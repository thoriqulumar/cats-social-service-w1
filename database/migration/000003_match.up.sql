CREATE TABLE "match" (
     "id" serial PRIMARY KEY,
     "issuedId" integer,
     "matchCatId" integer,
     "userCatId" integer,
     "message" varchar,
     "isApprovedOrRejected" boolean,
     "createdAt" date
);

ALTER TABLE "match" ADD FOREIGN KEY ("issuedId") REFERENCES "user" ("id");

ALTER TABLE "match" ADD FOREIGN KEY ("matchCatId") REFERENCES "cat" ("id");

ALTER TABLE "match" ADD FOREIGN KEY ("userCatId") REFERENCES "cat" ("id");