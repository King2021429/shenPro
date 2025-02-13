-- 创建用户表
CREATE TABLE `users` (
                         `id` INT NOT NULL AUTO_INCREMENT,
                         `username` VARCHAR(255) NOT NULL UNIQUE,
                         `password` VARCHAR(255) NOT NULL,
                         `email` VARCHAR(255),
                         `avatar` VARCHAR(255),
                         `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         INDEX `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建文章表
CREATE TABLE `articles` (
                            `id` INT NOT NULL AUTO_INCREMENT,
                            `uid` INT NOT NULL,
                            `title` VARCHAR(255) NOT NULL,
                            `content` TEXT NOT NULL,
                            `published_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            INDEX `idx_deleted_at` (`deleted_at`),
                            INDEX `idx_uid` (`uid`),
                            FOREIGN KEY (`uid`) REFERENCES `users`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建评论表
CREATE TABLE `comments` (
                            `id` INT NOT NULL AUTO_INCREMENT,
                            `uid` INT NOT NULL,
                            `article_id` INT NOT NULL,
                            `parent_comment_id` INT DEFAULT NULL,
                            `content` VARCHAR(255) NOT NULL,
                            `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            INDEX `idx_deleted_at` (`deleted_at`),
                            INDEX `idx_uid` (`uid`),
                            INDEX `idx_article_id` (`article_id`),
                            INDEX `idx_parent_comment_id` (`parent_comment_id`),
                            FOREIGN KEY (`uid`) REFERENCES `users`(`id`),
                            FOREIGN KEY (`article_id`) REFERENCES `articles`(`id`),
                            FOREIGN KEY (`parent_comment_id`) REFERENCES `comments`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建用户关注表
CREATE TABLE `user_follows` (
                                `id` INT NOT NULL AUTO_INCREMENT,
                                `follower_id` INT NOT NULL,
                                `followee_id` INT NOT NULL,
                                `status` INT NOT NULL DEFAULT 0,
                                `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                INDEX `idx_deleted_at` (`deleted_at`),
                                INDEX `idx_follower_id` (`follower_id`),
                                INDEX `idx_followee_id` (`followee_id`),
                                FOREIGN KEY (`follower_id`) REFERENCES `users`(`id`),
                                FOREIGN KEY (`followee_id`) REFERENCES `users`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建文章收藏表
CREATE TABLE `article_collections` (
                                       `id` INT NOT NULL AUTO_INCREMENT,
                                       `uid` INT NOT NULL,
                                       `article_id` INT NOT NULL,
                                       `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                       `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                       `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                                       PRIMARY KEY (`id`),
                                       INDEX `idx_deleted_at` (`deleted_at`),
                                       INDEX `idx_uid` (`uid`),
                                       INDEX `idx_article_id` (`article_id`),
                                       FOREIGN KEY (`uid`) REFERENCES `users`(`id`),
                                       FOREIGN KEY (`article_id`) REFERENCES `articles`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建文章点赞表
-- 创建文章点赞表
CREATE TABLE `article_likes` (
                                 `id` INT NOT NULL AUTO_INCREMENT,
                                 `uid` INT NOT NULL,
                                 `article_id` INT NOT NULL,
                                 `status` INT NOT NULL DEFAULT 0,
                                 `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                                 PRIMARY KEY (`id`),
                                 INDEX `idx_uid` (`uid`),
                                 INDEX `idx_article_id` (`article_id`),
                                 FOREIGN KEY (`uid`) REFERENCES `users`(`id`),
                                 FOREIGN KEY (`article_id`) REFERENCES `articles`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 用户表（users）：
-- id 是用户的唯一标识，自增主键。
-- username 字段唯一且不能为空，用于存储用户名。
-- password 字段存储用户密码。
-- email 和 avatar 字段分别存储用户的电子邮箱和头像链接。
-- created_at 和 updated_at 字段记录用户创建和更新的时间。
-- deleted_at 字段用于软删除，默认值为 NULL，表示未删除。添加了 idx_deleted_at 索引用于优化软删除相关的查询。

-- 文章表（articles）：
-- id 是文章的唯一标识，自增主键。
-- uid 是文章作者的用户 ID，通过外键关联到 users 表的 id 字段。
-- title 和 content 字段分别存储文章的标题和内容。
-- published_at 字段记录文章的发布时间。
-- 同样包含 created_at、updated_at 和 deleted_at 字段，以及相应的索引。

-- 评论表（comments）：
-- id 是评论的唯一标识，自增主键。
-- uid 和 article_id 分别是评论者的用户 ID 和被评论文章的 ID，通过外键关联到 users 表和 articles 表。
-- parent_comment_id 字段用于实现多级评论，可为 NULL。
-- 包含多个索引，以优化不同查询条件下的性能。

-- 用户关注表（user_follows）：
-- id 是关注关系的唯一标识，自增主键。
-- follower_id 和 followee_id 分别是关注者和被关注者的用户 ID，通过外键关联到 users 表。

-- 文章收藏表（article_collections）：
-- id 是收藏关系的唯一标识，自增主键。
-- uid 和 article_id 分别是收藏者的用户 ID 和被收藏文章的 ID，通过外键关联到 users 表和 articles 表。

-- 文章点赞表（article_likes）：
-- id 是点赞关系的唯一标识，自增主键。
-- uid 和 article_id 分别是点赞者的用户 ID 和被点赞文章的 ID，通过外键关联到 users 表和 articles 表。