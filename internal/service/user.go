package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kl09/users"
	"github.com/oklog/ulid/v2"
)

// UserService is an implementation of users.UserService.
type UserService struct {
	usersRepo     users.UserRepository
	usersNotifier users.UserNotifier
	ulidGenerator func() (ulid.ULID, error)
	nowFn         func() time.Time
}

// NewUserService creates a new instance of users.UserService.
func NewUserService(
	usersRepo users.UserRepository,
	usersNotifier users.UserNotifier,
	nowFn func() time.Time,
	ulidGenerator func() (ulid.ULID, error),
) *UserService {
	s := UserService{
		usersRepo:     usersRepo,
		usersNotifier: usersNotifier,
		nowFn:         nowFn,
		ulidGenerator: ulidGenerator,
	}

	return &s
}

// Create creates a new users.User.
func (s *UserService) Create(ctx context.Context, u users.User) (users.User, error) {
	var err error

	_, err = s.usersRepo.ByEmail(ctx, u.Email)
	if err == nil {
		return users.User{}, users.NewError(users.ErrCodeEmailExists, "User with this email already exists.")
	}

	if users.ErrorCode(err) != users.ErrCodeUserNotFound {
		return users.User{}, fmt.Errorf("error checking existed user: %s", err)
	}

	u.PasswordHash, err = hashAndSalt(u.Password)
	if err != nil {
		return users.User{}, fmt.Errorf("error hashing user password: %s", err)
	}

	id, err := s.ulidGenerator()
	if err != nil {
		return users.User{}, fmt.Errorf("error generating ulid: %s", err)
	}

	u.ID = id
	u.CreatedAt = s.nowFn()
	u.UpdatedAt = s.nowFn()

	err = s.usersRepo.Transact(ctx, func(tx *sql.Tx) error {
		errTx := s.usersRepo.Create(ctx, tx, u)
		if errTx != nil {
			return fmt.Errorf("error creating user: %s", errTx)
		}

		errTx = s.usersNotifier.UserCreated(ctx, u)
		if errTx != nil {
			return fmt.Errorf("error sending notification about creating user id: %s, err: %w", id.String(), errTx)
		}

		return nil
	})

	return u, err
}

// Update updates User identified by u.ID with the data provided on u and generates an event UserUpdated.
func (s *UserService) Update(ctx context.Context, id ulid.ULID, u users.User) (users.User, error) {
	// TODO: make sure that user wants to update himself: add tokens logic.

	err := s.usersRepo.Transact(ctx, func(tx *sql.Tx) error {
		_, lockErr := s.usersRepo.Lock(ctx, tx, id)
		if lockErr != nil {
			return fmt.Errorf("lock record: %w", lockErr)
		}

		u.UpdatedAt = s.nowFn()

		errTx := s.usersRepo.Update(ctx, tx, id, u)
		if errTx != nil {
			return fmt.Errorf("error updating user id: %s, err: %w", id.String(), errTx)
		}

		errTx = s.usersNotifier.UserUpdated(ctx, u)
		if errTx != nil {
			return fmt.Errorf("error sending notification about updating user id: %s, err: %w", id.String(), errTx)
		}

		return nil
	})

	return u, err
}

// ByID retrieves User from repository by id.
func (s *UserService) ByID(ctx context.Context, id ulid.ULID) (users.User, error) {
	u, err := s.usersRepo.ByID(ctx, id)
	if err != nil {
		return users.User{}, fmt.Errorf("error getting user id: %s, err: %w", id.String(), err)
	}

	return u, nil
}

// Delete makes a soft User delete and generates an event UserDeleted.
func (s *UserService) Delete(ctx context.Context, id ulid.ULID) error {
	err := s.usersRepo.Transact(ctx, func(tx *sql.Tx) error {
		_, lockErr := s.usersRepo.Lock(ctx, tx, id)
		if lockErr != nil {
			return fmt.Errorf("lock record: %w", lockErr)
		}

		errTx := s.usersRepo.Delete(ctx, tx, id)
		if errTx != nil {
			return fmt.Errorf("error deleting user id: %s, err: %w", id.String(), errTx)
		}

		errTx = s.usersNotifier.UserDeleted(ctx, id)
		if errTx != nil {
			return fmt.Errorf("error sending notification about deleting user id: %s, err: %w", id.String(), errTx)
		}

		return nil
	})

	return err
}

// Search retrieves User from repository with pagination filtering by certain fields.
func (s *UserService) Search(
	ctx context.Context,
	country string,
	since, until time.Time,
	limit int,
	orderASC bool,
) ([]users.User, error) {
	return s.usersRepo.Search(ctx, country, since, until, limit, orderASC)
}
