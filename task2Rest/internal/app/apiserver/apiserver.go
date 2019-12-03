package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trzhensimekh/cursesGo/task2Rest/store"
	"net/http"
	_ "encoding/json" // ...
)

//APIServer ...
type APIServer struct{
	config *Config
	router *mux.Router
	store *store.Store
}

// New ...
func New(config*Config)*APIServer {
	return &APIServer{
		config: config,
		router:mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer)Start() error {
	s.configureRouter()
	if err :=s.configureStore(); err!=nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr,s.router)
}

func (s *APIServer) configureRouter(){
	s.router.HandleFunc("/users",s.HandleUsers())
}

func (s *APIServer) configureStore() error {
st := store.New(s.config.Store)
if err:=st.Open(); err!=nil{
	return err
}
s.store = st
return nil
}

func (s *APIServer) HandleUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		users,_:=s.store.User().GetUsers();
		json.NewEncoder(w).Encode(users)
	/*
		if err!=nil{
			io.WriteString(w, "error")
		}
		for u := range users {

			io.WriteString(w, string(u))
		}
	 */

	}
}

