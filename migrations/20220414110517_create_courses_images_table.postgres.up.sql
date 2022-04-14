create table courses_images(
	id serial primary key,
	image text not null,
	course_id int
);

alter table courses_images 
add constraint fk_ci_course_id foreign key(course_id)
references courses(id)