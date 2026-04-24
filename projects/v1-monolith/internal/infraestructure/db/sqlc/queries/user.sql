-- name: CreateUser :exec
INSERT INTO users (id, full_name, email, role)
VALUES (?, ?, ?, ?);

-- name: UpdateUser :exec 
update users set full_name = ?, email = ? where id = ?;

-- name: DeleteUser :exec
DELETE FROM users where id = ?;


-- name: ListUsers :many
select * from users;

-- name: ListById :one
select * from users where id = ?;
