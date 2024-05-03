migrate create -ext sql -dir db/migrations nama_file_migration

go run main.go -migrate="migrate_name"