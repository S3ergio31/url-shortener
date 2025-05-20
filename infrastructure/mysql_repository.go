package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/S3ergio31/url-shortener/domain"
	"github.com/go-sql-driver/mysql"
)

type MySqlRepository struct {
}

func (repo MySqlRepository) Save(short domain.Short) {
	_, err := repo.FindByShortCode(short.ShortCode)
	if err != nil {
		log.Println("MySqlRepository save -> insert")
		insert(short)
	} else {
		log.Println("MySqlRepository save -> update")
		update(short)
	}
}

func insert(short domain.Short) {
	query := "INSERT INTO shorts (id,url,short_code,created_at,updated_at,access_count) values(?,?,?,?,?,?)"
	_, err := connect().Exec(query, short.Id, short.Url, short.ShortCode, short.CreatedAt, short.UpdatedAt, short.AccessCount)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(query)
}

func update(short domain.Short) {
	query := "UPDATE shorts SET url=?,updated_at=?,access_count=? WHERE id = ?"
	_, err := connect().Exec(query, short.Url, short.UpdatedAt, short.AccessCount, short.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(query)
}

func (repo MySqlRepository) FindByShortCode(shortCode string) (*domain.Short, *domain.ShortNotFound) {
	var short domain.Short
	query := "SELECT * FROM shorts WHERE short_code = ?"
	rows, err := connect().Query(query, shortCode)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&short.Id, &short.Url, &short.ShortCode, &short.CreatedAt, &short.UpdatedAt, &short.AccessCount); err != nil {
			log.Fatal(err)
		}
		log.Println(query)
		return &short, nil
	}

	return nil, &domain.ShortNotFound{Key: shortCode}
}

func (repo MySqlRepository) Delete(shortCode string) {
	query := "DELETE FROM shorts WHERE short_code = ?"
	_, err := connect().Exec(query, shortCode)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(query)
}

func connect() *sql.DB {
	config := mysql.NewConfig()
	config.User = os.Getenv("DATABASE_USER")
	config.Passwd = os.Getenv("DATABASE_PASSWORD")
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"))
	config.DBName = os.Getenv("DATABASE_NAME")
	config.ParseTime = true

	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")

	return db
}

var mySqlRepository *MySqlRepository

func BuildMySqlRepository() MySqlRepository {
	if mySqlRepository == nil {
		mySqlRepository = &MySqlRepository{}
	}

	return *mySqlRepository
}
