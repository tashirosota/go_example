package area

import (
	"encoding/json"
	area "gigphil/app/models/area"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	names := area.AllNames()
	res, _ := json.Marshal(names)
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","GET" )
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
