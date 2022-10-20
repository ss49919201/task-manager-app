DROP TABLE IF EXISTS task_priorities;

CREATE TABLE
    IF NOT EXISTS task_priorities (
        `id` INT UNSIGNED NOT NULL,
        `value` INT UNSIGNED NOT NULL,
        PRIMARY KEY (`id`)
    ) DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;