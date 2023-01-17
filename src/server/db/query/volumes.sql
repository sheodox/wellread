-- name: ListVolumes :many
select volumes.*, series.name as series_name
from volumes
inner join series on volumes.series_id = series.id
where volumes.user_id = @user_id
order by name asc;

-- name: ListVolumesBySeries :many
select volumes.*, series.name as series_name
from volumes
inner join series on volumes.series_id = series.id
where volumes.user_id = @user_id
and series_id = @series_id order by name asc;

-- name: GetVolume :one
select volumes.*, series.name as series_name
from volumes
inner join series on volumes.series_id = series.id
where volumes.user_id = @user_id and volumes.id = @volume_id;


-- name: ListVolumesByStatus :many
select volumes.*, series.name as series_name
from volumes
inner join series on volumes.series_id = series.id
where volumes.user_id = @user_id
and status = @status
order by name asc;

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
