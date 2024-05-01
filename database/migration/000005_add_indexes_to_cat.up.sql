ALTER TABLE "cat" RENAME COLUMN "isAlreadyMatched" TO "hasMatched";

CREATE INDEX "idx_cat_race" ON "cat" ("race");
CREATE INDEX "idx_cat_sex" ON "cat" ("sex");
CREATE INDEX "idx_cat_age" ON "cat" ("ageInMonth");
CREATE INDEX "idx_cat_owner" ON "cat" ("ownerId");
CREATE INDEX "idx_cat_text_search" ON "cat" ("name", "race", "sex", "description");
CREATE INDEX "idx_cat_matched" ON "cat" ("hasMatched");