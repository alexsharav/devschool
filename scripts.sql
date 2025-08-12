create table users (
    Id serial primary key,
    username varchar(50) not null,
    email varchar(100) not null,
    password text not null,
    role varchar(10) not null default 'student',
    courses varchar(100)[]
);

