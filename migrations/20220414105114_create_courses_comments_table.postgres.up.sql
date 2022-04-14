create table courses_comments(
	id serial primary key,
	course_id int,
	user_id int,
	comment varchar(512),
	date timestamp 
);

alter table courses_comments
add constraint fk_cc_course_id foreign key(course_id) 
references courses(id);

alter table courses_comments
add constraint fk_cc_user_id foreign key(user_id)
references users(id);