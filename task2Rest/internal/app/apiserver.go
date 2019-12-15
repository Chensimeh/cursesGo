package app

import (
	"encoding/json"
	_ "encoding/json" // ...
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/trzhensimekh/cursesGo/task2Rest/internal/data"
	"log"
	"net/http"
	"strconv"
	"time"
)

type APIServer struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
	}
}

func (s *APIServer) Start() error {
	s.configureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureRouter() {
	s.router.Use(s.headRequest)
	s.router.Use(s.logRequest)
	s.router.HandleFunc("/users", s.HandleUsers()).Methods("GET")
	s.router.HandleFunc("/users/{id}", s.FindUserById()).Methods("GET")
	s.router.HandleFunc("/users", s.UserCreaterHandler()).Methods("POST")
	s.router.HandleFunc("/users/{id}", s.UpdateUserById()).Methods("PUT")
	s.router.HandleFunc("/users/{id}", s.DeleteUserById()).Methods("DELETE")
	s.router.HandleFunc("/users/{id}/messages/{msg_id}", s.FindMsgById()).Methods("GET")
	s.router.HandleFunc("/users/{id}/messages", s.HandleUserMsgs()).Methods("GET")
	s.router.HandleFunc("/users/{id}/messages", s.MsgCreaterHandler()).Methods("POST")
	s.router.HandleFunc("/users/{id}/messages/{msg_id}", s.UpdateMsgById()).Methods("PUT")
	s.router.HandleFunc("/users/{id}/messages/{msg_id}", s.DeleteMsgById()).Methods("DELETE")
}

func (s *APIServer) HandleUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := (&data.User{}).GetUsers()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(users)
	}
}

func (s *APIServer) UserCreaterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		var user *data.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		err := user.CreateUser()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (s *APIServer) FindUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		user, err := (&data.User{Id: id}).FindByID()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (s *APIServer) UpdateUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var user *data.User
		_ = json.NewDecoder(r.Body).Decode(&user)
		id, _ := strconv.Atoi(params["id"])
		user.Id = id
		err := user.UpdatedByID()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (s *APIServer) DeleteUserById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		user := new(data.User)
		id, _ := strconv.Atoi(params["id"])
		user.Id = id
		err := user.DeleteByID()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(user)
	}
}

func (s *APIServer) FindMsgById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["msg_id"])
		msg, err := (&data.Message{Id: id}).FindMsgByID()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(msg)
	}
}

func (s *APIServer) HandleUserMsgs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		messeges, err := (&data.Message{Id: id}).GetUserMsg()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(messeges)
	}
}

func (s *APIServer) MsgCreaterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		var message *data.Message
		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])
		_ = json.NewDecoder(r.Body).Decode(&message)
		message.UserId = id
		err := message.CreateMsg()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(message)
	}
}

func (s *APIServer) UpdateMsgById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var msg *data.Message
		_ = json.NewDecoder(r.Body).Decode(&msg)
		id, _ := strconv.Atoi(params["msg_id"])
		msg.Id = id
		err := msg.UpdatedByID()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(msg)
	}
}

func (s *APIServer) DeleteMsgById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//	w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		msg := new(data.Message)
		_ = json.NewDecoder(r.Body).Decode(&msg)
		id, _ := strconv.Atoi(params["msg_id"])
		msg.Id = id
		err := msg.DeleteByID()
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(msg)
	}
}

func (s *APIServer) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(w, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *APIServer) headRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
