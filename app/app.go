package app

import (
	"database/sql"
	"log"
	"net/http"
	"io/ioutil"
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
	_, err = app.db.Exec(string(fake_data))
	if err != nil {
		panic(err)
	}

	userGateway := userGateway.NewGateway(app.db)
	userUseCase := userCore.NewUserUseCase(userGateway)
	userHandler := userApi.NewHandler(userUseCase)


	mux := http.NewServeMux()

	mux.HandleFunc("/user/balans/increase", userHandler.IncreaseBalance)
	mux.HandleFunc("/user/balans/decrease", userHandler.DecreaseBalance)

	log.Fatal(http.ListenAndServe(":8080", mux))
}