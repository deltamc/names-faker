package main

import (
	"fmt"
	"github.com/deltamc/names-faker/user"
	"github.com/deltamc/names-faker/db"
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	dbPool := db.OpenDB()
	defer dbPool.Close()

	stmt, err := dbPool.Prepare(
		"INSERT INTO " +
			"`users` (`login`, `password`, `first_name`, `last_name`, `age`, `sex`, `interests`, `city`) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	fmt.Println("Введите кол-во аккаунтов")
	var countStr string
	fmt.Scan(&countStr)
	count, _ := strconv.Atoi(countStr)
	for i:=1; i <= count; i++ {
		u := user.GetUser()
		stmt.Exec(u.Login, u.Password, u.FirstName, u.LastName, u.Age, u.Sex, u.Interests, u.City)
	}
}


