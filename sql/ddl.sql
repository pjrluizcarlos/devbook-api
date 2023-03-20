CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS post;
DROP TABLE IF EXISTS follower;
DROP TABLE IF EXISTS user;

CREATE TABLE user (
    id int auto_increment primary key,
    name varchar(50) not null, 
    nick varchar(50) not null unique,
    email varchar(50) not null unique, 
    password varchar(200) not null, 
    created_at timestamp default current_timestamp()  
) ENGINE=INNODB;

CREATE TABLE follower (
    user_id int not null, 
    follower_id int not null, 
    primary key (user_id, follower_id),
    foreign key user_id_fk (user_id) references user(id) on delete cascade,
    foreign key follower_id_fk (follower_id) references user(id) on delete cascade
) ENGINE=INNODB;

CREATE TABLE post (
    id int auto_increment primary key,
    title varchar(50) not null, 
    content varchar(300) not null, 
    user_id int not null, 
    likes int default 0,
    created_at timestamp default current_timestamp(),
    foreign key user_fk (user_id) references user(id) on delete cascade

) ENGINE=INNODB;

CREATE USER IF NOT EXISTS 'devbookapi'@'localhost' IDENTIFIED BY 'devbookapi';
GRANT ALL PRIVILEGES ON devbook.* TO 'devbookapi'@'localhost';
