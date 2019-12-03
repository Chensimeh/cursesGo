package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // ...
)

//Store ...
type Store struct {
	config *Config
	db *sql.DB
	userRep *UserRep
}


// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open ...
func (s *Store) Open() error {
db,err :=sql.Open("postgres",s.config.DatabaseURL)
if err!=nil{
	return err
}

if err:=db.Ping(); err!=nil{
	return err
}
s.db = db
	fmt.Println("Data base ping done!")
return nil
}

// Close ...
func (s *Store) Close() {
 s.db.Close()
}

//User ...
func (s *Store) User() *UserRep{
	if s.userRep!=nil{
		return s.userRep
	}

	s.userRep = &UserRep{
		Store: s,
	}
	return s.userRep
}
// store.User().Create() - create new User

