package delivery

import (
	m "cakestore/models"
	uc "cakestore/usecase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to cake store !"))
}

func HandlerDetailCakes(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" || id == "0" {
		m := "Body Not Allowed"
		res := Response(http.StatusBadRequest, m, nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	cId, err := strconv.Atoi(id)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		log.Print(m)
		res := Response(http.StatusInternalServerError, m, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	d, err := uc.DetailCake(cId)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		log.Print(m)
		res := Response(http.StatusInternalServerError, m, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := Response(http.StatusOK, "success", d)
	json.NewEncoder(w).Encode(res)

}

func HandlerListCakes(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	if limit == "" || limit == "0" {
		limit = "10"
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	l, err := strconv.Atoi(limit)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		log.Print(m)
		res := Response(http.StatusInternalServerError, m, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	p, err := strconv.Atoi(page)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		log.Print(m)
		res := Response(http.StatusInternalServerError, m, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	p -= p

	d, err := uc.GetListCake(l, p)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		log.Print(m)
		res := Response(http.StatusInternalServerError, m, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return

	}

	res := Response(http.StatusOK, "success", d)
	json.NewEncoder(w).Encode(res)
}

func HandlerAddCake(w http.ResponseWriter, r *http.Request) {
	var res m.Response
	reqBody, _ := ioutil.ReadAll(r.Body)
	var cake m.Cake
	err := json.Unmarshal(reqBody, &cake)
	if err != nil {
		res = Response(http.StatusBadRequest, "body not allowed", nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	d, err := uc.AddCake(cake)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		res = Response(http.StatusInternalServerError, m, nil)
	} else {
		res = Response(http.StatusOK, "success", d)
	}

	json.NewEncoder(w).Encode(res)
}

func HandlerUpdateCake(w http.ResponseWriter, r *http.Request) {
	var res m.Response
	var cake m.Cake

	id := mux.Vars(r)["id"]
	if id == "" || id == "0" {
		m := "Body Not Allowed"
		res := Response(http.StatusBadRequest, m, nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &cake)
	if err != nil {
		res = Response(http.StatusBadRequest, "body not allowed", nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	cId, err := strconv.Atoi(id)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		log.Print(m)
		res := Response(http.StatusInternalServerError, m, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	d, err := uc.UpdateCake(cId, cake)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		res = Response(http.StatusInternalServerError, m, nil)
	} else {
		res = Response(http.StatusOK, "success", d)
	}

	json.NewEncoder(w).Encode(res)
}

func HandlerDeleteCake(w http.ResponseWriter, r *http.Request) {
	var res m.Response

	id := mux.Vars(r)["id"]
	if id == "" || id == "0" {
		m := "Body Not Allowed"
		res := Response(http.StatusBadRequest, m, nil)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	cId, err := strconv.Atoi(id)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		log.Print(m)
		res := Response(http.StatusInternalServerError, m, nil)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	err = uc.DeleteCake(cId)
	if err != nil {
		m := fmt.Sprintf("Some error occured. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		res = Response(http.StatusInternalServerError, m, nil)
	} else {
		res = Response(http.StatusOK, "success", nil)
	}

	json.NewEncoder(w).Encode(res)
}

func Response(status int, message string, data interface{}) (r m.Response) {
	r.Status = status
	r.Message = message
	r.Data = data

	return r
}
