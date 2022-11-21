package test

import (
	m "cakestore/models"
	repo "cakestore/repo"
	"context"
	"fmt"
	"log"
	"testing"
)

func TestGetListCake(t *testing.T) {
	db := repo.GetConn()

	var (
		limit  int = 10
		offset int = 0
		ctx        = context.Background()
	)

	rows, err := db.QueryContext(ctx, repo.QueryGetListCake, limit, offset)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cakes := []m.Cake{}

	for rows.Next() {
		var c m.Cake

		err = rows.Scan(&c.ID, &c.Title, &c.Desc, &c.Rating, &c.Image, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			panic(err)
		}

		cakes = append(cakes, c)
	}

	fmt.Println(cakes)
}

func TestGetDetailCake(t *testing.T) {
	db := repo.GetConn()

	var (
		id   int = 1
		cake m.Cake
	)

	err := db.QueryRow(repo.QueryGetDetailCake, id).Scan(&cake.ID, &cake.Title, &cake.Desc, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err != nil {
		panic(err)
	}

	fmt.Println(cake)
	return
}

func TestAddCake(t *testing.T) {
	db := repo.GetConn()
	defer db.Close()
	ctx := context.Background()

	_, err := db.ExecContext(ctx, repo.QueryInsertCake)
	if err != nil {
		panic(err)
	}

	fmt.Println("success !")

}

func TestUpdateCake(t *testing.T) {
	db := repo.GetConn()
	ctx := context.Background()

	defer db.Close()

	var cake m.Cake
	var id int = 2

	err := db.QueryRow(repo.QueryGetDetailCake, id).Scan(&cake.ID, &cake.Title, &cake.Desc, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err != nil {
		panic(err)
	}

	if cake.ID == 0 {
		log.Print("cake not found")
		return
	}

	cake.Title = "update 2"

	_, err = db.ExecContext(ctx,
		repo.QueryUpdateCake,
		cake.Title,
		cake.Desc,
		cake.Rating,
		cake.Image,
		cake.ID,
	)
	if err != nil {
		panic(err)
	}
}

func TestDeleteCake(t *testing.T) {
	db := repo.GetConn()
	ctx := context.Background()

	defer db.Close()

	var r m.Cake
	var id int = 2

	err := db.QueryRow(repo.QueryGetDetailCake, id).Scan(&r.ID, &r.Title, &r.Desc, &r.Rating, &r.Image, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		panic(err)
	}

	if r.ID == 0 {
		log.Print("cake not found")
		return
	}

	_, err = db.ExecContext(ctx,
		repo.QuerySoftDeleteCake,
		r.ID,
	)
	if err != nil {
		panic(err)
	}
}
