package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/trzhensimekh/cursesGo/task2Rest/model"
	"github.com/trzhensimekh/cursesGo/task2Rest/store"
	"log"
	"net/http"
	_ "encoding/json" // ...
	"strconv"
)


type APIServer struct{
	config *Config
	router *mux.Router
	store *store.Store
}


func New(config*Config)*APIServer {
	return &APIServer{
		config: config,
		router:mux.NewRouter(),
	}
}


func (s *APIServer)Start() error {
	s.configureRouter()
	if err :=s.configureStore(); err!=nil {
		return err
	}

	return http.ListenAndServe(s.config.BindAddr,s.router)
}

func (s *APIServer) configureRouter(){
	s.router.HandleFunc("/users",s.HandleUsers()).Methods("GET")
	s.router.HandleFunc("/users/{id}",s.FindUserById()).Methods("GET")
	s.router.HandleFunc("/users",s.UserCreaterHandler()).Methods("POST")
	s.router.HandleFunc("/users/{id}",s.UpdateUserById()).Methods("PUT")
	s.router.HandleFunc("/users/{id}",s.DeleteUserById()).Methods("DELETE")
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

func (s *APIServer) UserCreaterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		var user *model.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		err:=s.store.User().CreateUser(user);
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (s *APIServer) FindUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id,_:=strconv.Atoi(params["id"])
		user,err := s.store.User().FindByID(id);
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (s *APIServer) UpdateUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var user *model.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		id,_:=strconv.Atoi(params["id"])
		user.Id = id
		err := s.store.User().UpdatedByID(user);
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (s *APIServer) DeleteUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		user:= new(model.User)
		id, _ := strconv.Atoi(params["id"])
		user.Id = id
		err:= s.store.User().DeleteByID(user);
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}