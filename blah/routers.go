package blah

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddPerson",
		"POST",
		"/people",
		AddPerson,
	},

	Route{
		"AddPersonPhoto",
		"POST",
		"/people/{personName}/photos",
		AddPersonPhoto,
	},

	Route{
		"DeletePeople",
		"DELETE",
		"/people",
		DeletePeople,
	},

	Route{
		"DeletePerson",
		"DELETE",
		"/people/{personName}",
		DeletePerson,
	},

	Route{
		"DeletePersonPhoto",
		"DELETE",
		"/people/{personName}/photos/{id}",
		DeletePersonPhoto,
	},

	Route{
		"FindPeople",
		"GET",
		"/people",
		FindPeople,
	},

	Route{
		"FindPersonByName",
		"GET",
		"/people/{personName}",
		FindPersonByName,
	},

	Route{
		"GetPersonPhoto",
		"GET",
		"/people/{personName}/photos/{id}",
		GetPersonPhoto,
	},

	Route{
		"PeoplePersonNamePatch",
		"PATCH",
		"/people/{personName}",
		PeoplePersonNamePatch,
	},

	Route{
		"PeoplePersonNamePhotosGet",
		"GET",
		"/people/{personName}/photos",
		PeoplePersonNamePhotosGet,
	},

	Route{
		"RootGet",
		"GET",
		"/",
		RootGet,
	},

}
