CREATE TABLE `tab_user` (
  `name` varchar(140) NOT NULL,
  `creation` datetime(6) DEFAULT NULL,
  `modified` datetime(6) DEFAULT NULL,
  `modified_by` varchar(140) DEFAULT NULL,
  `owner` varchar(140) DEFAULT NULL,
  `docstatus` int(1) NOT NULL DEFAULT 0,
  `parent` varchar(140) DEFAULT NULL,
  `parentfield` varchar(140) DEFAULT NULL,
  `parenttype` varchar(140) DEFAULT NULL,
  `idx` int(8) NOT NULL DEFAULT 0,
  `enabled` int(1) NOT NULL DEFAULT 1,
  `email` varchar(140) DEFAULT NULL,
  `first_name` varchar(140) DEFAULT NULL,
  `middle_name` varchar(140) DEFAULT NULL,
  `last_name` varchar(140) DEFAULT NULL,
  `full_name` varchar(140) DEFAULT NULL,
  `username` varchar(140) DEFAULT NULL,
  `language` varchar(140) DEFAULT NULL,
  `time_zone` varchar(140) DEFAULT NULL,
  `user_image` text DEFAULT NULL,
  `role_profile_name` varchar(140) DEFAULT NULL,
  `gender` varchar(140) DEFAULT NULL,
  `birth_date` date DEFAULT NULL,
  `interest` text DEFAULT NULL,
  `phone` varchar(140) DEFAULT NULL,
  `location` varchar(140) DEFAULT NULL,
  `bio` text DEFAULT NULL,
  `mute_sounds` int(1) NOT NULL DEFAULT 0,
  `mobile_no` varchar(140) DEFAULT NULL,
  `new_password` text DEFAULT NULL,
  `logout_all_sessions` int(1) NOT NULL DEFAULT 1,
  `reset_password_key` varchar(140) DEFAULT NULL,
  `last_password_reset_date` date DEFAULT NULL,
  `redirect_url` text DEFAULT NULL,
  `document_follow_notify` int(1) NOT NULL DEFAULT 0,
  `document_follow_frequency` varchar(140) DEFAULT 'Daily',
  `email_signature` text DEFAULT NULL,
  `thread_notify` int(1) NOT NULL DEFAULT 1,
  `send_me_a_copy` int(1) NOT NULL DEFAULT 0,
  `allowed_in_mentions` int(1) NOT NULL DEFAULT 1,
  `module_profile` varchar(140) DEFAULT NULL,
  `home_settings` longtext DEFAULT NULL,
  `simultaneous_sessions` int(11) NOT NULL DEFAULT 1,
  `restrict_ip` text DEFAULT NULL,
  `last_ip` varchar(140) DEFAULT NULL,
  `login_after` int(11) NOT NULL DEFAULT 0,
  `user_type` varchar(140) DEFAULT 'System User',
  `last_active` datetime(6) DEFAULT NULL,
  `login_before` int(11) NOT NULL DEFAULT 0,
  `bypass_restrict_ip_check_if_2fa_enabled` int(1) NOT NULL DEFAULT 0,
  `last_login` varchar(140) DEFAULT NULL,
  `last_known_versions` text DEFAULT NULL,
  `api_key` varchar(140) DEFAULT NULL,
  `api_secret` text DEFAULT NULL,
  `_user_tags` text DEFAULT NULL,
  `_comments` text DEFAULT NULL,
  `_assign` text DEFAULT NULL,
  `_liked_by` text DEFAULT NULL,
  PRIMARY KEY (`name`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `mobile_no` (`mobile_no`),
  UNIQUE KEY `api_key` (`api_key`),
  KEY `parent` (`parent`),
  KEY `modified` (`modified`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPRESSED
