package hello

import (
	"net/http"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir("./")))

}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	//fmt.Fprint(w, "Site is under construction!!!")
// }
