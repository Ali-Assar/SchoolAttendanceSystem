// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	// Queries for Admin
	CreateAdmin(ctx context.Context, arg CreateAdminParams) (string, error)
	// Queries for Entrance
	CreateEntrance(ctx context.Context, arg CreateEntranceParams) (int64, error)
	// Queries for Exit
	CreateExit(ctx context.Context, arg CreateExitParams) (int64, error)
	// Queries for Students
	CreateStudent(ctx context.Context, arg CreateStudentParams) (int64, error)
	// Queries for Teachers
	CreateTeacher(ctx context.Context, arg CreateTeacherParams) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	DeleteAdmin(ctx context.Context, userName string) error
	DeleteEntrance(ctx context.Context, id int64) error
	DeleteExit(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, userID int64) error
	GetAdminByUserName(ctx context.Context, userName string) (Admin, error)
	GetEntrancesByUserID(ctx context.Context, userID int64) ([]Entrance, error)
	GetExitsByUserID(ctx context.Context, userID int64) ([]Exit, error)
	GetStudentByID(ctx context.Context, userID int64) (GetStudentByIDRow, error)
	GetTeacherByID(ctx context.Context, userID int64) (GetTeacherByIDRow, error)
	// Queries for Time Range
	GetTimeRange(ctx context.Context, arg GetTimeRangeParams) ([]GetTimeRangeRow, error)
	GetTimeRangeByUserID(ctx context.Context, arg GetTimeRangeByUserIDParams) ([]GetTimeRangeByUserIDRow, error)
	GetUserByID(ctx context.Context, userID int64) (User, error)
	GetUserByName(ctx context.Context, arg GetUserByNameParams) ([]GetUserByNameRow, error)
	UpdateAdmin(ctx context.Context, arg UpdateAdminParams) error
	UpdateEntrance(ctx context.Context, arg UpdateEntranceParams) error
	UpdateExit(ctx context.Context, arg UpdateExitParams) error
	UpdateStudentAllowedTime(ctx context.Context, arg UpdateStudentAllowedTimeParams) error
	UpdateTeacherAllowedTime(ctx context.Context, arg UpdateTeacherAllowedTimeParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	UpdateUserBiometric(ctx context.Context, arg UpdateUserBiometricParams) error
}

var _ Querier = (*Queries)(nil)
