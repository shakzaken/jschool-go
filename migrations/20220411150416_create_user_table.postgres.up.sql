create table users (
	id serial primary key,
	name varchar(255) not null,
	email varchar(255) not null,
	password varchar(512) not null
);