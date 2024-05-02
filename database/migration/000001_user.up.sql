CREATE TABLE "user" (
    "id" serial PRIMARY KEY ,
    "email" varchar UNIQUE,
    "name" varchar,
    "password" varchar
);

CREATE INDEX idx_user_email ON "user"(email);