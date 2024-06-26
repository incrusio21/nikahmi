package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Storage interface that is implemented by storage providers
type Storage struct {
	db         *sql.DB
	gcInterval time.Duration
	done       chan struct{}

	sqlSelect     string
	sqlSelectUser string
	sqlInsert     string
	sqlDelete     string
	sqlReset      string
	sqlGC         string
}

var (
	checkSchemaMsg = "The `value` row has an incorrect data type. " +
		"It should be BLOB but is instead %s. This will cause encoding-related panics if the DB is not migrated (see https://github.com/gofiber/storage/blob/main/MIGRATE.md)."
	dropQuery = "DROP TABLE IF EXISTS %s;"
	initQuery = []string{
		`CREATE TABLE IF NOT EXISTS %s ( 
			sess_key  VARCHAR(64) NOT NULL DEFAULT '', 
			value  BLOB NOT NULL, 
			exp  BIGINT NOT NULL DEFAULT '0', 
			name  VARCHAR(140) NOT NULL DEFAULT '', 
			PRIMARY KEY (sess_key)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
	}
	checkSchemaQuery = `SELECT DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS
		WHERE table_name = '%s' AND COLUMN_NAME = 'value';`
)

// New creates a new storage
func New(config ...Config) *Storage {
	var err error
	var db *sql.DB

	// Set default config
	cfg := configDefault(config...)

	if cfg.Db != nil {
		// Use passed db
		db = cfg.Db
	} else {
		// Create db
		db, err = sql.Open("mysql", cfg.dsn())
		if err != nil {
			panic(err)
		}

		// Set options
		db.SetMaxOpenConns(cfg.maxOpenConns)
		db.SetMaxIdleConns(cfg.maxIdleConns)
		db.SetConnMaxLifetime(cfg.connMaxLifetime)
	}

	// Ping database to ensure a connection has been made
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Drop table if Clear set to true
	if cfg.Reset {
		query := fmt.Sprintf(dropQuery, cfg.Table)
		if _, err = db.Exec(query); err != nil {
			_ = db.Close()
			panic(err)
		}
	}

	// Init database queries
	for _, query := range initQuery {
		query = fmt.Sprintf(query, cfg.Table)
		if _, err := db.Exec(query); err != nil {
			_ = db.Close()
			panic(err)
		}
	}

	// Create storage
	store := &Storage{
		gcInterval:    cfg.GCInterval,
		db:            db,
		done:          make(chan struct{}),
		sqlSelect:     fmt.Sprintf("SELECT value, exp, name FROM %s WHERE sess_key=?;", cfg.Table),
		sqlSelectUser: fmt.Sprintf("SELECT count(*) FROM %s WHERE name=?;", cfg.Table),
		sqlInsert:     fmt.Sprintf("INSERT INTO %s (sess_key, value, exp, name) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE value = ?, exp = ?", cfg.Table),
		sqlDelete:     fmt.Sprintf("DELETE FROM %s WHERE sess_key=?", cfg.Table),
		sqlReset:      fmt.Sprintf("TRUNCATE TABLE %s;", cfg.Table),
		sqlGC:         fmt.Sprintf("DELETE FROM %s WHERE exp <= ? AND e != 0", cfg.Table),
	}

	store.checkSchema(cfg.Table)

	// Start garbage collector
	go store.gcTicker()

	return store
}

// Get value by key
func (s *Storage) Get(key string) ([]byte, string, error) {
	if len(key) <= 0 {
		return nil, "", nil
	}
	row := s.db.QueryRow(s.sqlSelect, key)

	// Add db response to data

	var (
		data []byte
		exp  int64
		name string
	)

	if err := row.Scan(&data, &exp, &name); err != nil {
		if err == sql.ErrNoRows {
			return nil, "", nil
		}
		return nil, "", err
	}

	// If the expiration time has already passed, then return nil
	if exp != 0 && exp <= time.Now().Unix() {
		return nil, "", nil
	}

	return data, "", nil
}

// Get value by key
func (s *Storage) GetUser(user string, max_session int) error {
	if len(user) <= 0 || max_session == 0 {
		return nil
	}
	row := s.db.QueryRow(s.sqlSelectUser, user)

	// Add db response to data

	var (
		total_session int
	)

	if err := row.Scan(&total_session); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	// If the expiration time has already passed, then return nil
	if max_session <= total_session {
		return errors.New("user is unable to log in because they have exceeded the maximum number of devices")
	}

	return nil
}

// Set key with value
func (s *Storage) Set(key string, val []byte, exp time.Duration, user string) error {
	// Ain't Nobody Got Time For That
	if len(key) <= 0 || len(val) <= 0 {
		return nil
	}
	var expSeconds int64
	if exp != 0 {
		expSeconds = time.Now().Add(exp).Unix()
	}
	_, err := s.db.Exec(s.sqlInsert, key, val, expSeconds, user, val, expSeconds)
	return err
}

// Delete key by key
func (s *Storage) Delete(key string) error {
	// Ain't Nobody Got Time For That
	if len(key) <= 0 {
		return nil
	}
	_, err := s.db.Exec(s.sqlDelete, key)
	return err
}

// Reset all keys
func (s *Storage) Reset() error {
	_, err := s.db.Exec(s.sqlReset)
	return err
}

// Close the database
func (s *Storage) Close() error {
	s.done <- struct{}{}
	return s.db.Close()
}

// Return database client
func (s *Storage) Conn() *sql.DB {
	return s.db
}

// gcTicker starts the gc ticker
func (s *Storage) gcTicker() {
	ticker := time.NewTicker(s.gcInterval)
	defer ticker.Stop()
	for {
		select {
		case <-s.done:
			return
		case t := <-ticker.C:
			s.gc(t)
		}
	}
}

// gc deletes all expired entries
func (s *Storage) gc(t time.Time) {
	_, _ = s.db.Exec(s.sqlGC, t.Unix())
}

func (s *Storage) checkSchema(tableName string) {
	var data []byte

	row := s.db.QueryRow(fmt.Sprintf(checkSchemaQuery, tableName))
	if err := row.Scan(&data); err != nil {
		panic(err)
	}

	if strings.ToLower(string(data)) != "blob" {
		fmt.Printf(checkSchemaMsg, string(data))
	}
}
