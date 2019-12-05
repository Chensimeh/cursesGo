package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trzhensimekh/cursesGo/task2Rest/model"
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
	s.router.HandleFunc("/users",s.HandleUsers()).Methods("GET")
	s.router.HandleFunc("/users",s.HandleUserCreater()).Methods("POST")
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
		w.Header().Set("Content-Type", "application/json")
		users,_:=s.store.User().GetUsers();
		json.NewEncoder(w).Encode(users)
	}
}

func (s *APIServer) HandleUserCreater() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		var user model.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		user,_=s.store.User().CreateUser(user);
		json.NewEncoder(w).Encode(user)
	}
}

