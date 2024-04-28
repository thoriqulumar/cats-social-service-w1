CREATE TABLE "match" (
     "id" integer PRIMARY KEY,
     "issuedId" integer,
     "matchCatId" integer,
     "userCatId" integer,
     "message" varchar,
     "createdAt" date
);

ALTER TABLE "match" ADD FOREIGN KEY ("issuedId") REFERENCES "user" ("id");

ALTER TABLE "match" ADD FOREIGN KEY ("matchCatId") REFERENCES "cat" ("id");

ALTER TABLE "match" ADD FOREIGN KEY ("userCatId") REFERENCES "cat" ("id");