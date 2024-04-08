-- name: CreateKeyValue :one
INSERT INTO kv(k,v,created_at)
VALUES($1,$2,$3)
RETURNING *;
