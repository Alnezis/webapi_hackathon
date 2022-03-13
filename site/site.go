package site

import (
	"api_hackathon/api"
	"api_hackathon/app"
	"api_hackathon/manager"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
)

func Site() {
	_ = app.CFG

	_cors := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	router := mux.NewRouter()
	router.HandleFunc("/manager/{email}/{password}", getManager).Methods("GET")

	router.HandleFunc("/placeholders", getPlaceholders).Methods("GET")

	cert := "/etc/letsencrypt/live/alnezis.qweri-craft.ru/fullchain.pem"
	key := "/etc/letsencrypt/live/alnezis.qweri-craft.ru/privkey.pem"

	handler := _cors.Handler(router)
	err := http.ListenAndServeTLS(":1920", cert, key, handler)
	if err == nil {
	} else {
		api.CheckErrInfo(err, "site")
		return
	}
	api.CheckErrInfo(err, "site 2")
}

type Result struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

func getManager(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	params := mux.Vars(r)
	res, err := manager.Get(params["email"], params["password"])
	if err != nil {
		json.NewEncoder(w).Encode(&Result{nil, err.Error()})
		return
	}
	json.NewEncoder(w).Encode(&Result{res, nil})
}

func getPlaceholders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	plan, err := ioutil.ReadFile("output.json") // filename is the JSON file to read
	api.CheckErrInfo(err, "errr json")

	var data interface{}
	err = json.Unmarshal(plan, &data)
	if err != nil {
		log.Println("Cannot unmarshal the json ", err)
	}

	json.NewEncoder(w).Encode(&data)
}
