package main

import (
	"GOlang/models"

	//"errors"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1903"
	dbname   = "postgres"
)

//func main() {
//	psqlInfo := fmt.Sprintf(
//		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//	us, err := models.NewUserService(psqlInfo)
//	if err != nil {
//		panic(err)
//	}
//
//	defer us.Close()
//
//	us.DestructiveReset()
//
//	// Create a user
//	user := models.User{
//		Name:  "Lewis Hamilton",
//		Email: "lewishamilton@ferrari.com",
//	}
//
//	if err := us.Create(&user); err != nil {
//		panic(err)
//	}
//
//	foundUser, err := us.ByID(1)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(foundUser)
//
//	// Update a user
//	user.Name = "Billie Eilish"
//	if err := us.Update(&user); err != nil {
//		panic(err)
//	}
//
//	foundUser, err = us.ByEmail("lewishamilton@ferrari.com")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(foundUser)
//
//	//Delete a user
//	if err := us.Delete(foundUser.ID); err != nil {
//		panic(err)
//	}
//
//	_, err = us.ByID(foundUser.ID)
//	if !errors.Is(err, models.ErrNotFound) {
//		panic(err)
//	}
//}

func main() {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}

	defer us.Close()

	us.DestructiveReset()

	user := models.User{
		Name:     "Michael Scott",
		Email:    "michael@dundermifflin.com",
		Password: "bestboss",
	}
	err = us.Create(&user)
	if err != nil {
		panic(err)
	}
	// Verify that the user has a Remember and RememberHash
	fmt.Printf("%+v\n", user)
	if user.Remember == "" {
		panic("Invalid remember token")
	}
	// Now verify that we can look up a user with that remember
	// token
	user2, err := us.ByRemember(user.Remember)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", *user2)
}
