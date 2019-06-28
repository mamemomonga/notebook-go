package web
import (
	"net/http"
	"sort"
	"fmt"
	"log"
	"html/template"
	"github.com/gobuffalo/packr/v2"
)


type KV struct {
	K string
	V interface{}
}

type tmplRH struct {
	Req     interface{}
	Headers []KV
}

func (this *Server) rh2tmpl(r *http.Request) tmplRH {
	t := tmplRH{}
	t.Req = r

	tb := []KV{}
	for k,v := range r.Header {
		tb = append(tb, KV{K: k, V: v})
	}
	sort.Slice(tb, func(i,j int) bool {
		return tb[i].K < tb[j].K
	})
	t.Headers = tb

	return t
}

func (this *Server) handlerShowHeaders(w http.ResponseWriter, r *http.Request) {
	log.Printf("PATH: %s", r.URL.Path)

	box := packr.New("templates","../../assets/templates")
	s, err := box.FindString("headers.tpl.html")
	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}

	t,err := template.New("T").Parse(s)
	// log.Print( "debug: \n" + spew.Sdump(tp) )

	err = t.Execute(w, this.rh2tmpl(r))

	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}

}
