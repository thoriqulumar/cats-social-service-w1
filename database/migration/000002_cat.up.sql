CREATE TYPE "race" AS ENUM (
        'Persian',
        'Maine Coon',
        'Siamese',
        'Ragdoll',
        'Bengal',
        'Sphynx',
        'British Shorthair',
        'Abyssinian',
        'Scottish Fold',
        'Birman'
    );

CREATE TYPE "sex" AS ENUM (
    'male',
    'female'
);


CREATE TABLE "cat" (
    "id" serial PRIMARY KEY,
    "ownerId" integer,
    "name" varchar,
    "race" race,
    "sex" sex,
    "ageInMonth" integer,
    "description" varchar,
    "imageUrls" varchar[],
    "isAlreadyMatched" boolean,
    "isDeleted" boolean,
    "createdAt" date
);

ALTER TABLE "cat" ADD FOREIGN KEY ("ownerId") REFERENCES "user" ("id");
