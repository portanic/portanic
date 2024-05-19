-- +goose Up
CREATE TABLE "service_entry" (
  "id" uuid PRIMARY KEY,
  "catalog_id" uuid,
  "data" jsonb,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp
);

ALTER TABLE "service_entry" ADD FOREIGN KEY ("catalog_id") REFERENCES "service_catalog" ("id");

-- +goose Down
DROP TABLE "service_entry" CASCADE;
