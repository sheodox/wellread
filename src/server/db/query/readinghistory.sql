-- name: ListReadingHistory :many
select *
from reading_history
where volume_id = @volume_id and user_id = @user_id
order by created_at desc;

-- name: AddReadingHistory :exec
insert into reading_history
(volume_id, current_page, created_at, user_id, pages_read)
values (@volume_id, @current_page, @created_at, @user_id, @pages_read);

-- name: DeleteReadingHistory :exec
delete from reading_history
where id = @reading_history_id and user_id = @user_id;
