package app

import (
	"database/sql"
	"log"
	"net"
	"time"
	"net/http"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
	"api/app/infrastructure/db/postgresql"
	userCore "api/app/core/user"
	userGateway "api/app/infrastructure/db/postgresql/gateway/user"
	userApi "api/app/presentation/api/user"

)

type Application struct {
	db *sql.DB
}

func (app *Application) Inicialize() {
	db := postgresql.DbConnect{}
	db.Connect()
	app.db = db.GetCur()

	file, err := ioutil.ReadFile("app/infrastructure/db/postgresql/create_tables.sql")
	if err != nil {
		panic(err)
	}
	_, err = app.db.Exec(string(file))
	if err != nil {
		panic(err)
	}

	fake_data, err := ioutil.ReadFile("app/infrastructure/db/postgresql/fake_data.sql")
	if err != nil {
		panic(err)
	}
	_, _ = app.db.Exec(string(fake_data))

	userGateway := userGateway.NewGateway(app.db)
	userUseCase := userCore.NewUserUseCase(userGateway)
	userHandler := userApi.NewHandler(userUseCase)

	router := httprouter.New()
	userHandler.Register(router)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,		
	}

	server.Serve(listener)

	log.Fatal()
}