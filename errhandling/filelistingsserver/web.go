package main

import (
	"net/http"
	"learngo/errhandling/filelistingsserver/filelisting"
	"os"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic : %v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(w, r)
		if err != nil {
			// log.Warn("Error handling request: %s", err.Error())
			log.Printf("Error occured handling request : %s", err.Error())
			if userError, ok := err.(userError); ok {
				http.Error(w, userError.Message(), http.StatusBadRequest)
			} else {
				code := http.StatusOK
				switch {
				case os.IsNotExist(err):
					code = http.StatusNotFound
				case os.IsPermission(err):
					code = http.StatusForbidden
				default:
					code = http.StatusInternalServerError
				}
				http.Error(w, http.StatusText(code), code)
			}
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
