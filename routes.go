package app

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type HTTPMethod int

const (
	GET HTTPMethod = 1 << iota
	POST
	PUT
	DELETE
)

type RestEndpoint func(http.ResponseWriter, *http.Request, httprouter.Params) interface{}

type JSONRoute struct {
	Path    string
	Method  HTTPMethod
	Handler RestEndpoint
}

type FileRoute struct {
	Path string
	File string
}

type CronJob struct {
	Path string
	Job  httprouter.Handle
}

type JSONRoutes []JSONRoute
type FileRoutes []FileRoute
type CronJobs []CronJob

var JSON_ROUTES = JSONRoutes{
	JSONRoute{
		"/api/test_results",
		GET,
		TestHander,
	},
}

var FILE_ROUTES = FileRoutes{
	FileRoute{
		"/",
		"./static/index.html",
	},
}

var CRONJOBS = CronJobs{
	CronJob{
		"/tasks/test",
		TestCronjob,
	},
}
