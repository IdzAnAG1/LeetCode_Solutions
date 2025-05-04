CREATE TABLE "level" (
                         "level_id" SERIAL PRIMARY KEY,
                         "difficulty_level" int
);

CREATE TABLE "category" (
                            "category_id" SERIAL PRIMARY KEY,
                            "category_name" varchar(128) UNIQUE NOT NULL
);

CREATE TABLE "language" (
                            "language_id" SERIAL PRIMARY KEY,
                            "language_name" varchar(32) UNIQUE NOT NULL
);

CREATE TABLE "tasks" (
                         "task_id" SERIAL PRIMARY KEY,
                         "task_number" int UNIQUE NOT NULL,
                         "task_name" varchar(256),
                         "task_description" text,
                         "level_id" int,
                         FOREIGN KEY ("level_id") REFERENCES "level" ("level_id")
);

CREATE TABLE "solution" (
                            "solution_id" SERIAL PRIMARY KEY,
                            "task_id" int NOT NULL,
                            "language_id" int NOT NULL,
                            "solution_text" text,
                            "complexity" varchar(8),
                            "update_at" timestamp,
                            "created_at" timestamp,
                            UNIQUE ("task_id", "language_id"),
                            FOREIGN KEY ("task_id") REFERENCES "tasks" ("task_id"),
                            FOREIGN KEY ("language_id") REFERENCES "language" ("language_id")
);

CREATE TABLE "task_category" (
                                 "task_id" int NOT NULL,
                                 "category_id" int NOT NULL,
                                 PRIMARY KEY ("task_id", "category_id"),
                                 FOREIGN KEY ("task_id") REFERENCES "tasks" ("task_id"),
                                 FOREIGN KEY ("category_id") REFERENCES "category" ("category_id")
);
