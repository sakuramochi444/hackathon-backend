package usecase

import (
	"database/sql"
	"kaizen/dao"
	"kaizen/model"
	"log"
)

func UserSearchUsecase(name string, db *sql.DB) (error, []model.User) {
	err, users := dao.SearchUsersByName(name)
	if err != nil {
		log.Printf("fail: , %v\n", err)
		return err, nil
	}

	return nil, users
}
