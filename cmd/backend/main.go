package main

import (
	"gebes.io/sticker_backend/pkg/database"
	"gebes.io/sticker_backend/pkg/logger"
	"gebes.io/sticker_backend/pkg/router"
	"gebes.io/sticker_backend/pkg/signals"
	"net/http"
	"os"
)

func main() {
	signals.ListenForSigterm()
	time, err := database.Ping()
	if err != nil {
		logger.Error.Fatalln("Could not ping the database:", err)
	}
	logger.Info.Println("Postgres database delay is", time)

	logger.Info.Println("Starting to listen on port 8080 on PID", os.Getpid())
	err = router.Listen()
	if err != nil && err != http.ErrServerClosed {
		logger.Error.Fatalln("Could not start the router:", err)
	}

	signals.WaitForCleanup()
}
