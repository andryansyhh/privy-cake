-- +migrate Up
CREATE TABLE cake (
	id int NOT NULL AUTO_INCREMENT,
	title text NULL,
	description text NULL,
	rating numeric NULL,
	image varchar(255) NULL,
	created_at timestamp NULL DEFAULT now(),
	updated_at timestamp NULL DEFAULT now(),
	deleted_at timestamp NULL,
	PRIMARY KEY (id)
);
