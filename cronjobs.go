package app

import (
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

func TestCronjob(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(r)
	status := Test{
		Timestamp: time.Now().Unix() * 1000, // ms
		Message:   "Hello, World",
	}
	_, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Test", nil), &status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
