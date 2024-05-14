-- +goose Up
CREATE TABLE "organizations" (
  id uuid PRIMARY KEY,
  name varchar,
  created_at timestamp,
  updated_at timestamp
);

CREATE TABLE "mapped_data" (
  id uuid PRIMARY KEY,
  mapping_id uuid,
  data jsonb,
  created_at timestamp
);

CREATE TABLE "organizations_users" (
  organization_id uuid,
  user_id uuid,
  PRIMARY KEY(organization_id,user_id)
);

CREATE TABLE "users" (
  id uuid PRIMARY KEY,
  name varchar,
  email varchar UNIQUE NOT NULL,
  created_at timestamp,
  updated_at timestamp,
  last_login timestamp
);

CREATE TABLE "roles" (
  id uuid PRIMARY KEY,
  name varchar,
  description text
);

CREATE TABLE "organizations_users_roles" (
  organization_id uuid,
  user_id uuid,
  role_id uuid,
  PRIMARY KEY(organization_id,user_id,role_id)
);

CREATE TABLE "service_catalog" (
  id uuid PRIMARY KEY,
  template uuid
);

CREATE TABLE "service_templates" (
  id uuid PRIMARY KEY,
  name varchar NOT NULL,
  fields jsonb NOT NULL,
  created_by varchar NOT NULL,
  created_at timestamp DEFAULT (now()),
  updated_at timestamp
);

CREATE TABLE "data_source" (
  id uuid PRIMARY KEY,
  name varchar
);

CREATE TABLE "mappings" (
  id uuid,
  template uuid,
  data_source uuid
);

ALTER TABLE "service_catalog" ADD FOREIGN KEY ("template") REFERENCES "service_templates" ("id");
ALTER TABLE "mappings" ADD FOREIGN KEY ("template") REFERENCES "service_templates" ("id");
ALTER TABLE "mappings" ADD FOREIGN KEY ("data_source") REFERENCES "data_source" ("id");
ALTER TABLE "organizations_users_roles" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");
ALTER TABLE "organizations_users_roles" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "organizations_users_roles" ADD FOREIGN KEY ("role_id") REFERENCES "roles" ("id");
ALTER TABLE "organizations_users" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");
ALTER TABLE "organizations_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +goose Down
DROP TABLE "mapped_data" CASCADE;
DROP TABLE "mappings" CASCADE;
DROP TABLE "data_source" CASCADE;
DROP TABLE "service_templates" CASCADE;
DROP TABLE "service_catalog" CASCADE;
DROP TABLE "organizations_users_roles" CASCADE;
DROP TABLE "users" CASCADE;
DROP TABLE "roles" CASCADE;
DROP TABLE "organizations_users" CASCADE;
DROP TABLE "organizations" CASCADE;

