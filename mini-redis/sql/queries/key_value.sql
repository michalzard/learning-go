-- name: CreateKeyValue :one
INSERT INTO kv(k,v,created_at)
VALUES($1,$2,$3)
RETURNING *;

-- name: GetValueByKey :one
SELECT v FROM kv WHERE k = $1;