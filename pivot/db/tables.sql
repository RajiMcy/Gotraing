CREATE TABLE employee (
	attuid VARCHAR(6) PRIMARY KEY,
	first_nm TEXT NOT NULL,
	last_nm TEXT NOT NULL,
	email VARCHAR NOT NULL UNIQUE,
	mgr_id VARCHAR(6),
	status TEXT NOT NULL,
	last_modified DATETIME NOT NULL DEFAULT (strftime('%Y-%m-%d %H:%M:%f', 'now', 'localtime'))
);

CREATE TABLE skill (
	skill_id INTEGER PRIMARY KEY AUTOINCREMENT,
	skill_nm VARCHAR NOT NULL UNIQUE
);

CREATE TABLE employeeskill (
	attuid VARCHAR(6),
	skill_id INTEGER,
	proficiency INTEGER NOT NULL,
	version VARCHAR,
	last_used DATE NOT NULL,
	is_primary CHAR NOT NULL,
	total_exp_in_month INTEGER NOT NULL,
	PRIMARY KEY (attuid, skill_id),
    FOREIGN KEY (attuid) 
      REFERENCES employee (attuid) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION,
    FOREIGN KEY (skill_id) 
      REFERENCES skill (skill_id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
);
