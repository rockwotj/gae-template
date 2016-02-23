package app

import (
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func TestHander(w http.ResponseWriter, r *http.Request, _ httprouter.Params) interface{} {
	ctx := appengine.NewContext(r)
	q := datastore.NewQuery("Test").Order("-Timestamp")
	statusList := TestResults{}
	if _, err := q.GetAll(ctx, &statusList); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return Error{
			Message: "Cannot query for status",
			Code:    500,
		}
	}
	return statusList
}
