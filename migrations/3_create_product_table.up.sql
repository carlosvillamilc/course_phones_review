CREATE TABLE product(
    `id`                varchar(150) not null PRIMARY KEY,
    `name`              varchar(150),
    `price`             int(10),
    `created_at`        DATETIME DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        DATETIME ON UPDATE CURRENT_TIMESTAMP    
) engine = InnoDB DEFAULT charset = utf8;