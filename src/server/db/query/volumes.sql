-- name: ListVolumes :many
with vols as (
	select volumes.*, series.name as series_name
	from volumes
	inner join series on volumes.series_id = series.id
	where volumes.user_id = @user_id 
	and volumes.status = coalesce(sqlc.narg('status'), volumes.status)
	and volumes.name like coalesce(sqlc.narg('name'), volumes.name)
	and series_id = coalesce(sqlc.narg('series_id'), volumes.series_id) 
	order by name asc
),
paged as (
	select vols.*
	from vols
	limit @page_size
	offset @page_offset
)
select paged.*, (select count(vols.id) from vols) as total_results
from paged;

-- name: GetVolume :one
select volumes.*, series.name as series_name
from volumes
inner join series on volumes.series_id = series.id
where volumes.user_id = @user_id and volumes.id = @volume_id;


-- name: AddVolume :one
insert into volumes
(series_id, name, created_at, user_id)
values (@series_id, @name, @created_at, @user_id)
returning *;

-- name: DeleteVolume :exec
delete from volumes where id = @volume_id and user_id = @user_id;

-- name: UpdateVolume :exec
update volumes
set notes = @notes, current_page = @current_page, name = @name, status = @status
where id = @volume_id and user_id = @user_id;
