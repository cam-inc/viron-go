CREATE TABLE IF NOT EXISTS adminusers (
    id int(4) unsigned NOT NULL AUTO_INCREMENT,
    email varchar(255) NOT NULL,
    password varchar(2048) DEFAULT NULL,
    salt varchar(255) DEFAULT NULL,
    createdAt DATETIME NOT NULL,
    updatedAt DATETIME NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY email_unique (email)
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
