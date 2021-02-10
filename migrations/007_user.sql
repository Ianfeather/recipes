DROP TABLE recipe_user;
DROP TABLE user;
ALTER TABLE recipe ADD user_id varchar(255) NOT NULL;
ALTER TABLE list MODIFY user_id varchar(255) NOT NULL;
