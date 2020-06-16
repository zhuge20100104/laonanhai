-- ---------------------------------
-- 这是注释
-- BMS
-- ---------------------------------

CREATE TABLE book(
    id bigint(20) AUTO_INCREMENT PRIMARY KEY,
    title varchar(20) NOT NULL,
    price double(16,2) NOT NULL
) engine=InnoDB default charset=utf8;