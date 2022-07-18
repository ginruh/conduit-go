package seeders

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"log"
)

type FakeUser struct {
	Username string `faker:"username"`
	Email    string `faker:"email"`
	Password string `faker:"password"`
}

func (s Seed) seedUsers() {
	user := FakeUser{}
	err := faker.FakeData(&user)
	if err != nil {
		log.Fatalln("Unable to get fake users")
		return
	}
	_, err = s.q.CreateUser(context.Background(), sqlc.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Unable to seed user to db")
		return
	}
	log.Println("Fake users saved successfully")
}
