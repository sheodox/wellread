-- user
alter table users
alter column provider_id drop not null;

-- series
alter table series
alter column user_id drop not null;

alter table series
alter column notes drop not null;

-- volumes
alter table volumes
alter column user_id drop not null;

alter table volumes
alter column status drop not null;

alter table volumes
alter column notes drop not null;

alter table volumes
alter column current_page drop not null;

-- reading history
alter table reading_history
alter column user_id drop not null;

alter table reading_history
alter column volume_id drop not null;

alter table reading_history
alter column current_page drop not null;

alter table reading_history
alter column pages_read drop not null;
