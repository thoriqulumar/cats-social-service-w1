CREATE TYPE "matchStatus" AS ENUM (
    'waiting_for_approval',
    'approved',
    'rejected',
    'deleted'
);

CREATE TABLE "match" (
     "id" serial PRIMARY KEY,
     "issuedId" integer,
     "receiverId" integer,
     "matchCatId" integer,
     "userCatId" integer,
     "message" varchar,
     "status" "matchStatus",
     "createdAt" date
);

ALTER TABLE "match" ADD FOREIGN KEY ("issuedId") REFERENCES "user" ("id");

ALTER TABLE "match" ADD FOREIGN KEY ("receiverId") REFERENCES "user" ("id");

ALTER TABLE "match" ADD FOREIGN KEY ("matchCatId") REFERENCES "cat" ("id");

ALTER TABLE "match" ADD FOREIGN KEY ("userCatId") REFERENCES "cat" ("id");