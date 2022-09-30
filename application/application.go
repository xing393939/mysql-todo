package application

import (
	"github.com/carbocation/interpose"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	"net/http"

	"mysql-todo/handlers"
	"mysql-todo/middlewares"
)

// New is the constructor for Application struct.
func New(config *viper.Viper) (*Application, error) {
	cookieStoreSecret := config.Get("cookie_secret").(string)

	app := &Application{}
	app.config = config
	app.dsn = ""
	app.sessionStore = sessions.NewCookieStore([]byte(cookieStoreSecret))

	return app, nil
}

// Application is the application object that runs HTTP server.
type Application struct {
	config       *viper.Viper
	dsn          string
	sessionStore sessions.Store
}

func (app *Application) MiddlewareStruct() (*interpose.Middleware, error) {
	middle := interpose.New()
	middle.Use(middlewares.SetSessionStore(app.sessionStore))
	middle.UseHandler(app.mux())
	return middle, nil
}

func (app *Application) mux() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(handlers.GetIndex)).Methods("GET")
	router.Handle("/brand/{id:[0-9a-z]+}", http.HandlerFunc(handlers.GetBrand)).Methods("GET")
	router.Handle("/product/{b:[0-9a-z]+}/{p:[\\-0-9a-z]+}", http.HandlerFunc(handlers.GetBrandProduct)).Methods("GET")
	router.Handle("/users", http.HandlerFunc(handlers.GetAccountHome)).Methods("GET")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return router
}
