	-- Create tables.
	DROP TABLE IF EXISTS "usersList";
	DROP TABLE IF EXISTS "forums";
	DROP TABLE IF EXISTS "interestList";
	DROP TABLE IF EXISTS "users";
	CREATE TABLE "users"
	(
		"id"   SERIAL PRIMARY KEY,
		"name" VARCHAR(50) NOT NULL UNIQUE
	);

	CREATE TABLE "interestList"
	(
		"id"   SERIAL PRIMARY KEY,
		"interest" VARCHAR(50) NULL,
		"userID" INT NOT NULL,
		FOREIGN KEY ("userID")
		REFERENCES users(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	CREATE TABLE "forums"
	(
		"id"   SERIAL PRIMARY KEY,
		"name" VARCHAR(50) NOT NULL UNIQUE,
		"topicKeyword" VARCHAR(50) NOT NULL UNIQUE
	);

	CREATE TABLE "usersList"
	(
		"id"   SERIAL PRIMARY KEY,
		"forumsID" INT NOT NULL,
		FOREIGN KEY ("forumsID")
		REFERENCES forums(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE,
		"userID" INT NULL,
		FOREIGN KEY ("userID")
		REFERENCES users(id)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);

	-- Insert demo data.
	INSERT INTO "users" (name) VALUES ('Bob');
	INSERT INTO "users" (name) VALUES ('Nick');
	INSERT INTO "users" (name) VALUES ('Simon');
	INSERT INTO "interestList" ("interest", "userID") VALUES ('Jojo References', 1);
	INSERT INTO "interestList" ("interest", "userID") VALUES ('Games', 1);
	INSERT INTO "interestList" ("interest", "userID") VALUES ('Games', 2);
	INSERT INTO "interestList" ("interest", "userID") VALUES ('Books', 3);
	INSERT INTO "forums" ("name", "topicKeyword") VALUES ('Jojo References', 'jojo bizzare adventure');
	INSERT INTO "forums" ("name", "topicKeyword") VALUES ('Movies fan', 'Movies');
	INSERT INTO "forums" ("name", "topicKeyword") VALUES ('Book enjoyer', 'Books');
	INSERT INTO "forums" ("name", "topicKeyword") VALUES ('Gaming', 'Games');
	INSERT INTO "usersList" ("forumsID", "userID") VALUES (1, 1);
	INSERT INTO "usersList" ("forumsID", "userID") VALUES (4, 1);
	INSERT INTO "usersList" ("forumsID", "userID") VALUES (4, 2);
	INSERT INTO "usersList" ("forumsID", "userID") VALUES (3, 3);