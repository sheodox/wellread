-- user
alter table users
alter column provider_id set not null;

-- series
alter table series
alter column user_id set not null;

alter table series
alter column notes set not null;

-- volumes
alter table volumes
alter column user_id set not null;

alter table volumes
alter column status set not null;

alter table volumes
alter column notes set not null;

alter table volumes
alter column current_page set not null;

-- reading history
alter table reading_history
alter column user_id set not null;

alter table reading_history
alter column volume_id set not null;

alter table reading_history
alter column current_page set not null;

alter table reading_history
alter column pages_read set not null;
