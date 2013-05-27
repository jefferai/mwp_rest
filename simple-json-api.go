package main

import (
	"github.com/coocood/jas"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type ReceivedData struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Guess    string `json:"guess"`
	NextChar string `json:"nextchar"`
}

func sigintCatcher() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT)
	<-ch
	os.Exit(0)
}

func main() {
	// Catch signals
	go sigintCatcher()

	router := jas.NewRouter(new(Wordgame))
	router.RequestErrorLogger = router.InternalErrorLogger

	log.Println("Starting serving on paths:\n" + router.HandledPaths(true))
	http.Handle(router.BasePath, router)
	err := http.ListenAndServe("127.0.0.1:8888", nil)
	if err != nil {
		panic(err)
	}
}
