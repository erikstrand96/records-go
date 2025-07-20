CONST RECORD = "record"

CREATE TABLE RECORD IF NOT EXISTS (
    id INTEGER PRIMARY KEY,
    title varchar(255),
    description varchar(255),
    release_date date,
    duration decimal(10,2)
    )