package context

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/godemo/common"
	"github.com/godemo/handler"
)

// App struct
type App struct {
	Router      chi.Router
	UserHandler *handler.UserHandler
}

// Initialize router and handler
func (a *App) Initialize(user, password, dbname string) {
	// connectionString :=
	// 	fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	// var err error
	// a.DB, err = sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	a.Router = chi.NewRouter()
	a.initializeRoutes()
	a.UserHandler = &handler.UserHandler{}
}

// Run runs app at the addr
func (a *App) Run(addr string) {
	log.Println("App Run() at ", addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	// a.Router.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"*"}}))
	a.Router.Use((cors.AllowAll().Handler))

	a.Router.Get("/api/users/{id:[0-9]+}", a.UserHandler.GetUser)
	// Serve Swagger API Docs
	a.Router.Mount("/doc", http.StripPrefix("/doc/", http.FileServer(http.Dir("./documentation"))))
	a.Router.Get("/doc", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/doc/", http.StatusFound)
	})
	a.Router.NotFound(unhandled)
}

func unhandled(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI, " not found")
	common.RespondWithError(w, http.StatusNotFound, "Please check api doc")
}
func redirect() {

}
