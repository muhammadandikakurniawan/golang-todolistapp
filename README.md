# Golang todolist app


## Sql : 
```sql
CREATE TABLE `activities` (
    `activity_id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(500) NOT NULL,
    `email` varchar(100) NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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

```


Open api docs endpoint :
http://localhost:3030/api-docs/index.html


![api doc preview](/assets/api-docs.png "api doc preview")


#
## 🔗 Links
[![github](https://img.shields.io/badge/repository-black?style=for-the-badge&logo=github&logoColor=white)](https://github.com/muhammadandikakurniawan/golang-todolistapp/tree/main)

