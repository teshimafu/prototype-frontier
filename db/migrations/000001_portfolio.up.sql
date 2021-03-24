CREATE TABLE IF NOT EXISTS portfolio (
	id serial PRIMARY KEY,
	title VARCHAR (256),
	uid VARCHAR (256),
	author VARCHAR (256),
	abstruct TEXT,
	source VARCHAR (1024),
	link VARCHAR (1024),
	readme TEXT,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);