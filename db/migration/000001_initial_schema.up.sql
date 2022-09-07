CREATE TABLE "categories" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "breeds" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "category_id" int
);

CREATE TABLE "locations" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "pets" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "breed_id" int,
  "age" float,
  "location_id" int,
  "image_url" varchar,
  "description" varchar
);

ALTER TABLE "breeds" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "pets" ADD FOREIGN KEY ("breed_id") REFERENCES "breeds" ("id");

ALTER TABLE "pets" ADD FOREIGN KEY ("location_id") REFERENCES "locations" ("id");