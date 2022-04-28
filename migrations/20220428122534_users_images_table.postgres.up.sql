create table users_images(
	id serial primary key,
	image text not null,
	user_id int
);

alter table users_images 
add constraint fk_ui_user_id foreign key(user_id)
references users(id)