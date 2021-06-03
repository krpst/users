package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kl09/users"
	"github.com/kl09/users/internal/mock"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

var (
	ulidExample = ulid.MustParse("01AN4Z07BY79KA1307SR9X4MV3")
	userExample = users.User{
		ID:           ulidExample,
		PasswordHash: "some_hash",
		Password:     "some_pswd",
		Email:        "example@example.org",
		FirstName:    "John",
		LastName:     "Smith",
		Country:      "UK",
		CreatedAt:    time.Date(2020, time.April, 15, 0, 0, 0, 0, time.UTC),
		UpdatedAt:    time.Date(2020, time.April, 15, 0, 0, 0, 0, time.UTC),
		Deleted:      false,
	}
)

func TestUserService_ByID(t *testing.T) {
	now := time.Date(2020, time.April, 16, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		testName      string
		userService   users.UserRepository
		expectedUser  users.User
		expectedError error
	}{
		{
			testName: "success get user id",
			userService: &mock.UserRepositoryMock{
				ByIDFunc: func(ctx context.Context, id ulid.ULID) (users.User, error) {
					return userExample, nil
				},
			},
			expectedUser: userExample,
		},
		{
			testName: "error user not found",
			userService: &mock.UserRepositoryMock{
				ByIDFunc: func(ctx context.Context, id ulid.ULID) (users.User, error) {
					return users.User{}, users.NewError(users.ErrCodeUserNotFound, "User not found")
				},
			},
			expectedError: fmt.Errorf(
				"error getting user id: %s, err: %w",
				ulidExample.String(),
				users.NewError(users.ErrCodeUserNotFound, "User not found"),
			),
		},
		{
			testName: "error - some error from repository",
			userService: &mock.UserRepositoryMock{
				ByIDFunc: func(ctx context.Context, id ulid.ULID) (users.User, error) {
					return users.User{}, fmt.Errorf("some error")
				},
			},
			expectedError: fmt.Errorf(
				"error getting user id: %s, err: %w",
				ulidExample.String(),
				fmt.Errorf("some error"),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			s := NewUserService(
				tc.userService,
				nil,
				func() time.Time { return now },
				nil,
			)

			u, err := s.ByID(context.Background(), ulidExample)
			if diff := cmp.Diff(tc.expectedUser, u); diff != "" {
				t.Fatal(diff)
			}
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestUserService_Update(t *testing.T) {
	now := time.Date(2020, time.April, 16, 0, 0, 0, 0, time.UTC)
	userUpdated := userExample
	userUpdated.UpdatedAt = now

	testCases := []struct {
		testName      string
		userService   users.UserRepository
		userNotify    users.UserNotifier
		expectedUser  users.User
		expectedError error
	}{
		{
			testName: "success update user id",
			userService: &mock.UserRepositoryMock{
				TransactFunc: func(ctx context.Context, atomic func(*sql.Tx) error) error {
					return atomic(nil)
				},
				LockFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID) (users.User, error) {
					return userExample, nil
				},
				UpdateFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID, u users.User) error {
					return nil
				},
			},
			userNotify: &mock.UserNotifierMock{
				UserUpdatedFunc: func(ctx context.Context, u users.User) error {
					return nil
				},
			},
			expectedUser: userUpdated,
		},
		{
			testName: "error from notifier",
			userService: &mock.UserRepositoryMock{
				TransactFunc: func(ctx context.Context, atomic func(*sql.Tx) error) error {
					return atomic(nil)
				},
				LockFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID) (users.User, error) {
					return userExample, nil
				},
				UpdateFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID, u users.User) error {
					return nil
				},
			},
			userNotify: &mock.UserNotifierMock{
				UserUpdatedFunc: func(ctx context.Context, u users.User) error {
					return errors.New("some error")
				},
			},
			expectedUser: userUpdated,
			expectedError: fmt.Errorf(
				"error sending notification about updating user id: %s, err: %w",
				ulidExample.String(),
				fmt.Errorf("some error"),
			),
		},
		{
			testName: "error from repository",
			userService: &mock.UserRepositoryMock{
				TransactFunc: func(ctx context.Context, atomic func(*sql.Tx) error) error {
					return atomic(nil)
				},
				LockFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID) (users.User, error) {
					return userExample, nil
				},
				UpdateFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID, u users.User) error {
					return errors.New("some error")
				},
			},
			expectedUser: userUpdated,
			expectedError: fmt.Errorf(
				"error updating user id: %s, err: %w",
				ulidExample.String(),
				fmt.Errorf("some error"),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			s := NewUserService(
				tc.userService,
				tc.userNotify,
				func() time.Time { return now },
				nil,
			)

			u, err := s.Update(context.Background(), ulidExample, userExample)
			if diff := cmp.Diff(tc.expectedUser, u); diff != "" {
				t.Fatal(diff)
			}
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

// TODO: add more tests for all services functions
