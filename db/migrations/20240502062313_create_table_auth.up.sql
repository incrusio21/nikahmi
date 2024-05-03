CREATE TABLE `__Auth` (
  `doctype` varchar(140) NOT NULL,
  `name` varchar(255) NOT NULL,
  `fieldname` varchar(140) NOT NULL,
  `password` text NOT NULL,
  `encrypted` int(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`doctype`,`name`,`fieldname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPRESSED