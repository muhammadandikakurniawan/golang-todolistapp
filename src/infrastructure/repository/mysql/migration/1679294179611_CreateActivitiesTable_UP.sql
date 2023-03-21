CREATE TABLE `activities` (
    `activity_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(500) NOT NULL,
    `email` varchar(100) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
