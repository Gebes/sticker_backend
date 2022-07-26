package logger

import (
	"io"
	"log"
	"os"
)

const (
	flags       = log.Ldate | log.Ltime | log.Lshortfile
	logFileName = "sticker"
)

var (
	Info  = log.New(os.Stdout, "INFO  ", flags)
	Debug = log.New(os.Stdout, "DEBUG ", flags)
	Error = log.New(os.Stderr, "ERROR ", flags)
)

func init() {
	logFile, err := os.OpenFile(logFileName+".log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		Error.Fatalln("Could not initialize log file:", err)
	}

	output := io.MultiWriter(os.Stdout, logFile)
	Info.SetOutput(output)
	Error.SetOutput(output)
	Debug.SetOutput(output)

	log.SetOutput(Debug.Writer())
	log.SetPrefix("DEBUG ")
	log.SetFlags(flags)
}
