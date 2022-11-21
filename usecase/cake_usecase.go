package usecase

import (
	m "cakestore/models"
	repo "cakestore/repo"
)

func GetListCake(limit, page int) (r []m.Cake, err error) {
	r, err = repo.GetListCake(limit, page)
	return r, err
}

func AddCake(cake m.Cake) (r m.Cake, err error) {
	r, err = repo.AddCake(cake)
	return r, err
}

func DetailCake(id int) (r m.Cake, err error) {
	r, err = repo.DetailCake(id)
	return r, err
}

func UpdateCake(id int, cake m.Cake) (r m.Cake, err error) {
	r, err = repo.UpdateCake(id, cake)
	return r, err
}

func DeleteCake(id int) (err error) {
	err = repo.DeleteCake(id)
	return err
}
