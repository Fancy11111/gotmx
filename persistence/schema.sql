create table if not exists projects (
    id INTEGER PRIMARY KEY,
    name varchar(64) not null
);

create table if not exists tasks (
    id INTEGER PRIMARY KEY,
    name varchar(64) not null,
    project_id INTEGER references projects(id) on delete set null
);

create table if not exists timings (
    id INTEGER PRIMARY KEY,
    start timestamp not null,
    stop timestamp,
    task_id INTEGER references task(id) on delete cascade
);

create table if not exists tags (
    id INTEGER PRIMARY KEY,
    name varchar(64) not null,
    color integer
);

create table if not exists task_tag (
    tag_id integer not null references tags(id) on delete set null,
    task_id integer not null references tasks(id) on delete set null,
    primary key(tag_id, task_id)
);
