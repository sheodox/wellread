create table if not exists series (
	id int primary key generated always as identity,
	name varchar(1000) not null,
	notes text default '',
	created_at timestamp not null
);

create table if not exists volumes (
	id int primary key generated always as identity,
	series_id int not null,
	name varchar(1000) not null,
	notes text default '',
	current_page int default 0,
	created_at timestamp not null,

	constraint fk_series
		foreign key(series_id)
			references series(id) on delete cascade
);

create index on volumes (series_id);

create table if not exists reading_history (
	id int primary key generated always as identity,
	volume_id int,
	current_page int,
	created_at timestamp not null,

	constraint fk_volume
		foreign key(volume_id)
			references volumes(id) on delete cascade
);

create index on reading_history (volume_id);
