// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createAdmin = `-- name: CreateAdmin :one
INSERT INTO admin (user_name, password) 
VALUES (?, ?)
RETURNING user_name
`

type CreateAdminParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (q *Queries) CreateAdmin(ctx context.Context, arg CreateAdminParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createAdmin, arg.UserName, arg.Password)
	var user_name string
	err := row.Scan(&user_name)
	return user_name, err
}

const createAttendance = `-- name: CreateAttendance :one
INSERT INTO attendance (user_id, date, entry_time, exit_time) 
VALUES (?, ?, ?, ?)
RETURNING attendance_id
`

type CreateAttendanceParams struct {
	UserID    int64         `json:"user_id"`
	Date      int64         `json:"date"`
	EntryTime sql.NullInt64 `json:"entry_time"`
	ExitTime  sql.NullInt64 `json:"exit_time"`
}

func (q *Queries) CreateAttendance(ctx context.Context, arg CreateAttendanceParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createAttendance,
		arg.UserID,
		arg.Date,
		arg.EntryTime,
		arg.ExitTime,
	)
	var attendance_id int64
	err := row.Scan(&attendance_id)
	return attendance_id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id) 
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING user_id
`

type CreateUserParams struct {
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	PhoneNumber       string         `json:"phone_number"`
	ImagePath         sql.NullString `json:"image_path"`
	IsTeacher         bool           `json:"is_teacher"`
	IsBiometricActive sql.NullBool   `json:"is_biometric_active"`
	FingerID          sql.NullString `json:"finger_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.ImagePath,
		arg.IsTeacher,
		arg.IsBiometricActive,
		arg.FingerID,
	)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}

const deleteAdmin = `-- name: DeleteAdmin :exec
DELETE FROM admin 
WHERE user_name = ?
`

func (q *Queries) DeleteAdmin(ctx context.Context, userName string) error {
	_, err := q.db.ExecContext(ctx, deleteAdmin, userName)
	return err
}

const deleteAttendance = `-- name: DeleteAttendance :exec
DELETE FROM attendance WHERE attendance_id = ?
`

func (q *Queries) DeleteAttendance(ctx context.Context, attendanceID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAttendance, attendanceID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getAbsentUsersUntil9AM = `-- name: GetAbsentUsersUntil9AM :many
SELECT 
    users.user_id, 
    users.first_name, 
    users.last_name, 
    users.phone_number
FROM 
    users
LEFT JOIN 
    attendance ON users.user_id = attendance.user_id AND attendance.date = ? 
WHERE 
    attendance.entry_time IS NULL OR attendance.entry_time > 32400
`

type GetAbsentUsersUntil9AMRow struct {
	UserID      int64  `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

func (q *Queries) GetAbsentUsersUntil9AM(ctx context.Context, date int64) ([]GetAbsentUsersUntil9AMRow, error) {
	rows, err := q.db.QueryContext(ctx, getAbsentUsersUntil9AM, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAbsentUsersUntil9AMRow
	for rows.Next() {
		var i GetAbsentUsersUntil9AMRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
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

const getAdminByUserName = `-- name: GetAdminByUserName :one
SELECT user_name, password
FROM admin
WHERE user_name = ?
`

func (q *Queries) GetAdminByUserName(ctx context.Context, userName string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdminByUserName, userName)
	var i Admin
	err := row.Scan(&i.UserName, &i.Password)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id 
FROM users
`

type GetAllUsersRow struct {
	UserID            int64          `json:"user_id"`
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	PhoneNumber       string         `json:"phone_number"`
	ImagePath         sql.NullString `json:"image_path"`
	IsTeacher         bool           `json:"is_teacher"`
	IsBiometricActive sql.NullBool   `json:"is_biometric_active"`
	FingerID          sql.NullString `json:"finger_id"`
}

func (q *Queries) GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersRow
	for rows.Next() {
		var i GetAllUsersRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.ImagePath,
			&i.IsTeacher,
			&i.IsBiometricActive,
			&i.FingerID,
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

const getAllUsersAttendanceByDate = `-- name: GetAllUsersAttendanceByDate :many
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
    attendance.date = ?
`

type GetAllUsersAttendanceByDateRow struct {
	AttendanceID int64         `json:"attendance_id"`
	UserID       int64         `json:"user_id"`
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	Date         int64         `json:"date"`
	EntryTime    sql.NullInt64 `json:"entry_time"`
	ExitTime     sql.NullInt64 `json:"exit_time"`
}

func (q *Queries) GetAllUsersAttendanceByDate(ctx context.Context, date int64) ([]GetAllUsersAttendanceByDateRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsersAttendanceByDate, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllUsersAttendanceByDateRow
	for rows.Next() {
		var i GetAllUsersAttendanceByDateRow
		if err := rows.Scan(
			&i.AttendanceID,
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.Date,
			&i.EntryTime,
			&i.ExitTime,
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

const getAttendanceByUserIDAndDate = `-- name: GetAttendanceByUserIDAndDate :one
SELECT attendance_id, user_id, date, entry_time, exit_time 
FROM attendance 
WHERE user_id = ? AND date = ?
`

type GetAttendanceByUserIDAndDateParams struct {
	UserID int64 `json:"user_id"`
	Date   int64 `json:"date"`
}

func (q *Queries) GetAttendanceByUserIDAndDate(ctx context.Context, arg GetAttendanceByUserIDAndDateParams) (Attendance, error) {
	row := q.db.QueryRowContext(ctx, getAttendanceByUserIDAndDate, arg.UserID, arg.Date)
	var i Attendance
	err := row.Scan(
		&i.AttendanceID,
		&i.UserID,
		&i.Date,
		&i.EntryTime,
		&i.ExitTime,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id
FROM users 
WHERE user_id = ?
`

type GetUserByIDRow struct {
	UserID            int64          `json:"user_id"`
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	PhoneNumber       string         `json:"phone_number"`
	ImagePath         sql.NullString `json:"image_path"`
	IsTeacher         bool           `json:"is_teacher"`
	IsBiometricActive sql.NullBool   `json:"is_biometric_active"`
	FingerID          sql.NullString `json:"finger_id"`
}

func (q *Queries) GetUserByID(ctx context.Context, userID int64) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var i GetUserByIDRow
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.ImagePath,
		&i.IsTeacher,
		&i.IsBiometricActive,
		&i.FingerID,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id
FROM users
WHERE first_name = ? AND last_name = ?
`

type GetUserByNameParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUserByNameRow struct {
	UserID            int64          `json:"user_id"`
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	PhoneNumber       string         `json:"phone_number"`
	ImagePath         sql.NullString `json:"image_path"`
	IsTeacher         bool           `json:"is_teacher"`
	IsBiometricActive sql.NullBool   `json:"is_biometric_active"`
	FingerID          sql.NullString `json:"finger_id"`
}

func (q *Queries) GetUserByName(ctx context.Context, arg GetUserByNameParams) (GetUserByNameRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, arg.FirstName, arg.LastName)
	var i GetUserByNameRow
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.ImagePath,
		&i.IsTeacher,
		&i.IsBiometricActive,
		&i.FingerID,
	)
	return i, err
}

const getUserByPhoneNumber = `-- name: GetUserByPhoneNumber :one
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id
FROM users
WHERE phone_number = ?
`

type GetUserByPhoneNumberRow struct {
	UserID            int64          `json:"user_id"`
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	PhoneNumber       string         `json:"phone_number"`
	ImagePath         sql.NullString `json:"image_path"`
	IsTeacher         bool           `json:"is_teacher"`
	IsBiometricActive sql.NullBool   `json:"is_biometric_active"`
	FingerID          sql.NullString `json:"finger_id"`
}

func (q *Queries) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (GetUserByPhoneNumberRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByPhoneNumber, phoneNumber)
	var i GetUserByPhoneNumberRow
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.ImagePath,
		&i.IsTeacher,
		&i.IsBiometricActive,
		&i.FingerID,
	)
	return i, err
}

const updateAdmin = `-- name: UpdateAdmin :exec
UPDATE admin
SET password = ?
WHERE user_name = ?
`

type UpdateAdminParams struct {
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

func (q *Queries) UpdateAdmin(ctx context.Context, arg UpdateAdminParams) error {
	_, err := q.db.ExecContext(ctx, updateAdmin, arg.Password, arg.UserName)
	return err
}

const updateAttendance = `-- name: UpdateAttendance :exec
UPDATE attendance 
SET entry_time = ?, exit_time = ? 
WHERE user_id = ? AND attendance_id = ? AND date = ?
`

type UpdateAttendanceParams struct {
	EntryTime    sql.NullInt64 `json:"entry_time"`
	ExitTime     sql.NullInt64 `json:"exit_time"`
	UserID       int64         `json:"user_id"`
	AttendanceID int64         `json:"attendance_id"`
	Date         int64         `json:"date"`
}

func (q *Queries) UpdateAttendance(ctx context.Context, arg UpdateAttendanceParams) error {
	_, err := q.db.ExecContext(ctx, updateAttendance,
		arg.EntryTime,
		arg.ExitTime,
		arg.UserID,
		arg.AttendanceID,
		arg.Date,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users 
SET first_name = ?, last_name = ?, phone_number = ?, image_path = ?, is_teacher = ?, is_biometric_active = ?, finger_id = ? 
WHERE user_id = ?
`

type UpdateUserParams struct {
	FirstName         string         `json:"first_name"`
	LastName          string         `json:"last_name"`
	PhoneNumber       string         `json:"phone_number"`
	ImagePath         sql.NullString `json:"image_path"`
	IsTeacher         bool           `json:"is_teacher"`
	IsBiometricActive sql.NullBool   `json:"is_biometric_active"`
	FingerID          sql.NullString `json:"finger_id"`
	UserID            int64          `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.ImagePath,
		arg.IsTeacher,
		arg.IsBiometricActive,
		arg.FingerID,
		arg.UserID,
	)
	return err
}
