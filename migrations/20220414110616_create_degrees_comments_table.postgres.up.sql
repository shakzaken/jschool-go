create table degrees_comments(
	id serial primary key,
	degree_id int,
	user_id int,
	comment varchar(512),
	date timestamp 
);

alter table degrees_comments
add constraint fk_dc_degree_id foreign key(degree_id) 
references degrees(id);

alter table degrees_comments
add constraint fk_dc_user_id foreign key(user_id)
references users(id);