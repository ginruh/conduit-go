package seeders

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/iyorozuya/real-world-app/internal/sqlc"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"strconv"
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
	bcryptCost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)

	fmt.Println("=============== SEEDED USER(S) ===================")
	fmt.Printf("Email: %s, Password: %s", user.Email, user.Password)
	fmt.Println()

	_, err = s.q.CreateUser(context.Background(), sqlc.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: string(passwordHash),
	})
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Unable to seed user to db")
		return
	}
	log.Println("Fake users saved successfully")
}
