CREATE TABLE IF NOT EXISTS candidate (
	CREATE TABLE IF NOT EXISTS candidate (
	id SERIAL PRIMARY KEY,
	fullname VARCHAR(100) NOT NULL,
	about VARCHAR(100),
	jobtitle VARCHAR(100),
	experience VARCHAR(100),
    country VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS job (
	id SERIAL PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	jobName VARCHAR(255) NOT NULL
);
);

CREATE TABLE IF NOT EXISTS profile (
	id Auto_Increment PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	profilelink VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS skills (
	id SERIAL PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	skillsname VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS socialmedia (
	id INT PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	platform VARCHAR(255) NOT NULL,
	url VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS location (
	id SERIAL PRIMARY KEY,
	candidate_id INT REFERENCES candidate(id),
	locationname VARCHAR(255) NOT NULL,
	locationtype VARCHAR(255) NOT NULL
);`