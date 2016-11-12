package blah

import (
	"net/http"
	"log"
	"encoding/json"
	"strings"
)

// members must be camel case, json.Marshal is lowercase
type Default struct {
	Name string `json:"name"`
	Message []string `json:"message"`
}

var m = make(map[string][]string)

func AddPerson(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var p Default
		if r.Body == nil {
            	   http.Error(w, "Please send a request body", 400)
            	   return
        	}
        	err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
            	   http.Error(w, err.Error(), 400)
            	   return
		}
        	log.Println("AddPerson", p.Name)
		m[p.Name] = p.Message
		defer r.Body.Close()
		w.WriteHeader(http.StatusCreated)
}

func AddPersonPhoto(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func DeletePeople(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func DeletePersonPhoto(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func FindPeople(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")	
		var ret = []Default{}
		for n, v := range m {
		    ret = append(ret, Default{n, v})
		}
	        js, err := json.Marshal(ret)
		if err != nil {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
		    return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		w.WriteHeader(http.StatusOK)
}

func FindPersonByName(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		uriSegments := strings.Split(r.URL.Path,"/")
		n := uriSegments[2]
		log.Println("FindPersonByName:", n)
		var ret = [1]Default{}
		ret[0] = Default{n, m[n]}
	        js, err := json.Marshal(ret)
		if err != nil {
		    http.Error(w, err.Error(), http.StatusInternalServerError)
		    return
		}		
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		w.WriteHeader(http.StatusOK)
}

func GetPersonPhoto(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func PeoplePersonNamePatch(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func PeoplePersonNamePhotosGet(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

func RootGet(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
}

