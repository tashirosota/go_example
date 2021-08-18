package chi

import (
	"fmt"
	"gigphil/app/controllers/area"
	"gigphil/app/controllers/root" // request増えるたびに都度増やすの苦痛だな memo: packageは1dirに付き１つ
	"net/http"
  "os"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
func Setup() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port =  "4000"
	}
	fmt.Printf("\n===== Listening on http://127.0.0.1:" + port + " =====\n\n")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", root.Get)
	r.Get("/areas", area.GetList)
	http.ListenAndServe(":"+port, r)
}
