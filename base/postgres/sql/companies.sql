create table companies (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
name VARCHAR(15) UNIQUE NOT NULL,
description TEXT,
employees INT NOT NULL,
registered BOOLEAN NOT NULL,
type VARCHAR(20) NOT NULL
v text DEFAULT v1,
uts timestamp default current_timestamp
);
