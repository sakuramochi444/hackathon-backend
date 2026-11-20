package usecase

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"kaizen/dao"
	"kaizen/model"
	"log"
	"time"
	"unicode/utf8"

	"github.com/oklog/ulid"
)

var ErrInvalidUser = errors.New("invalid user")

func UserRegisterUsecase(context context.Context, req model.User, db *sql.DB) (error, ulid.ULID) {
	err := validateUser(req)
	if err != nil {
		return err, ulid.ULID{}
	}

	id, err := ulid.New(ulid.Timestamp(time.Now()), rand.Reader)
	if err != nil {
		log.Printf("fail: ulid.New, %v", err)
		return err, id
	}

	err = dao.InsertUser(context, id, req)
	if err != nil {
		log.Printf("fail: , %v", err)
		return err, id
	}

	return nil, id
}

func validateUser(req model.User) error {
	if req.Name == "" || utf8.RuneCountInString(req.Name) > 50 {
		return ErrInvalidUser
	}
	if req.Age < 20 || req.Age > 80 {
		return ErrInvalidUser
	}

	return nil
}
