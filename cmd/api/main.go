package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Hello")

	logger := httplog.NewLogger("dummy-service", httplog.Options{
		LogLevel:        slog.LevelDebug,
		Concise:         false,
		RequestHeaders:  true,
		ResponseHeaders: true,
	})

	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Get("/success", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Success"))
	})
	r.Get("/notfound", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(404), 404)
	})
	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(400), 400)
	})
	r.Get("/error/{code}", func(w http.ResponseWriter, r *http.Request) {
		codeString := chi.URLParam(r, "code")
		code, err := strconv.Atoi(codeString)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			http.Error(w, http.StatusText(code), code)
		}
	})
	r.Get("/timeout", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(30 * time.Second)
		w.Write([]byte("done after 30 seconds"))
	})
	r.Get("/timeout/{duration}", func(w http.ResponseWriter, r *http.Request) {
		durationString := chi.URLParam(r, "duration")
		duration, err := strconv.Atoi(durationString)
		if err != nil {
			http.Error(w, err.Error(), 500)
		} else {
			time.Sleep(time.Duration(duration) * time.Second)
			w.Write([]byte(fmt.Sprintf("done after %d seconds", duration)))
		}
	})
	http.ListenAndServe(":9292", r)
}
