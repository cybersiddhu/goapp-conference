CREATE TABLE users (
	 id INTEGER PRIMARY KEY autoincrement not null,
	 email TEXT UNIQUE,
	 password TEXT UNIQUE,
	 firstname TEXT,
	 lastname TEXT
);
