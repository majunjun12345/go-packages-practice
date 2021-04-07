package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/dig"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "./example.db",
		Port:         "8000",
	}
}

func ConnectDatabase(config *Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}

// dao
// PersonRepository----------------------------------------------------------------PersonRepository
type PersonRepository struct {
	database *sql.DB
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}

func (repository *PersonRepository) FindAll() []*Person {
	rows, _ := repository.database.Query("SELECT id, name, age FROM people;")
	defer rows.Close()
	people := []*Person{}
	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)
		rows.Scan(&id, &name, &age)
		people = append(people, &Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}
	return people
}

// service----------------------------------------------------------------service
type PersonService struct {
	config     *Config
	repository *PersonRepository
}

func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
	return &PersonService{config: config, repository: repository}
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}
	return []*Person{}
}

// server---------------------------------------------------------------- server

type Server struct {
	config        *Config
	personService *PersonService
}

func NewServer(config *Config, service *PersonService) *Server {
	return &Server{
		config:        config,
		personService: service,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/people", s.people)
	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}
	fmt.Println("======listen:")
	httpServer.ListenAndServe()
}

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func main() {
	// // origin
	// config := NewConfig()
	// db, err := ConnectDatabase(config)
	// if err != nil {
	// 	panic(err)
	// }

	// personRepository := NewPersonRepository(db)
	// personService := NewPersonService(config, personRepository)
	// server := NewServer(config, personService)

	// server.Run()

	/*
		c := dig.New()创建一个实例，在这个实例上执行Provide注入对象构造器。
		Invoke依赖注入的构造器创建一个用户希望得到的对象。
	*/
	container := builderContanier()
	err := container.Invoke(func(server *Server) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}

func builderContanier() *dig.Container {
	container := dig.New()
	container.Provide(NewConfig)
	container.Provide(ConnectDatabase)
	container.Provide(NewPersonRepository)
	container.Provide(NewPersonService)
	container.Provide(NewServer)
	return container
}
