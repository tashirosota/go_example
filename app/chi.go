package chi

import (
	"fmt"
	"net/http"
	"gigphil/app/requests"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
const port string = "4000"
func Setup(){
	fmt.Printf("\n===== Listening on http://127.0.0.1:" + port + " =====\n\n")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", root.Get)
	http.ListenAndServe(":" + port, r)
}