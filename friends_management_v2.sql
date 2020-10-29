CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
)

CREATE TABLE friendship(
	id SERIAL,
	first_user TEXT NOT NULL,
	second_user TEXT NOT NULL,
	is_friend BOOL NULL DEFAULT false,
	update_status BOOL NULL DEFAULT false,
	PRIMARY KEY (id, first_user, second_user),
  	FOREIGN KEY (first_user)
      REFERENCES users (email),
  	FOREIGN KEY (second_user)
      REFERENCES users (email)
)