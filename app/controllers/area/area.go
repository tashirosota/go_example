package area

import (
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("welcome areas"))
}