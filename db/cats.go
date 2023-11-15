package db

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
)

type Cat struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Age           int     `json:"age"`
	TimeInShelter float64 `json:"timeInShelter"`
	Adopted       string  `json:"adopted"`
	Description   string  `json:"description"`
	Photo         []byte  `json:"photo"`
}

func GetAllCats(connection string) ([]Cat, error) {
	conn, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(conn, context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM Cat;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Cat
	for rows.Next() {
		user := Cat{}
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.TimeInShelter, &user.Adopted, &user.Description, &user.Photo)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func InsertCat(cat Cat, connection string) (error, int) {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		return err, 0
	}
	defer db.Close()

	var id int
	err = db.QueryRow("SELECT nextval('cat1_id_seq')").Scan(&id)
	if err != nil {
		return err, id
	}

	stmt, err := db.Prepare("INSERT INTO Cat(id, name, age, timeInShelter, adopted, description, photo) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err, id
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, cat.Name, cat.Age, cat.TimeInShelter, cat.Adopted, cat.Description, cat.Photo)
	if err != nil {
		return err, id
	}

	return nil, id
}
