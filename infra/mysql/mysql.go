package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type connect struct {
	Database *sql.DB
}

var App connect

func InitMysql() *sql.DB {
	godotenv.Load()
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	database := os.Getenv("DB")
	host := os.Getenv("DATABASE_HOST")

	log.Println(username, password, database, host)

	dsn := fmt.Sprintf("%v:%v@tcp(%s)/%v?parseTime=true", username, password, host, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	log.Println("masuk")

	// migrations := &migrate.FileMigrationSource{
	// 	Dir: "./migration",
	// }

	// log.Println("masuk2")

	// n, err := migrate.Exec(db, "mysql", migrations, migrate.Up)
	// if err != nil {
	// 	log.Println("error migratin", err)
	// }
	// fmt.Printf("Applied %d migrations!\n", n)
	CreateTableCake(db)

	return db
}

func CreateTableCake(db *sql.DB) {
	_, err := db.Exec(`DROP TABLE IF EXISTS cake;`)
	if err != nil {
		log.Println("error drop table", err)
	}

	_, err = db.Exec(`CREATE TABLE cake (
		id int NOT NULL AUTO_INCREMENT,
		title text NULL,
		description text NULL,
		rating numeric NULL,
		image varchar(255) NULL,
		created_at timestamp NULL DEFAULT now(),
		updated_at timestamp NULL DEFAULT now(),
		deleted_at timestamp NULL,
		PRIMARY KEY (id)
	);`)
	if err != nil {
		log.Println("error create table", err)
	}
}
