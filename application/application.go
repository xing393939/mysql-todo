package application

import (
	"github.com/carbocation/interpose"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"net/http"

	"mysql-todo/handlers"
	"mysql-todo/middlewares"
)

// New is the constructor for Application struct.
func New(config *viper.Viper) (*Application, error) {
	dsn := config.Get("dsn").(string)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	cookieStoreSecret := config.Get("cookie_secret").(string)

	app := &Application{}
	app.config = config
	app.dsn = dsn
	app.db = db
	app.sessionStore = sessions.NewCookieStore([]byte(cookieStoreSecret))

	return app, nil
}

// Application is the application object that runs HTTP server.
type Application struct {
	config       *viper.Viper
	dsn          string
	db           *sqlx.DB
	sessionStore sessions.Store
}

func (app *Application) MiddlewareStruct() (*interpose.Middleware, error) {
	middle := interpose.New()
	middle.Use(middlewares.SetDB(app.db))
	middle.Use(middlewares.SetSessionStore(app.sessionStore))
	middle.UseHandler(app.mux())
	return middle, nil
}

func (app *Application) mux() *mux.Router {
	MustLogin := middlewares.MustLogin

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(handlers.GetIndex)).Methods("GET")
	router.Handle("/brand", http.HandlerFunc(handlers.GetBrand)).Methods("GET")
	router.Handle("/brand/product", http.HandlerFunc(handlers.GetBrandProduct)).Methods("GET")
	router.Handle("/users", http.HandlerFunc(handlers.GetAccountHome)).Methods("GET")
	router.HandleFunc("/users/signup", handlers.GetSignup).Methods("GET")
	router.HandleFunc("/users/signup", handlers.PostSignup).Methods("POST")
	router.HandleFunc("/users/login", handlers.GetLogin).Methods("GET")
	router.HandleFunc("/users/login", handlers.PostLogin).Methods("POST")
	router.HandleFunc("/users/logout", handlers.GetLogout).Methods("GET")
	router.Handle("/users/{id:[0-9]+}", MustLogin(http.HandlerFunc(handlers.PostPutDeleteUsersID))).Methods("POST", "PUT", "DELETE")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return router
}
