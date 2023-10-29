CREATE DATABASE IF NOT EXISTS CraftEdu2;
USE CraftEdu2;

DROP TABLE IF EXISTS Users;

CREATE TABLE Users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(200) not null unique,
    nickname varchar(50) not null unique,
    password varchar(200) not null unique,
    dateCreated timestamp default current_timestamp()

) ENGINE=INNODB;