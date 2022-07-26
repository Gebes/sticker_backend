package signals

import (
	"gebes.io/sticker_backend/pkg/database"
	"gebes.io/sticker_backend/pkg/logger"
	"gebes.io/sticker_backend/pkg/router"
	"os"
	"os/signal"
	"syscall"
)

var done = make(chan bool, 1)

func ListenForSigterm() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		logger.Debug.Println("Received signal", sig)
		err := router.Shutdown()
		if err != nil {
			logger.Error.Println("Could not stop router:", err)
		}
		logger.Info.Println("Closed router")

		err = database.Close()
		if err != nil {
			logger.Error.Println("Could not close database connection:", err)
		}
		logger.Info.Println("Closed database")

		logger.Info.Println("Stopping Sticker backend")
		done <- true
	}()

}

func WaitForCleanup() {
	<-done
}
