package main

import (
	"todo-list/handlers"

	"github.com/labstack/echo/v4"
	bolt "go.etcd.io/bbolt"
)

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		return err
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	e := echo.New()

	db, err := initDB()
	if err != nil {
		e.Logger.Fatal("Error creating database table: ", err)
	}

	handlers.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
