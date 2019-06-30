package web
import (
	"net/http"
	"sort"
	"fmt"
	"log"
	"html/template"
	"github.com/gobuffalo/packr/v2"
)

func (t *Server) handlerShowHeaders(w http.ResponseWriter, r *http.Request) {

	// リクエストヘッダを表にする
	type keyValue struct {
		K string
		V interface{}
	}
	type rh2tmplRet struct {
		Req interface{}
		Headers []keyValue
	}
	rh2tmpl:=func(r *http.Request) rh2tmplRet {
		rh := rh2tmplRet{}
		rh.Req = r
		tb := []keyValue{}
		for k,v := range r.Header {
			tb = append(tb, keyValue{K: k, V: v})
		}
		sort.Slice(tb, func(i,j int) bool {
			return tb[i].K < tb[j].K
		})
		rh.Headers = tb
		return rh
	}

	log.Printf("PATH: %s", r.URL.Path)
	box := packr.New("templates","../../assets/templates")
	s, err := box.FindString("headers.tpl.html")
	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}

	tp,err := template.New("T").Parse(s)
	// log.Print( "debug: \n" + spew.Sdump(tp) )

	err = tp.Execute(w, rh2tmpl(r))

	if err != nil {
		fmt.Fprintf(w,"error")
		log.Printf("warn: %s",err)
	}

}
