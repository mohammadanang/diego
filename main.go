package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       int
	Username string
	FullName string
}

func NewUser(username, fullName string) User {
	return User{0, username, fullName}
}

type Repository interface {
	Store(user User) error
}

func NewMysqlRepository(sqlDB *sql.DB) Repository {
	return MysqlRepository{sqlDB: sqlDB}
}

type MysqlRepository struct {
	sqlDB *sql.DB
}

func (sql MysqlRepository) Store(user User) error {
	statement, err := sql.sqlDB.Prepare("INSERT INTO users (username, full_name) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Username, user.FullName); err != nil {
		return err
	}
	return nil
}

func newSql(mysqlUser, mysqlPassword string) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@/diegodb", mysqlUser, mysqlPassword)
	sqlDB, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return sqlDB, nil
}

type Service struct {
	userRepo Repository
}

func NewService(userRepo Repository) Service {
	return Service{userRepo}
}

func (service Service) Create(username, fullName string) error {
	user := NewUser(username, fullName)
	err := service.userRepo.Store(user)
	if err != nil {
		return err
	}

	return nil
}

var (
	mysqlHost     = "localhost"
	mysqlUser     = "root"
	mysqlPassword = "password"
)

func main() {
	sqlDb, err := newSql(mysqlUser, mysqlPassword)
	if err != nil {
		log.Fatalf("Error sql %v", err)
	}

	mysqlRepo := NewMysqlRepository(sqlDb)
	userService := NewService(mysqlRepo)
	if err := userService.Create("brood", "Brodn Way"); err != nil {
		log.Fatalf("Error service %v", err)
	}

	log.Printf("Done")
}
