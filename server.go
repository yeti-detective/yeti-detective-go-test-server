// this bizniz produced by following article here: https://blog.logrocket.com/creating-a-web-server-with-golang/
// pushed to heroku following instructions here: https://devcenter.heroku.com/articles/getting-started-with-go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"yetiserver/controllers"
)

func main() {
	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}
	fmtPort := fmt.Sprintf(":%s", port)

	http.Handle("/", controllers.FileHandlerStatic())
	http.HandleFunc("/hello", controllers.HelloHandler)
	http.HandleFunc("/form", controllers.FormHandler)

	fmt.Printf("Starting server at %s\n", fmtPort)
	if err := http.ListenAndServe(fmtPort, nil); err != nil {
		log.Fatal(err)
	}

}
