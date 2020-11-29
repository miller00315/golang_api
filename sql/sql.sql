CREATE USER 'golang'@'localhost' IDENTIFIED BY 'golang';

CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id  int not null,
    follower_id int not null,
    
    FOREIGN KEY (user_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE,

    FOREIGN KEY (follower_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE,

    primary key(user_id, follower_id)
) ENGINE=INNODB;

GRANT ALL PRIVILEGES ON devbook.* TO 'golang'@'localhost';