-- name: CreateCategory :exec
insert into categories (id, name)
values (?,?);

-- name: CreateCourse :exec
insert into courses (id, name, category_id)
values (?,?,?);