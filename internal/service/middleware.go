package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kl09/users"
	"github.com/oklog/ulid/v2"
)

// NewLoggingMiddleware makes a logging middleware for users.UserService.
func NewLoggingMiddleware(l log.Logger, s users.UserService) users.UserService {
	return &loggingMW{
		logger: l,
		next:   s,
	}
}

type loggingMW struct {
	logger log.Logger
	next   users.UserService
}

func (mw *loggingMW) Create(ctx context.Context, u users.User) (createdUser users.User, err error) {
	defer func(begin time.Time) {
		level.Debug(mw.logger).Log(
			"method", "CreateUser",
			"err", err,
			"took", time.Since(begin),
			"user", u.ForLog,
			"createdUser", createdUser.ForLog,
		)
	}(time.Now().UTC())

	createdUser, err = mw.next.Create(ctx, u)
	return
}

func (mw *loggingMW) Update(ctx context.Context, id ulid.ULID, u users.User) (userUpdated users.User, err error) {
	defer func(begin time.Time) {
		level.Debug(mw.logger).Log(
			"method", "UpdateUser",
			"err", err,
			"took", time.Since(begin),
			"id", id.String(),
			"user", u.ForLog,
			"userUpdated", userUpdated.ForLog(),
		)
	}(time.Now().UTC())

	userUpdated, err = mw.next.Create(ctx, u)
	return
}

func (mw *loggingMW) ByID(ctx context.Context, id ulid.ULID) (u users.User, err error) {
	defer func(begin time.Time) {
		level.Debug(mw.logger).Log(
			"method", "UserByID",
			"err", err,
			"took", time.Since(begin),
			"id", id.String(),
			"user", u.ForLog,
		)
	}(time.Now().UTC())

	u, err = mw.next.Create(ctx, u)
	return
}

func (mw *loggingMW) Delete(ctx context.Context, id ulid.ULID) (err error) {
	defer func(begin time.Time) {
		level.Debug(mw.logger).Log(
			"method", "DeleteUser",
			"err", err,
			"took", time.Since(begin),
			"id", id.String(),
		)
	}(time.Now().UTC())

	err = mw.next.Delete(ctx, id)
	return
}

func (mw *loggingMW) Search(
	ctx context.Context,
	country string,
	since, until time.Time,
	limit int,
	orderASC bool,
) (users []users.User, err error) {
	defer func(begin time.Time) {
		usersCount := 0
		if users != nil {
			usersCount = len(users)
		}

		level.Debug(mw.logger).Log(
			"method", "SearchUsers",
			"err", err,
			"took", time.Since(begin),
			"country", country,
			"since", since,
			"until", until,
			"limit", limit,
			"orderASC", limit,
			"usersCount", usersCount,
		)
	}(time.Now().UTC())

	users, err = mw.next.Search(ctx, country, since, until, limit, orderASC)
	return
}
