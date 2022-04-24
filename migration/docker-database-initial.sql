create table users(
    id serial primary key,
    name varchar,
    type varchar
);

INSERT INTO users(name, type) VALUES
('Test Admin', 'ADMIN'),
('Test Basic', 'BASIC');