package chi

import (
	"fmt"
	"gigphil/app/controllers/area"
	"gigphil/app/controllers/root" // request増えるたびに都度増やすの苦痛だな memo: packageは1dirに付き１つ
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Setup() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}
	fmt.Printf("\n===== Listening on http://127.0.0.1:" + port + " =====\n\n")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Get("/", root.Get)
	r.Get("/areas", area.GetList)
	http.ListenAndServe(":"+port, r)
}
