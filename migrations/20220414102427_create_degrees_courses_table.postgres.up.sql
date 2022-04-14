create table degrees_courses (
	id serial primary key,
	degree_id int ,
	course_id int ,
	unique(degree_id,course_id)
);

alter table degrees_courses
add constraint fk_dc_degree_id foreign key(degree_id) references degrees(id);

alter table degrees_courses 
add constraint fk_dc_course_id foreign key(course_id) references courses(id);