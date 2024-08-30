-- name: CreateUser :exec
INSERT INTO users (first_name, last_name, phone_number, image_path, role_id, is_admin) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetUserByID :one
SELECT user_id, first_name, last_name, phone_number, image_path, role_id, is_admin 
FROM users 
WHERE user_id = ?;

-- name: UpdateUser :exec
UPDATE users 
SET first_name = ?, last_name = ?, phone_number = ?, image_path = ?, role_id = ?, is_admin = ? 
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = ?;

-- Roles Queries

-- name: CreateRole :exec
INSERT INTO roles (role_name) 
VALUES (?);

-- name: GetRoleByID :one
SELECT role_id, role_name 
FROM roles 
WHERE role_id = ?;

-- name: UpdateRole :exec
UPDATE roles 
SET role_name = ? 
WHERE role_id = ?;

-- name: DeleteRole :exec
DELETE FROM roles WHERE role_id = ?;

-- name: CreateAttendance :exec
INSERT INTO attendance (user_id, date, entry_time, exit_time) 
VALUES (?, ?, ?, ?);

-- name: GetAttendanceByUserIDAndDate :one
SELECT attendance_id, user_id, date, entry_time, exit_time 
FROM attendance 
WHERE user_id = ? AND date = ?;

-- name: UpdateAttendance :exec
UPDATE attendance 
SET entry_time = ?, exit_time = ? 
WHERE user_id = ? AND date = ?;

-- name: DeleteAttendance :exec
DELETE FROM attendance WHERE attendance_id = ?;

-- name: GetAllUsersAttendanceByDate :many
SELECT 
    attendance.attendance_id, 
    attendance.user_id, 
    users.first_name, 
    users.last_name, 
    attendance.date, 
    attendance.entry_time, 
    attendance.exit_time
FROM 
    attendance
INNER JOIN 
    users ON attendance.user_id = users.user_id
WHERE 
    attendance.date = ?;

-- name: GetUserAttendanceBetweenDates :many
SELECT 
    attendance.attendance_id, 
    attendance.user_id, 
    users.first_name, 
    users.last_name, 
    attendance.date, 
    attendance.entry_time, 
    attendance.exit_time
FROM 
    attendance
INNER JOIN 
    users ON attendance.user_id = users.user_id
WHERE 
    attendance.user_id = ? 
AND 
    attendance.date BETWEEN ? AND ?;