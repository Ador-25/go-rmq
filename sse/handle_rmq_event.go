package sse

import (
	"fmt"
	"net/http"
	"rmq/utils"
)

var MessageChannel = make(chan string)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	utils.AllowCors(w)
	for {
		msg := <-MessageChannel
		fmt.Fprintf(w, "data: %s\n\n", msg)
		w.(http.Flusher).Flush()
	}
	closeNotify := w.(http.CloseNotifier).CloseNotify()
	<-closeNotify
}
