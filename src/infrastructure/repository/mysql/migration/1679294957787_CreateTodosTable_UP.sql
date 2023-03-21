CREATE TABLE `todos` (
    `todo_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `activity_group_id` bigint unsigned NOT NULL,
    `title` varchar(500) NOT NULL,
    `priority` varchar(50) NOT NULL,
    `is_active` tinyint NOT NULL DEFAULT 0,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`todo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;