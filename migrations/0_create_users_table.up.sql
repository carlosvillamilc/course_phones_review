CREATE TABLE user (
    `id`                int(11) not null auto_increment PRIMARY KEY,
    `username`          varchar(150),
    `password`		      varchar(150),
    `created_at`        DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        DATETIME ON UPDATE CURRENT_TIMESTAMP
) engine = InnoDB DEFAULT charset = utf8;	
