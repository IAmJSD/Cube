
CREATE TABLE IF NOT EXISTS `bans` (
  `user_id` text DEFAULT NULL,
  `staff_id` text DEFAULT NULL,
  `server_id` text DEFAULT NULL,
  `reason` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `bot_admins` (
  `user_id` mediumtext COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `custom_commands` (
  `server_id` mediumtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `command` mediumtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `response` mediumtext COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `custom_prefixes` (
  `prefix` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `server_id` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `kicks` (
  `user_id` text DEFAULT NULL,
  `staff_id` text DEFAULT NULL,
  `server_id` text DEFAULT NULL,
  `reason` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `logging_channels` (
  `server_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `channel_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `main_channels` (
  `server_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `channel_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `member_roles` (
  `part_of_role` text DEFAULT NULL,
  `server_id` text DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `merits` (
  `merit_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `staff_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `server_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `merit_reason` text COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE IF NOT EXISTS `msg_logs` (
  `message_id` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `server_id` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `channel_id` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `attachments` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `content` mediumtext COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `strikes` (
  `strike_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `staff_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `server_id` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `strike_reason` text COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
