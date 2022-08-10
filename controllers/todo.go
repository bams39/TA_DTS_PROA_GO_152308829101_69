package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ichtrojan/go-todo/config"
	"github.com/ichtrojan/go-todo/models"
)

var (
	id           int
	petugas      string
	tugas        string
	tenggatWaktu string
	completed    int
	view         = template.Must(template.ParseFiles("./views/index.html"))
	database     = config.Database()
)

func Show(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT * FROM daftarTugas`)

	if err != nil {
		fmt.Println(err)
	}

	var todos []models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &petugas, &tugas, &tenggatWaktu, &completed)

		if err != nil {
			fmt.Println(err)
		}

		todo := models.Todo{
			Id:           id,
			Petugas:      petugas,
			Tugas:        tugas,
			TenggatWaktu: tenggatWaktu[0:10],
			Completed:    completed,
		}

		todos = append(todos, todo)
	}

	data := models.View{
		Todos: todos,
	}

	_ = view.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

	tugas := r.FormValue("tugas")
	petugas := r.FormValue("petugas")
	tenggatWaktu := r.FormValue("tenggatWaktu")

	_, err := database.Exec(`INSERT INTO daftarTugas (tugas, petugas, tenggatWaktu) VALUES (?, ?, ?)`, tugas, petugas, tenggatWaktu)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {

	updateId := r.FormValue("updateId")
	updatePetugas := r.FormValue("updatePetugas")
	updateTugas := r.FormValue("updateTugas")
	updateTenggatWaktu := r.FormValue("updateTenggatWaktu")

	// fmt.Println(updateId)
	// fmt.Println(updatePetugas)
	// fmt.Println(updateTugas)
	// fmt.Println(updateTenggatWaktu)

	_, err := database.Exec(`update daftarTugas set petugas = ?, tugas = ?, tenggatWaktu = ? where id = ?`, updatePetugas, updateTugas, updateTenggatWaktu, updateId)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`DELETE FROM daftarTugas WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Complete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE daftarTugas SET status = 1 WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}
