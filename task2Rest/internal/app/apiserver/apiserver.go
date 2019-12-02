package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/trzhensimekh/cursesGo/task2Rest/store"
	"io"
	"net/http"
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
	s.router.HandleFunc("/hello",s.HandleHello())
}

func (s *APIServer) configureStore() error {
st := store.New(s.config.Store)
if err:=st.Open(); err!=nil{
	return err
}
s.store = st
return nil
}

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Hello")
}
}

