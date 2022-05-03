CREATE TABLE t_User (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),
    password VARCHAR(255),
    email VARCHAR(255),
    created_at VARCHAR(50),
    updated_at VARCHAR(50)
);

CREATE TABLE t_Token (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    user_id INTEGER,
    name VARCHAR(255),
    token VARCHAR(255),
    type SMALLINT,
    created_at VARCHAR(50),
    updated_at VARCHAR(50),
    expired_at VARCHAR(50)
);