package user

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/MikelSot/trival-proyect/model"
)

const (
	_totalUsers = 15000
	_timeout    = 2.25
)

type User struct {
	service DocService
}

func New(service DocService) User {
	return User{service}
}

func (u User) GetUniqueUsers() (model.Users, error) {
	start := time.Now()

	uniqueResults := make(map[string]model.Service)

	for len(uniqueResults) < _totalUsers {
		results, err := u.service.GetRandomUser()
		if err != nil {
			return nil, fmt.Errorf("user.service.GetRandomUser: %w", err)
		}

		fmt.Println("len(uniqueResults)")
		fmt.Println(len(uniqueResults))

		resultMap := results.UniqueMap()
		for _, result := range resultMap {

			if len(uniqueResults) >= _totalUsers {
				break
			}

			uniqueResults[result.Login.Uuid] = result
		}

	}

	elapsed := time.Since(start)
	if elapsed.Seconds() > _timeout {
		log.Warn("user.GetUniqueUsers(): timeout")

		return u.getUsers(uniqueResults), nil
	}

	return u.getUsers(uniqueResults), nil
}

func (u User) getUsers(results map[string]model.Service) model.Users {
	var users model.Users
	for _, result := range results {
		user := model.User{
			Uuid:      result.Login.Uuid,
			FirstName: result.Name.First,
			LastName:  result.Name.Last,
			Email:     result.Email,
			Gender:    result.Gender,
		}

		users = append(users, user)
	}

	return users
}
