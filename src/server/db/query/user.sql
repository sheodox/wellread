-- name: GetUser :one
select *
from users
where id = @user_id;

-- name: GetUserByFirebaseId :one
select *
from users
where firebase_user_id = @firebase_user_id;

-- name: AddUser :one
insert into users
(provider_id, firebase_user_id, email, display_name, created_at)
values (@provider_id, @firebase_user_id, @email, @display_name, @created_at)
returning *;
