create table degrees_images(
	id serial primary key,
	image text not null,
	degree_id int
);

alter table degrees_images 
add constraint fk_di_degree_id foreign key(degree_id)
references degrees(id)