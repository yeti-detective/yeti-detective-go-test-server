package controllers

import "net/http"

func FileHandlerStatic() http.Handler {
	return http.FileServer(http.Dir("./static"))
}
