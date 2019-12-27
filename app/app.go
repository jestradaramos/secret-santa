package app

import (
	"fmt"
	"net/url"
	"secret-santa/domain"
	"secret-santa/repo"
)

// App this is the thing we run
type App struct {
	// Currently nothing
}

// NewApp This will create our app with all dependencies and blah
func NewApp() (*App, error) {
	app := &App{}
	dbString := "postgres://postgres:docker@localhost:5432/postgres?sslmode=disable"
	dbURL, err := url.Parse(dbString)
	if err != nil {
		return nil, err
	}

	db, err := repo.InitDatabase(dbURL)
	if err != nil {
		return nil, err
	}

	// Init all the tables
	db.InitGroupTable()
	db.InitMemberTable()

	fmt.Print("Tables started")

	// Add a group
	g := &domain.Group{
		ID:           "Group 1",
		MoneyLimit:   10.00,
		Deadline:     "now",
		ExchangeDate: "tomorrow",
	}

	msg, err := db.AddGroup(g)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(msg)
	// Add a member with wrong group

	m1 := &domain.Member{
		Name:    "Not Me",
		Email:   "email@email.com",
		Spouse:  "Nope",
		GroupID: "Group 2",
	}

	err = db.AddMember(m1)
	if err != nil {
		fmt.Println("It worked")
	} else {
		fmt.Println("It did not work")
	}

	// Add a member with correct group
	m2 := &domain.Member{
		Name:    "Me",
		Email:   "email@email.com",
		Spouse:  "yes",
		GroupID: "Group 1",
	}

	err = db.AddMember(m2)
	if err != nil {
		fmt.Println("It did not work")
	} else {
		fmt.Println("It did work")
	}
	return app, nil

}
