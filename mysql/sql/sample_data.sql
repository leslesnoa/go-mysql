DROP TABLE IF EXISTS `users`;

create table IF not exists `users`
(
 `id`               INT(20) NOT NULL AUTO_INCREMENT,
 `name`             VARCHAR(50) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO users (name) VALUES
    ('田中'),
    ('高橋'),
    ('渡辺');