-- name: CreateTemplate :exec
INSERT INTO service_templates (id, name, fields, created_by, created_at, updated_at) VALUES($1, $2, $3, $4, now(), null);

-- name: CreateCatalog :exec
INSERT INTO service_catalog (id, template) VALUES($1, $2);

-- name: GetAllTemplates :many
SELECT * FROM service_templates ORDER BY created_at;

-- name: GetTemplateByID :one
SELECT * FROM service_templates WHERE id = $1;

-- name: GetCatalog :one
SELECT * FROM service_catalog LIMIT 1;

-- name: CreateEntry :exec
INSERT INTO service_entry (id, catalog_id, data, created_at, updated_at) VALUES ($1, $2, $3, now(), null);

-- name: GetAllEntries :many
SELECT * FROM service_entry ORDER BY created_at;
