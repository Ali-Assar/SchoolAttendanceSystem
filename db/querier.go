// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateAttendance(ctx context.Context, arg CreateAttendanceParams) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	DeleteAttendance(ctx context.Context, attendanceID int64) error
	DeleteUser(ctx context.Context, userID int64) error
	GetAllUsers(ctx context.Context) ([]GetAllUsersRow, error)
	GetAllUsersAttendanceByDate(ctx context.Context, date int64) ([]GetAllUsersAttendanceByDateRow, error)
	GetAttendanceByUserIDAndDate(ctx context.Context, arg GetAttendanceByUserIDAndDateParams) (Attendance, error)
	GetUserByID(ctx context.Context, userID int64) (GetUserByIDRow, error)
	GetUserByName(ctx context.Context, arg GetUserByNameParams) (GetUserByNameRow, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber int64) (GetUserByPhoneNumberRow, error)
	UpdateAttendance(ctx context.Context, arg UpdateAttendanceParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
