package main

import (
	web_smartphone "course-phones-review/gadgets/smartphones/web"
	web_user "course-phones-review/gadgets/users/web"
	"encoding/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	//reviews "course-phones-review/reviews/web"

	"net/http"
)

func Routes(
	sph *web_smartphone.CreateSmartphoneHandler,
	uh *web_user.CreateUserHandler,
	//reviewHandler *reviews.ReviewHandler,
) *chi.Mux {
	mux := chi.NewMux()

	// globals middleware
	mux.Use(
		middleware.Logger,    //log every http request
		middleware.Recoverer, // recover if a panic occurs
	)

	mux.Route("/users", func(r chi.Router) {
		r.Post("/auth", uh.AuthUserHandler)
		r.Post("/", uh.SaveUserHandler)
		r.Get("/{userID:[0-9]+}", uh.GetUserHandler)
	})

	mux.Route("/smartphones", func(r chi.Router) {
		r.Post("/", sph.SaveSmartphoneHandler)
		r.Get("/{smartphoneID:[0-9]+}", sph.GetSmartphoneHandler)
	})

	//mux.Post("/smartphones", sph.SaveSmartphoneHandler)
	mux.Get("/hello", helloHandler)
	//mux.Post("/reviews", reviewHandler.AddReviewHandler)

	return mux
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("done-by", "tomas")

	res := map[string]interface{}{"message": "hello world"}

	_ = json.NewEncoder(w).Encode(res)
}
