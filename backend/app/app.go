package app

import (
	"log"
	"net/url"
	"secret-santa/api/restapi"
	"secret-santa/api/restapi/operations"
	"secret-santa/repo"

	"github.com/go-openapi/loads"
)

// App this is the thing we run
type App struct {
	Server *restapi.Server
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

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewSecretSantaAPI(swaggerSpec)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parser := flags.NewParser(server, flags.Default)
	// parser.ShortDescription = "Secret Santa"
	// parser.LongDescription = swaggerSpec.Spec().Info.Description

	// server.ConfigureFlags()
	// for _, optsGroup := range api.CommandLineOptionsGroups {
	// 	_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }

	// if _, err := parser.Parse(); err != nil {
	// 	code := 1
	// 	if fe, ok := err.(*flags.Error); ok {
	// 		if fe.Type == flags.ErrHelp {
	// 			code = 0
	// 		}
	// 	}
	// 	os.Exit(code)
	// }
	server.ConfigureAPI(db)

	server.Port = 8080

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
	return app, nil

}

// Start ...
func (a *App) Start() {

}
