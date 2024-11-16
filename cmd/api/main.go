package main

import (
	"context"
	"database/sql"
	"flag"
<<<<<<< HEAD
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/Duane-Arzu/comments/internal/data"
	"github.com/Duane-Arzu/comments/internal/mailer"
	_ "github.com/lib/pq"
)

const appVersion = "7.0.0"
=======
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Duane-Arzu/comments/internal/data"
	_ "github.com/lib/pq"
)

const appVersion = "3.0.0"
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef

type serverConfig struct {
	port        int
	environment string
	db          struct {
		dsn string
	}
<<<<<<< HEAD
	limiter struct {
		rps     float64 // requests per second
		burst   int     // initial requests possible
		enabled bool    // enable or disable rate limiter
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type applicationDependencies struct {
	config       serverConfig
	logger       *slog.Logger
	commentModel data.CommentModel
	userModel    data.UserModel
	mailer       mailer.Mailer
	wg           sync.WaitGroup
	tokenModel   data.TokenModel
}

func main() {
	var setting serverConfig

	flag.IntVar(&setting.port, "port", 4000, "Server port")
	flag.StringVar(&setting.environment, "env", "development", "Environment (development|staging|production)")
	//read the dsn
	flag.StringVar(&setting.db.dsn, "db-dsn", "postgres://comments:comments@localhost/comments?sslmode=disable", "PostgreSQL DSN")
	flag.Float64Var(&setting.limiter.rps, "limiter-rps", 2, "Rate Limiter maximum requests per second")

	flag.IntVar(&setting.limiter.burst, "limiter-burst", 5, "Rate Limiter maximum burst")

	flag.BoolVar(&setting.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&setting.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
	// We have port 25, 465, 587, 2525. If 25 doesn't work choose another
	flag.IntVar(&setting.smtp.port, "smtp-port", 2525, "SMTP port")
	// Use your Username value provided by Mailtrap
	flag.StringVar(&setting.smtp.username, "smtp-username", "c3e1c1678d71c9", "SMTP username")

	flag.StringVar(&setting.smtp.password, "smtp-password", "38d1f200e85005", "SMTP password")

	flag.StringVar(&setting.smtp.sender, "smtp-sender", "Comments Community <no-reply@commentscommunity.alexperaza.net>", "SMTP sender")

=======
}

type applicationDependences struct {
	config       serverConfig
	logger       *slog.Logger
	commentModel data.CommentModel
}

func main() {
	var settings serverConfig
	flag.IntVar(&settings.port, "port", 4000, "Server Port")
	flag.StringVar(&settings.environment, "env", "development", "Environment(development|staging|production)")
	//read the dsn
	flag.StringVar(&settings.db.dsn, "db-dsn", "postgres://comments:comments@localhost/comments?sslmode=disable", "PostgreSQL DSN")
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

<<<<<<< HEAD
	// the call to openDB() sets up our connection pool
	db, err := openDB(setting)
	if err != nil {
		logger.Error("Database connection failed")
		os.Exit(1)
	}
	// release the database resources before exiting
	defer db.Close()

	logger.Info("Database connection pool established")

	appInstance := &applicationDependencies{
		config:       setting,
		logger:       logger,
		userModel:    data.UserModel{DB: db},
		commentModel: data.CommentModel{DB: db},
		tokenModel:   data.TokenModel{DB: db},
		mailer: mailer.New(setting.smtp.host, setting.smtp.port,
			setting.smtp.username, setting.smtp.password, setting.smtp.sender),
	}

	err = appInstance.serve()
=======
	//the call to openDB() sets up our connection pool
	db, err := openDB(settings)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	//release the database connection before exiting
	defer db.Close()

	logger.Info("Database Connection Pool Established")

	appInstance := &applicationDependences{
		config:       settings,
		logger:       logger,
		commentModel: data.CommentModel{DB: db},
	}

	apiServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", settings.port),
		Handler:      appInstance.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting Server", "address", apiServer.Addr, "environment", settings.environment)
	err = apiServer.ListenAndServe()
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(settings serverConfig) (*sql.DB, error) {
<<<<<<< HEAD
	// open a connection pool
=======
	//open a connection pool
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
	db, err := sql.Open("postgres", settings.db.dsn)
	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	// set a context to ensure DB operations don't take too long
	ctx, cancel := context.WithTimeout(context.Background(),
		5*time.Second)
	defer cancel()

	// let's test if the connection pool was created
	// we trying pinging it with a 5-second timeout
=======
	//set context to ensure DB operations dont take too long
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	//pinging connection pool to verify it was created, with a 5-second timeout
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

<<<<<<< HEAD
	// return the connection pool (sql.DB)
	return db, nil

}
=======
	//return the connection pool (sql.DB)
	return db, nil
}

//This is the command to push an existing repository from the command line
// git remote add origin https://github.com/Duane-Arzu/flower.git
// git branch -M main
// git push -u origin main

//hello
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
