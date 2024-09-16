// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateAdmin(ctx context.Context, arg CreateAdminParams) (string, error)
	CreateEntrance(ctx context.Context, arg CreateEntranceParams) (int64, error)
	CreateStudent(ctx context.Context, arg CreateStudentParams) (int64, error)
	CreateTeacher(ctx context.Context, arg CreateTeacherParams) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	DeleteAdmin(ctx context.Context, userName string) error
	DeleteAttendance(ctx context.Context, attendanceID int64) error
	DeleteUser(ctx context.Context, userID int64) error
	GetAbsentTeachersByDate(ctx context.Context, date int64) ([]GetAbsentTeachersByDateRow, error)
	GetAbsentUsersByDate(ctx context.Context, date int64) ([]GetAbsentUsersByDateRow, error)
	GetAdminByUserName(ctx context.Context, userName string) (Admin, error)
	GetAttendanceBetweenDates(ctx context.Context, arg GetAttendanceBetweenDatesParams) ([]GetAttendanceBetweenDatesRow, error)
	GetAttendanceByDate(ctx context.Context, date int64) ([]GetAttendanceByDateRow, error)
	GetAttendanceByUserID(ctx context.Context, userID int64) ([]Attendance, error)
	GetAttendanceByUserIDAndDate(ctx context.Context, arg GetAttendanceByUserIDAndDateParams) (Attendance, error)
	GetStudentAttendanceBetweenDates(ctx context.Context, arg GetStudentAttendanceBetweenDatesParams) ([]GetStudentAttendanceBetweenDatesRow, error)
	GetStudentAttendanceByDate(ctx context.Context, date int64) ([]GetStudentAttendanceByDateRow, error)
	GetStudentByID(ctx context.Context, userID int64) (GetStudentByIDRow, error)
	GetTeacherAttendanceBetweenDates(ctx context.Context, arg GetTeacherAttendanceBetweenDatesParams) ([]GetTeacherAttendanceBetweenDatesRow, error)
	GetTeacherAttendanceByDate(ctx context.Context, date int64) ([]GetTeacherAttendanceByDateRow, error)
	GetTeacherByID(ctx context.Context, userID int64) (GetTeacherByIDRow, error)
	GetUserByID(ctx context.Context, userID int64) (User, error)
	GetUserByName(ctx context.Context, arg GetUserByNameParams) ([]GetUserByNameRow, error)
	GetUsersWithFalseBiometric(ctx context.Context) ([]GetUsersWithFalseBiometricRow, error)
	GetUsersWithTrueBiometric(ctx context.Context) ([]GetUsersWithTrueBiometricRow, error)
	UpdateAdmin(ctx context.Context, arg UpdateAdminParams) error
	UpdateExit(ctx context.Context, arg UpdateExitParams) error
	UpdateStudentAllowedTime(ctx context.Context, arg UpdateStudentAllowedTimeParams) error
	UpdateTeacherAllowedTime(ctx context.Context, arg UpdateTeacherAllowedTimeParams) error
	UpdateUserBiometric(ctx context.Context, arg UpdateUserBiometricParams) error
	UpdateUserDetails(ctx context.Context, arg UpdateUserDetailsParams) error
}

var _ Querier = (*Queries)(nil)
