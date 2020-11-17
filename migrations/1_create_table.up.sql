CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL
);

CREATE TABLE friendship(
	id SERIAL,
	first_user TEXT NOT NULL,
	second_user TEXT NOT NULL,
	is_friend BOOL NULL DEFAULT false,
	update_status INTEGER NULL DEFAULT 0,
	PRIMARY KEY (id, first_user, second_user),
  	FOREIGN KEY (first_user)
      REFERENCES users (email),
  	FOREIGN KEY (second_user)
      REFERENCES users (email)
)