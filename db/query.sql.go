// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
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

const createEntrance = `-- name: CreateEntrance :one
INSERT INTO attendance (user_id, date, enter_time)
VALUES (?, ?, ?)
RETURNING attendance_id
`

type CreateEntranceParams struct {
	UserID    int64 `json:"user_id"`
	Date      int64 `json:"date"`
	EnterTime int64 `json:"enter_time"`
}

func (q *Queries) CreateEntrance(ctx context.Context, arg CreateEntranceParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createEntrance, arg.UserID, arg.Date, arg.EnterTime)
	var attendance_id int64
	err := row.Scan(&attendance_id)
	return attendance_id, err
}

const createStudent = `-- name: CreateStudent :one
INSERT INTO students (user_id, required_entry_time)
VALUES (?, ?)
RETURNING user_id
`

type CreateStudentParams struct {
	UserID            int64 `json:"user_id"`
	RequiredEntryTime int64 `json:"required_entry_time"`
}

func (q *Queries) CreateStudent(ctx context.Context, arg CreateStudentParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createStudent, arg.UserID, arg.RequiredEntryTime)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}

const createTeacher = `-- name: CreateTeacher :one
INSERT INTO teachers (user_id, sunday_entry_time, monday_entry_time, tuesday_entry_time, wednesday_entry_time, thursday_entry_time, friday_entry_time, saturday_entry_time)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING user_id
`

type CreateTeacherParams struct {
	UserID             int64 `json:"user_id"`
	SundayEntryTime    int64 `json:"sunday_entry_time"`
	MondayEntryTime    int64 `json:"monday_entry_time"`
	TuesdayEntryTime   int64 `json:"tuesday_entry_time"`
	WednesdayEntryTime int64 `json:"wednesday_entry_time"`
	ThursdayEntryTime  int64 `json:"thursday_entry_time"`
	FridayEntryTime    int64 `json:"friday_entry_time"`
	SaturdayEntryTime  int64 `json:"saturday_entry_time"`
}

func (q *Queries) CreateTeacher(ctx context.Context, arg CreateTeacherParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTeacher,
		arg.UserID,
		arg.SundayEntryTime,
		arg.MondayEntryTime,
		arg.TuesdayEntryTime,
		arg.WednesdayEntryTime,
		arg.ThursdayEntryTime,
		arg.FridayEntryTime,
		arg.SaturdayEntryTime,
	)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (first_name, last_name, phone_number, image_path, finger_id, is_biometric_active)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING user_id
`

type CreateUserParams struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	ImagePath         string `json:"image_path"`
	FingerID          string `json:"finger_id"`
	IsBiometricActive bool   `json:"is_biometric_active"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.ImagePath,
		arg.FingerID,
		arg.IsBiometricActive,
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
DELETE FROM attendance
WHERE attendance_id = ?
`

func (q *Queries) DeleteAttendance(ctx context.Context, attendanceID int64) error {
	_, err := q.db.ExecContext(ctx, deleteAttendance, attendanceID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
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

const getAttendanceByUserID = `-- name: GetAttendanceByUserID :many
SELECT attendance_id, user_id, date, enter_time, exit_time
FROM attendance
WHERE user_id = ?
`

func (q *Queries) GetAttendanceByUserID(ctx context.Context, userID int64) ([]Attendance, error) {
	rows, err := q.db.QueryContext(ctx, getAttendanceByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Attendance
	for rows.Next() {
		var i Attendance
		if err := rows.Scan(
			&i.AttendanceID,
			&i.UserID,
			&i.Date,
			&i.EnterTime,
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
SELECT attendance_id, user_id, date, enter_time, exit_time
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
		&i.EnterTime,
		&i.ExitTime,
	)
	return i, err
}

const getStudentByID = `-- name: GetStudentByID :one
SELECT s.user_id, u.first_name, u.last_name, s.required_entry_time
FROM students s
JOIN users u ON s.user_id = u.user_id
WHERE s.user_id = ?
`

type GetStudentByIDRow struct {
	UserID            int64  `json:"user_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	RequiredEntryTime int64  `json:"required_entry_time"`
}

func (q *Queries) GetStudentByID(ctx context.Context, userID int64) (GetStudentByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getStudentByID, userID)
	var i GetStudentByIDRow
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.RequiredEntryTime,
	)
	return i, err
}

const getTeacherByID = `-- name: GetTeacherByID :one
SELECT t.user_id, u.first_name, u.last_name, t.sunday_entry_time, t.monday_entry_time, t.tuesday_entry_time, t.wednesday_entry_time, t.thursday_entry_time, t.friday_entry_time, t.saturday_entry_time
FROM teachers t
JOIN users u ON t.user_id = u.user_id
WHERE t.user_id = ?
`

type GetTeacherByIDRow struct {
	UserID             int64  `json:"user_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	SundayEntryTime    int64  `json:"sunday_entry_time"`
	MondayEntryTime    int64  `json:"monday_entry_time"`
	TuesdayEntryTime   int64  `json:"tuesday_entry_time"`
	WednesdayEntryTime int64  `json:"wednesday_entry_time"`
	ThursdayEntryTime  int64  `json:"thursday_entry_time"`
	FridayEntryTime    int64  `json:"friday_entry_time"`
	SaturdayEntryTime  int64  `json:"saturday_entry_time"`
}

func (q *Queries) GetTeacherByID(ctx context.Context, userID int64) (GetTeacherByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getTeacherByID, userID)
	var i GetTeacherByIDRow
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.SundayEntryTime,
		&i.MondayEntryTime,
		&i.TuesdayEntryTime,
		&i.WednesdayEntryTime,
		&i.ThursdayEntryTime,
		&i.FridayEntryTime,
		&i.SaturdayEntryTime,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, first_name, last_name, phone_number, image_path, finger_id, is_biometric_active
FROM users
WHERE user_id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.ImagePath,
		&i.FingerID,
		&i.IsBiometricActive,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :many
SELECT user_id, first_name, last_name, phone_number, image_path, is_biometric_active
FROM users
WHERE first_name = ? AND last_name = ?
`

type GetUserByNameParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetUserByNameRow struct {
	UserID            int64  `json:"user_id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	PhoneNumber       string `json:"phone_number"`
	ImagePath         string `json:"image_path"`
	IsBiometricActive bool   `json:"is_biometric_active"`
}

func (q *Queries) GetUserByName(ctx context.Context, arg GetUserByNameParams) ([]GetUserByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserByName, arg.FirstName, arg.LastName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserByNameRow
	for rows.Next() {
		var i GetUserByNameRow
		if err := rows.Scan(
			&i.UserID,
			&i.FirstName,
			&i.LastName,
			&i.PhoneNumber,
			&i.ImagePath,
			&i.IsBiometricActive,
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

const updateExit = `-- name: UpdateExit :exec
UPDATE attendance
SET exit_time = ?
WHERE attendance_id = ?
`

type UpdateExitParams struct {
	ExitTime     int64 `json:"exit_time"`
	AttendanceID int64 `json:"attendance_id"`
}

func (q *Queries) UpdateExit(ctx context.Context, arg UpdateExitParams) error {
	_, err := q.db.ExecContext(ctx, updateExit, arg.ExitTime, arg.AttendanceID)
	return err
}

const updateStudentAllowedTime = `-- name: UpdateStudentAllowedTime :exec
UPDATE students
SET required_entry_time = ?
WHERE user_id = ?
`

type UpdateStudentAllowedTimeParams struct {
	RequiredEntryTime int64 `json:"required_entry_time"`
	UserID            int64 `json:"user_id"`
}

func (q *Queries) UpdateStudentAllowedTime(ctx context.Context, arg UpdateStudentAllowedTimeParams) error {
	_, err := q.db.ExecContext(ctx, updateStudentAllowedTime, arg.RequiredEntryTime, arg.UserID)
	return err
}

const updateTeacherAllowedTime = `-- name: UpdateTeacherAllowedTime :exec
UPDATE teachers
SET sunday_entry_time = ?, monday_entry_time = ?, tuesday_entry_time = ?, wednesday_entry_time = ?, thursday_entry_time = ?, friday_entry_time = ?, saturday_entry_time = ?
WHERE user_id = ?
`

type UpdateTeacherAllowedTimeParams struct {
	SundayEntryTime    int64 `json:"sunday_entry_time"`
	MondayEntryTime    int64 `json:"monday_entry_time"`
	TuesdayEntryTime   int64 `json:"tuesday_entry_time"`
	WednesdayEntryTime int64 `json:"wednesday_entry_time"`
	ThursdayEntryTime  int64 `json:"thursday_entry_time"`
	FridayEntryTime    int64 `json:"friday_entry_time"`
	SaturdayEntryTime  int64 `json:"saturday_entry_time"`
	UserID             int64 `json:"user_id"`
}

func (q *Queries) UpdateTeacherAllowedTime(ctx context.Context, arg UpdateTeacherAllowedTimeParams) error {
	_, err := q.db.ExecContext(ctx, updateTeacherAllowedTime,
		arg.SundayEntryTime,
		arg.MondayEntryTime,
		arg.TuesdayEntryTime,
		arg.WednesdayEntryTime,
		arg.ThursdayEntryTime,
		arg.FridayEntryTime,
		arg.SaturdayEntryTime,
		arg.UserID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET first_name = ?, last_name = ?, phone_number = ?, image_path = ?
WHERE user_id = ?
`

type UpdateUserParams struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	ImagePath   string `json:"image_path"`
	UserID      int64  `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.ImagePath,
		arg.UserID,
	)
	return err
}

const updateUserBiometric = `-- name: UpdateUserBiometric :exec
UPDATE users
SET image_path = ?, finger_id = ?, is_biometric_active = ?
WHERE user_id = ?
`

type UpdateUserBiometricParams struct {
	ImagePath         string `json:"image_path"`
	FingerID          string `json:"finger_id"`
	IsBiometricActive bool   `json:"is_biometric_active"`
	UserID            int64  `json:"user_id"`
}

func (q *Queries) UpdateUserBiometric(ctx context.Context, arg UpdateUserBiometricParams) error {
	_, err := q.db.ExecContext(ctx, updateUserBiometric,
		arg.ImagePath,
		arg.FingerID,
		arg.IsBiometricActive,
		arg.UserID,
	)
	return err
}
