package routes

import (
	"github.com/gorilla/mux"
)

var RegisterBookstoreRoutes = func(router, *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Method("POST")
	router.HandleFunc("/book", controllers.GetBook).Method("GET")
	router.HandleFunc("/book", controllers.GetBookById).Method("GET")
	router.HandleFunc("/book", controllers.UpdateBook).Method("PUT")
	router.HandleFunc("/book", controllers.DeleteBook).Method("DELETE")
}
