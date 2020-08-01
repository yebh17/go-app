package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	chiprometheus "github.com/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Welcome struct {
	Name string
	Time string
}

func main() {

	welcome := Welcome{"User, you are been hacked", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	n := chi.NewRouter()
	m := chiprometheus.NewMiddleware("api", 50, 100, 200, 500, 1000, 5000)

	n.Use(m)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	n.Handle("/metrics", promhttp.Handler())
	n.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on PORT 6777")
	fmt.Println(http.ListenAndServe(":6777", n))
}
