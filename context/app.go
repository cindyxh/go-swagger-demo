package context

import (
	"log"
	"net/http"

	"github.com/godemo/handler"
	"github.com/gorilla/mux"
)

type App struct {
	Router      *mux.Router
	UserHandler *handler.UserHandler
}

func (a *App) Initialize(user, password, dbname string) {
	// connectionString :=
	// 	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	// var err error
	// a.DB, err = sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	a.Router = mux.NewRouter()
	a.initializeRoutes()
	a.UserHandler = &handler.UserHandler{}
}

// Run runs app at the addr
func (a *App) Run(addr string) {
	log.Println("App Run() at ", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/users/{id:[0-9]+}", a.UserHandler.GetUser).Methods("GET")
	// Serve Swagger API Docs
	a.Router.PathPrefix("/api").Handler(http.StripPrefix("/api/", http.FileServer(http.Dir("./documentation"))))
}
