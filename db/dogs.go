package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type Dog struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Age           int     `json:"age"`
	TimeInShelter float64 `json:"timeInShelter"`
	Adopted       string  `json:"adopted"`
	Photo         string  `json:"photo"`
	Description   string  `json:"description"`
}

func GetAllDogs(connection string) ([]Dog, error) {
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

	rows, err := conn.Query(context.Background(), "SELECT * FROM Dog")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []Dog
	for rows.Next() {
		dog := Dog{}
		err = rows.Scan(&dog.ID, &dog.Name, &dog.Age, &dog.TimeInShelter, &dog.Adopted, &dog.Photo, &dog.Description)
		if err != nil {
			return nil, err
		}
		dogs = append(dogs, dog)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return dogs, nil
}
