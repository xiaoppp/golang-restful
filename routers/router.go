package routers

import "github.com/julienschmidt/httprouter"

func Config() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)

	return router
}
