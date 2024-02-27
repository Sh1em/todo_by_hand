CREATE TABLE users
(
    id serial not null UNIQUE,
    name varchar(255) not NULL,
    username VARCHAR(255) not NULL UNIQUE,
    password_hash VARCHAR(255) not NULL
);

CREATE TABLE todo_lists
(
    id serial not null UNIQUE,
    title VARCHAR(255) not NULL,
    description VARCHAR(255)
);

CREATE TABLE users_lists
(
    id serial not null UNIQUE,
    user_id int REFERENCES users(id) on delete cascade not NULL,
    list_id int REFERENCES todo_lists(id) on delete cascade not NULL

);


CREATE TABLE todo_items
(
    id serial not null UNIQUE,
    title VARCHAR(255) not NULL,
    description VARCHAR(255),
    done BOOLEAN not Null DEFAULT FALSE
);

CREATE TABLE lists_items
(
    id serial not null UNIQUE,
    item_id int REFERENCES todo_items(id) on delete cascade not NULL,
    list_id int REFERENCES todo_lists(id) on delete cascade not NULL

);