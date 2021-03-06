package web

import (
	"log"
	"net/http"
	"github.com/gobuffalo/packr/v2"
)

type Server struct {
}

func New() *Server {
	return new(Server)
}

func (t *Server) Run(listen string) error {

	box := packr.New("public","../../assets/static")

	http.HandleFunc("/",t.handlerShowHeaders)
	http.Handle("/static/",http.StripPrefix("/static",http.FileServer(box)))

	log.Printf("info: Start Listening at http://%s/", listen)
	if err := http.ListenAndServe(listen, nil); err != nil {
		return err
	}
	return nil
}
