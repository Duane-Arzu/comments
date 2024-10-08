package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

)

func main() {
	var settings serverConfig

    flag.IntVar(&settings.port, "port", 4000, "Server port")
    flag.StringVar(&settings.environment, "env", "development",
                  "Environment(development|staging|production)")
    flag.Parse()
    
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    appInstance := &applicationDependencies {
        config: settings,
        logger: logger,
    }

	router := http.NewServeMux()
    router.HandleFunc("/v1/healthcheck", appInstance.healthcheckHandler)

    apiServer := &http.Server {
        Addr: fmt.Sprintf(":%d", settings.port),
        Handler: appInstance.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
        ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
    }
	logger.Info("starting server", "address", apiServer.Addr,
	"environment", settings.environment)
err := apiServer.ListenAndServe()
logger.Error(err.Error())
os.Exit(1)

}

const appVersion = "1.0.0"

type serverConfig struct {
    port int 
    environment string
}

type applicationDependencies struct {
    config serverConfig
    logger *slog.Logger
}

//This is the command to push an existing repository from the command line
// git remote add origin https://github.com/Duane-Arzu/flower.git
// git branch -M main
// git push -u origin main