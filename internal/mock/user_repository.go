// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"database/sql"
	"github.com/kl09/users"
	"github.com/oklog/ulid/v2"
	"sync"
	"time"
)

// Ensure, that UserRepositoryMock does implement users.UserRepository.
// If this is not the case, regenerate this file with moq.
var _ users.UserRepository = &UserRepositoryMock{}

// UserRepositoryMock is a mock implementation of users.UserRepository.
//
//     func TestSomethingThatUsesUserRepository(t *testing.T) {
//
//         // make and configure a mocked users.UserRepository
//         mockedUserRepository := &UserRepositoryMock{
//             ByEmailFunc: func(ctx context.Context, email string) (users.User, error) {
// 	               panic("mock out the ByEmail method")
//             },
//             ByIDFunc: func(ctx context.Context, id ulid.ULID) (users.User, error) {
// 	               panic("mock out the ByID method")
//             },
//             CreateFunc: func(ctx context.Context, tx *sql.Tx, u users.User) error {
// 	               panic("mock out the Create method")
//             },
//             DeleteFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID) error {
// 	               panic("mock out the Delete method")
//             },
//             LockFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID) (users.User, error) {
// 	               panic("mock out the Lock method")
//             },
//             SearchFunc: func(ctx context.Context, country string, since time.Time, until time.Time, limit int, orderASC bool) ([]users.User, error) {
// 	               panic("mock out the Search method")
//             },
//             TransactFunc: func(ctx context.Context, atomic func(*sql.Tx) error) error {
// 	               panic("mock out the Transact method")
//             },
//             UpdateFunc: func(ctx context.Context, tx *sql.Tx, id ulid.ULID, u users.User) error {
// 	               panic("mock out the Update method")
//             },
//         }
//
//         // use mockedUserRepository in code that requires users.UserRepository
//         // and then make assertions.
//
//     }
type UserRepositoryMock struct {
	// ByEmailFunc mocks the ByEmail method.
	ByEmailFunc func(ctx context.Context, email string) (users.User, error)

	// ByIDFunc mocks the ByID method.
	ByIDFunc func(ctx context.Context, id ulid.ULID) (users.User, error)

	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, tx *sql.Tx, u users.User) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, tx *sql.Tx, id ulid.ULID) error

	// LockFunc mocks the Lock method.
	LockFunc func(ctx context.Context, tx *sql.Tx, id ulid.ULID) (users.User, error)

	// SearchFunc mocks the Search method.
	SearchFunc func(ctx context.Context, country string, since time.Time, until time.Time, limit int, orderASC bool) ([]users.User, error)

	// TransactFunc mocks the Transact method.
	TransactFunc func(ctx context.Context, atomic func(*sql.Tx) error) error

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, tx *sql.Tx, id ulid.ULID, u users.User) error

	// calls tracks calls to the methods.
	calls struct {
		// ByEmail holds details about calls to the ByEmail method.
		ByEmail []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Email is the email argument value.
			Email string
		}
		// ByID holds details about calls to the ByID method.
		ByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID ulid.ULID
		}
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx *sql.Tx
			// U is the u argument value.
			U users.User
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx *sql.Tx
			// ID is the id argument value.
			ID ulid.ULID
		}
		// Lock holds details about calls to the Lock method.
		Lock []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx *sql.Tx
			// ID is the id argument value.
			ID ulid.ULID
		}
		// Search holds details about calls to the Search method.
		Search []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Country is the country argument value.
			Country string
			// Since is the since argument value.
			Since time.Time
			// Until is the until argument value.
			Until time.Time
			// Limit is the limit argument value.
			Limit int
			// OrderASC is the orderASC argument value.
			OrderASC bool
		}
		// Transact holds details about calls to the Transact method.
		Transact []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Atomic is the atomic argument value.
			Atomic func(*sql.Tx) error
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tx is the tx argument value.
			Tx *sql.Tx
			// ID is the id argument value.
			ID ulid.ULID
			// U is the u argument value.
			U users.User
		}
	}
	lockByEmail  sync.RWMutex
	lockByID     sync.RWMutex
	lockCreate   sync.RWMutex
	lockDelete   sync.RWMutex
	lockLock     sync.RWMutex
	lockSearch   sync.RWMutex
	lockTransact sync.RWMutex
	lockUpdate   sync.RWMutex
}

// ByEmail calls ByEmailFunc.
func (mock *UserRepositoryMock) ByEmail(ctx context.Context, email string) (users.User, error) {
	callInfo := struct {
		Ctx   context.Context
		Email string
	}{
		Ctx:   ctx,
		Email: email,
	}
	mock.lockByEmail.Lock()
	mock.calls.ByEmail = append(mock.calls.ByEmail, callInfo)
	mock.lockByEmail.Unlock()
	if mock.ByEmailFunc == nil {
		var (
			out1 users.User
			out2 error
		)
		return out1, out2
	}
	return mock.ByEmailFunc(ctx, email)
}

// ByEmailCalls gets all the calls that were made to ByEmail.
// Check the length with:
//     len(mockedUserRepository.ByEmailCalls())
func (mock *UserRepositoryMock) ByEmailCalls() []struct {
	Ctx   context.Context
	Email string
} {
	var calls []struct {
		Ctx   context.Context
		Email string
	}
	mock.lockByEmail.RLock()
	calls = mock.calls.ByEmail
	mock.lockByEmail.RUnlock()
	return calls
}

// ByID calls ByIDFunc.
func (mock *UserRepositoryMock) ByID(ctx context.Context, id ulid.ULID) (users.User, error) {
	callInfo := struct {
		Ctx context.Context
		ID  ulid.ULID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockByID.Lock()
	mock.calls.ByID = append(mock.calls.ByID, callInfo)
	mock.lockByID.Unlock()
	if mock.ByIDFunc == nil {
		var (
			out1 users.User
			out2 error
		)
		return out1, out2
	}
	return mock.ByIDFunc(ctx, id)
}

// ByIDCalls gets all the calls that were made to ByID.
// Check the length with:
//     len(mockedUserRepository.ByIDCalls())
func (mock *UserRepositoryMock) ByIDCalls() []struct {
	Ctx context.Context
	ID  ulid.ULID
} {
	var calls []struct {
		Ctx context.Context
		ID  ulid.ULID
	}
	mock.lockByID.RLock()
	calls = mock.calls.ByID
	mock.lockByID.RUnlock()
	return calls
}

// Create calls CreateFunc.
func (mock *UserRepositoryMock) Create(ctx context.Context, tx *sql.Tx, u users.User) error {
	callInfo := struct {
		Ctx context.Context
		Tx  *sql.Tx
		U   users.User
	}{
		Ctx: ctx,
		Tx:  tx,
		U:   u,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	if mock.CreateFunc == nil {
		var (
			out1 error
		)
		return out1
	}
	return mock.CreateFunc(ctx, tx, u)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedUserRepository.CreateCalls())
func (mock *UserRepositoryMock) CreateCalls() []struct {
	Ctx context.Context
	Tx  *sql.Tx
	U   users.User
} {
	var calls []struct {
		Ctx context.Context
		Tx  *sql.Tx
		U   users.User
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *UserRepositoryMock) Delete(ctx context.Context, tx *sql.Tx, id ulid.ULID) error {
	callInfo := struct {
		Ctx context.Context
		Tx  *sql.Tx
		ID  ulid.ULID
	}{
		Ctx: ctx,
		Tx:  tx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	if mock.DeleteFunc == nil {
		var (
			out1 error
		)
		return out1
	}
	return mock.DeleteFunc(ctx, tx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedUserRepository.DeleteCalls())
func (mock *UserRepositoryMock) DeleteCalls() []struct {
	Ctx context.Context
	Tx  *sql.Tx
	ID  ulid.ULID
} {
	var calls []struct {
		Ctx context.Context
		Tx  *sql.Tx
		ID  ulid.ULID
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Lock calls LockFunc.
func (mock *UserRepositoryMock) Lock(ctx context.Context, tx *sql.Tx, id ulid.ULID) (users.User, error) {
	callInfo := struct {
		Ctx context.Context
		Tx  *sql.Tx
		ID  ulid.ULID
	}{
		Ctx: ctx,
		Tx:  tx,
		ID:  id,
	}
	mock.lockLock.Lock()
	mock.calls.Lock = append(mock.calls.Lock, callInfo)
	mock.lockLock.Unlock()
	if mock.LockFunc == nil {
		var (
			out1 users.User
			out2 error
		)
		return out1, out2
	}
	return mock.LockFunc(ctx, tx, id)
}

// LockCalls gets all the calls that were made to Lock.
// Check the length with:
//     len(mockedUserRepository.LockCalls())
func (mock *UserRepositoryMock) LockCalls() []struct {
	Ctx context.Context
	Tx  *sql.Tx
	ID  ulid.ULID
} {
	var calls []struct {
		Ctx context.Context
		Tx  *sql.Tx
		ID  ulid.ULID
	}
	mock.lockLock.RLock()
	calls = mock.calls.Lock
	mock.lockLock.RUnlock()
	return calls
}

// Search calls SearchFunc.
func (mock *UserRepositoryMock) Search(ctx context.Context, country string, since time.Time, until time.Time, limit int, orderASC bool) ([]users.User, error) {
	callInfo := struct {
		Ctx      context.Context
		Country  string
		Since    time.Time
		Until    time.Time
		Limit    int
		OrderASC bool
	}{
		Ctx:      ctx,
		Country:  country,
		Since:    since,
		Until:    until,
		Limit:    limit,
		OrderASC: orderASC,
	}
	mock.lockSearch.Lock()
	mock.calls.Search = append(mock.calls.Search, callInfo)
	mock.lockSearch.Unlock()
	if mock.SearchFunc == nil {
		var (
			out1 []users.User
			out2 error
		)
		return out1, out2
	}
	return mock.SearchFunc(ctx, country, since, until, limit, orderASC)
}

// SearchCalls gets all the calls that were made to Search.
// Check the length with:
//     len(mockedUserRepository.SearchCalls())
func (mock *UserRepositoryMock) SearchCalls() []struct {
	Ctx      context.Context
	Country  string
	Since    time.Time
	Until    time.Time
	Limit    int
	OrderASC bool
} {
	var calls []struct {
		Ctx      context.Context
		Country  string
		Since    time.Time
		Until    time.Time
		Limit    int
		OrderASC bool
	}
	mock.lockSearch.RLock()
	calls = mock.calls.Search
	mock.lockSearch.RUnlock()
	return calls
}

// Transact calls TransactFunc.
func (mock *UserRepositoryMock) Transact(ctx context.Context, atomic func(*sql.Tx) error) error {
	callInfo := struct {
		Ctx    context.Context
		Atomic func(*sql.Tx) error
	}{
		Ctx:    ctx,
		Atomic: atomic,
	}
	mock.lockTransact.Lock()
	mock.calls.Transact = append(mock.calls.Transact, callInfo)
	mock.lockTransact.Unlock()
	if mock.TransactFunc == nil {
		var (
			out1 error
		)
		return out1
	}
	return mock.TransactFunc(ctx, atomic)
}

// TransactCalls gets all the calls that were made to Transact.
// Check the length with:
//     len(mockedUserRepository.TransactCalls())
func (mock *UserRepositoryMock) TransactCalls() []struct {
	Ctx    context.Context
	Atomic func(*sql.Tx) error
} {
	var calls []struct {
		Ctx    context.Context
		Atomic func(*sql.Tx) error
	}
	mock.lockTransact.RLock()
	calls = mock.calls.Transact
	mock.lockTransact.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *UserRepositoryMock) Update(ctx context.Context, tx *sql.Tx, id ulid.ULID, u users.User) error {
	callInfo := struct {
		Ctx context.Context
		Tx  *sql.Tx
		ID  ulid.ULID
		U   users.User
	}{
		Ctx: ctx,
		Tx:  tx,
		ID:  id,
		U:   u,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	if mock.UpdateFunc == nil {
		var (
			out1 error
		)
		return out1
	}
	return mock.UpdateFunc(ctx, tx, id, u)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedUserRepository.UpdateCalls())
func (mock *UserRepositoryMock) UpdateCalls() []struct {
	Ctx context.Context
	Tx  *sql.Tx
	ID  ulid.ULID
	U   users.User
} {
	var calls []struct {
		Ctx context.Context
		Tx  *sql.Tx
		ID  ulid.ULID
		U   users.User
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
