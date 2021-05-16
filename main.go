package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name     string `json: "name"`
	Balance  int    `json: "balance"`
	Aircraft string `json: "aircraft"`
	Sortie   string `json: "sortie"`
	Air      string `json: "air"`
	Ground   string `json: "ground"`
	Killed   string `json: "killed"`
}

func makeJson() {
	origin, err := os.Open("./userdata.csv")

	if err != nil {
		fmt.Println(err)
	}

	defer origin.Close()

	reader := csv.NewReader(origin)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var user User
	var data []User

	for index, record := range records {
		if index != 0 {
			user.Name = record[0]
			user.Balance, _ = strconv.Atoi(record[1])
			user.Aircraft = record[2]
			user.Sortie = record[3]
			user.Air = record[4]
			user.Ground = record[5]
			user.Killed = record[6]
			data = append(data, user)
		}
	}

	proc, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(string(proc))

	result, err := os.Create("./userdata.json")

	if err != nil {
		fmt.Println(err)
	}

	defer result.Close()

	result.Write(proc)
	result.Close()
}

func main() {
	app := fiber.New()

	app.Static("/", "./web")

	app.Get("/userdata.json", func(c *fiber.Ctx) error {
		makeJson()
		return c.SendFile("./userdata.json")
	})

	app.Listen(":8080")
}
