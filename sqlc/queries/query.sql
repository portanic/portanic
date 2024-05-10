-- name: CreateTemplate :exec
INSERT INTO service_templates (id, name, fields, created_by, created_at, updated_at) VALUES($1, $2, $3, $4, now(), null);

-- name: GetAllTemplates :many
SELECT * FROM service_templates ORDER BY created_at;
