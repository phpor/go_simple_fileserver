package controller
import (
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(r.RequestURI))
	switch r.RequestURI {
	case "/list":
		List(w, r)
		break
	case "/upload":
		Upload(w, r); break
	case "/download":
		Download(w, r); break
	default:
		http.NotFound(w, r);break
	}
}
