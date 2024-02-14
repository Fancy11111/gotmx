-- name: GetProjects :many
select * from projects 
order by name;

-- name: GetProjectById :one
select * from projects
where id = @id limit 1;

-- name: FindProjectsByName :many
select * from projects
where name like ('%' || @name || '%');

-- name: InsertProject :one
insert into projects(name) values (@name) returning *;

-- name: UpdateProject :one
update projects set name = @name where id = @id returning *; 

-- name: DeleteProject :exec
delete from projects where id = @id;
