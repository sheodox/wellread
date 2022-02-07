alter table reading_history
add column pages_read int default 0;

alter table volumes
add column status text default 'planning';
