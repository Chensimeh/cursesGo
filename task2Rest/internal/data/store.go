package data

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
}

func (s *Store) Open() (*sql.DB, error) {
	config := NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", config)
	d, err := sql.Open("postgres", "host=localhost port=5432 dbname=mydb user=root password=root sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := d.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Data base ping done!")
	return d, nil
}
