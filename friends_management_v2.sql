CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) UNIQUE NOT NULL,
	password VARCHAR(30) NOT NULL
)

CREATE TABLE friendship(
	id SERIAL,
	first_user VARCHAR(255) NOT NULL,
	second_user VARCHAR(255) NOT NULL,
	is_friend BOOL NULL DEFAULT false,
	update_status BOOL NULL DEFAULT false,
	PRIMARY KEY (id, first_user, second_user),
  	FOREIGN KEY (first_user)
      REFERENCES users (email),
  	FOREIGN KEY (second_user)
      REFERENCES users (email)
)