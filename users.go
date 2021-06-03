package users

import (
	"context"
	"database/sql"
	"time"

	"github.com/oklog/ulid/v2"
)

// User is a user's data.
type User struct {
	// ID is a User ID, uniquely identifies an User.
	ID ulid.ULID
	// PasswordHash is a hash of user password.
	PasswordHash string
	Password     string
	Email        string
	FirstName    string
	LastName     string
	Country      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	// Deleted is a soft delete.
	Deleted bool
}

// ForLog returns User struct for logging.
func (u User) ForLog() User {
	uCopied := u

	// we shouldn't store passwords into logs.
	u.PasswordHash = "***"
	u.Password = "***"

	// we shouldn't store personal data into logs.
	u.FirstName = "***"
	u.LastName = "***"

	return uCopied
}

// Storage provides transaction support with "Read committed" isolation level.
//
// Postgres provides "Read committed" isolation level by default - https://www.postgresql.org/docs/9.5/transaction-iso.html.
// Note, sometimes "Read committed" isolation level is not enough to provide sufficient consistency for an application,
// see this excellent read on consistency in databases - https://www.cockroachlabs.com/blog/consistency-model/.
//
//go:generate moq -stub -out internal/mock/dbstorage.go -pkg mock . Storage
type Storage interface {
	// Transact runs function atomic in a transaction.
	// If the function returns an error, transaction will be rolled back, otherwise transaction is committed.
	Transact(ctx context.Context, atomic func(*sql.Tx) error) error
}

// UserRepository is a storage for users.
//
//go:generate moq -stub -out internal/mock/user_repository.go -pkg mock . UserRepository
type UserRepository interface {
	Storage

	// Lock acquires a lock associated with user id.
	// Acquired lock is released upon tx commit / roll back.
	//
	// Lock is a blocking method.
	//
	// Since Storage provides only "Read committed" isolation level every User modification (create, update, delete)
	// through UserRepository must be coordinated via Lock method.
	Lock(ctx context.Context, tx *sql.Tx, id ulid.ULID) (User, error)

	// Create creates a new User.
	Create(ctx context.Context, tx *sql.Tx, u User) error
	// Update updates User.
	Update(ctx context.Context, tx *sql.Tx, id ulid.ULID, u User) error
	// Delete makes a soft User delete.
	Delete(ctx context.Context, tx *sql.Tx, id ulid.ULID) error

	// ByID retrieves User from repository by id.
	ByID(ctx context.Context, id ulid.ULID) (User, error)
	// ByEmail retrieves a User by email.
	ByEmail(ctx context.Context, email string) (User, error)

	// Search retrieves User from repository with pagination filtering by certain fields.
	Search(
		ctx context.Context,
		country string,
		since, until time.Time,
		limit int,
		orderASC bool,
	) ([]User, error)
}

// UserService aggregates some logic on managing users.
//
//go:generate moq -stub -out internal/mock/user_service.go -pkg mock . UserService
type UserService interface {
	// Create creates new User, persists it in UserRepository and generates an event UserCreated.
	//
	// It returns User enriched with some data (such as ID).
	Create(ctx context.Context, u User) (User, error)
	// Update updates User identified by u.ID with the data provided on u and generates an event UserUpdated.
	//
	// It returns existing User enriched with the data provided on u.
	Update(ctx context.Context, id ulid.ULID, u User) (User, error)
	// ByID retrieves User from repository by id.
	ByID(ctx context.Context, id ulid.ULID) (User, error)
	// Delete makes a soft User delete and generates an event UserDeleted.
	Delete(ctx context.Context, id ulid.ULID) error

	// Search retrieves User from repository with pagination filtering by certain fields.
	Search(
		ctx context.Context,
		country string,
		since, until time.Time,
		limit int,
		orderASC bool,
	) ([]User, error)
}

// UserNotifier notifies all the interested parties about User changes.
//
//go:generate moq -stub -out internal/mock/user_notifier.go -pkg mock . UserNotifier
type UserNotifier interface {
	// UserCreated notifies all the interested parties that a User was created.
	UserCreated(ctx context.Context, u User) error
	// UserUpdated notifies all the interested parties that a User was updated.
	UserUpdated(ctx context.Context, u User) error
	// UserDeleted notifies all the interested parties that a User was delete.
	UserDeleted(ctx context.Context, id ulid.ULID) error
}
