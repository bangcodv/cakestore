package repo

import (
	m "cakestore/models"
	"context"
	"errors"
	"log"
)

func GetListCake(limit, page int) (r []m.Cake, err error) {
	db := GetConn()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, QueryGetListCake, limit, page)
	if err != nil {
		log.Printf("error while get data : %s", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c m.Cake

		err = rows.Scan(&c.ID, &c.Title, &c.Desc, &c.Rating, &c.Image, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			log.Printf("error while scan : %s", err)
			return
		}

		r = append(r, c)
	}

	return r, err
}

func DetailCake(id int) (r m.Cake, err error) {
	db := GetConn()

	err = db.QueryRow(QueryGetDetailCake, id).Scan(&r.ID, &r.Title, &r.Desc, &r.Rating, &r.Image, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		log.Printf("error while scan : %s", err)
		return
	}

	return r, err
}

func AddCake(cake m.Cake) (r m.Cake, err error) {
	db := GetConn()

	defer db.Close()

	ctx := context.Background()
	_, err = db.ExecContext(ctx, QueryInsertCake, cake.Title, cake.Desc, cake.Rating, cake.Image)
	if err != nil {
		log.Printf("error while scan : %s", err)
		return
	}

	r = cake
	return r, err
}

func UpdateCake(id int, cake m.Cake) (r m.Cake, err error) {
	db := GetConn()
	ctx := context.Background()

	defer db.Close()

	err = db.QueryRow(QueryGetDetailCake, id).Scan(&r.ID, &r.Title, &r.Desc, &r.Rating, &r.Image, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		log.Printf("error query while : %s", err)
		return
	}

	if r.ID == 0 {
		err = errors.New("cake not found !")
		log.Println("cake not found", err)
		return
	}

	if cake.Title != "" {
		r.Title = cake.Title
	}

	if cake.Desc != "" {
		r.Desc = cake.Desc
	}

	if cake.Rating != 0 {
		r.Rating = cake.Rating
	}

	if cake.Image != "" {
		r.Image = cake.Image
	}

	_, err = db.ExecContext(ctx,
		QueryUpdateCake,
		r.Title,
		r.Desc,
		r.Rating,
		r.Image,
		r.ID,
	)
	if err != nil {
		log.Printf("error query while : %s", err)
		return
	}

	return r, err
}

func DeleteCake(id int) (err error) {
	db := GetConn()
	ctx := context.Background()

	defer db.Close()

	var r m.Cake
	err = db.QueryRow(QueryGetDetailCake, id).Scan(&r.ID, &r.Title, &r.Desc, &r.Rating, &r.Image, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		log.Printf("error query while : %s", err)
		return
	}

	if r.ID == 0 {
		err = errors.New("cake not found !")
		log.Println("cake not found", err)
		return
	}

	_, err = db.ExecContext(ctx,
		QuerySoftDeleteCake,
		r.ID,
	)
	if err != nil {
		log.Printf("error query while : %s", err)
		return
	}

	return err
}
