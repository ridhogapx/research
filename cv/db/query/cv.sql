-- name: CreateUser :exec
INSERT INTO cvzn_users (
	email
) VALUES (
	$1
);
