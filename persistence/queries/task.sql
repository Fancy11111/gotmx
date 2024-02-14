-- name: GetTasks :many
select * from tasks 
order by name;

-- name: GetTaskById :one
select * from tasks
where id = @id limit 1;

-- name: FindTasksByName :many
select * from tasks
where name like ('%' || @name || '%');

-- name: InsertTask :one
insert into tasks(name, project_id) values(?,?) returning *;
