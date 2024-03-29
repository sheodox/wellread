// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: series.sql

package query

import (
	"context"
	"time"
)

const addSeries = `-- name: AddSeries :one
insert into series (name, created_at, user_id)
values ($1, $2, $3)
returning id
`

type AddSeriesParams struct {
	Name      string
	CreatedAt time.Time
	UserID    int32
}

func (q *Queries) AddSeries(ctx context.Context, arg AddSeriesParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, addSeries, arg.Name, arg.CreatedAt, arg.UserID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteSeries = `-- name: DeleteSeries :exec
delete from series
where id = $1 and user_id = $2
`

type DeleteSeriesParams struct {
	SeriesID int32
	UserID   int32
}

func (q *Queries) DeleteSeries(ctx context.Context, arg DeleteSeriesParams) error {
	_, err := q.db.ExecContext(ctx, deleteSeries, arg.SeriesID, arg.UserID)
	return err
}

const getSeries = `-- name: GetSeries :one
select series.id, series.name, series.notes, series.created_at, series.user_id, count(volumes.id) as volume_count
from series
left join volumes on volumes.series_id = series.id
where series.user_id = $1 and series.id = $2
group by series.id
`

type GetSeriesParams struct {
	UserID   int32
	SeriesID int32
}

type GetSeriesRow struct {
	ID          int32
	Name        string
	Notes       string
	CreatedAt   time.Time
	UserID      int32
	VolumeCount int64
}

func (q *Queries) GetSeries(ctx context.Context, arg GetSeriesParams) (GetSeriesRow, error) {
	row := q.db.QueryRowContext(ctx, getSeries, arg.UserID, arg.SeriesID)
	var i GetSeriesRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Notes,
		&i.CreatedAt,
		&i.UserID,
		&i.VolumeCount,
	)
	return i, err
}

const listSeries = `-- name: ListSeries :many
select series.id, series.name, series.notes, series.created_at, series.user_id, count(volumes.id) as volume_count
from series
left join volumes on volumes.series_id = series.id
where series.user_id = $1
group by series.id
`

type ListSeriesRow struct {
	ID          int32
	Name        string
	Notes       string
	CreatedAt   time.Time
	UserID      int32
	VolumeCount int64
}

func (q *Queries) ListSeries(ctx context.Context, seriesID int32) ([]ListSeriesRow, error) {
	rows, err := q.db.QueryContext(ctx, listSeries, seriesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListSeriesRow
	for rows.Next() {
		var i ListSeriesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Notes,
			&i.CreatedAt,
			&i.UserID,
			&i.VolumeCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSeries = `-- name: UpdateSeries :exec
update series
set name = $1, notes = $2
where id = $3 and user_id = $4
`

type UpdateSeriesParams struct {
	Name     string
	Notes    string
	SeriesID int32
	UserID   int32
}

func (q *Queries) UpdateSeries(ctx context.Context, arg UpdateSeriesParams) error {
	_, err := q.db.ExecContext(ctx, updateSeries,
		arg.Name,
		arg.Notes,
		arg.SeriesID,
		arg.UserID,
	)
	return err
}
