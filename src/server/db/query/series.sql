-- name: ListSeries :many
select series.*, count(volumes.id) as volume_count
from series
left join volumes on volumes.series_id = series.id
where series.user_id = @series_id
group by series.id;

-- name: GetSeries :one
select series.*, count(volumes.id) as volume_count
from series
left join volumes on volumes.series_id = series.id
where series.user_id = @user_id and series.id = @series_id
group by series.id;

-- name: AddSeries :one
insert into series (name, created_at, user_id)
values (@name, @created_at, @user_id)
returning id;

-- name: DeleteSeries :exec
delete from series
where id = @series_id and user_id = @user_id;

-- name: UpdateSeries :exec
update series
set name = @name, notes = @notes
where id = @series_id and user_id = @user_id;
