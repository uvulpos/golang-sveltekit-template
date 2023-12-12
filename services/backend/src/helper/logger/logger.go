package logger

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Debug *log.Logger = log.New(
		os.Stdout,
		"🐞 DEBUG: ",
		log.Lshortfile,
	)

	Info *log.Logger = log.New(
		os.Stdout,
		"🧑🏻‍🏫 INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Warning *log.Logger = log.New(
		os.Stdout,
		"🚧 WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Error *log.Logger = log.New(
		os.Stdout,
		"🔥 ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Pretty = func(_ string, objs ...interface{}) {
		msg, err := json.MarshalIndent(objs, "", "  ")
		if err != nil {
			Error.Printf("unhandled error: %s", err)
		}

		Debug.Printf("%s", msg)
	}
)
