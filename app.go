package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	router := httprouter.New()
	for _, route := range JSON_ROUTES {
		handler := HandleJson(route.Handler)
		switch route.Method {
		case GET:
			router.GET(route.Path, handler)
		case POST:
			router.POST(route.Path, handler)
		case PUT:
			router.PUT(route.Path, handler)
		case DELETE:
			router.DELETE(route.Path, handler)
		}
	}
	for _, route := range FILE_ROUTES {
		router.GET(route.Path, HandleFile(route.File))
	}
	for _, route := range CRONJOBS {
		router.GET(route.Path, HandleCron(route.Job))
	}
	router.ServeFiles("/static/*filepath", http.Dir("./static/"))
	http.Handle("/", router)
}

func HandleCron(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		if r.Header.Get("X-Appengine-Cron") == "" {
			http.Error(w, "Forbidden", 403)
		} else {
			handler(w, r, p)
		}
	}
}

func HandleFile(file string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		http.ServeFile(w, r, file)
	}
}

func HandleJson(handler RestEndpoint) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := handler(w, r, p)
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			panic(err)
		}
	}
}
