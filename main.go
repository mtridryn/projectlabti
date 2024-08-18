package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const PORT = 22197

func main() {
	http.HandleFunc("/api/mahasiswa", user)
	http.Handle("/", http.FileServer(http.Dir("polymer")))
	fmt.Printf("Jalankan web pada Port %d. url adalah http://[::1]:%d/", PORT, PORT)
	fmt.Println()
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
}

type labti struct {
	Nama     string `json:"nama_mahasiswa"`
	Angkatan string `json:"angkatan_mahasiswa"`
	Foto     string `json:"foto_mahasiswa"`
}

var data_mahasiswa = []labti{
	{
		Nama:     "Kageyama Tobio",
		Angkatan: "2020",
		Foto:     "img/gambar1.jpg",
	},
	{
		Nama:     "Choi Yeonjun",
		Angkatan: "2021",
		Foto:     "img/gambar2.jpg",
	},
	{
		Nama:     "Mutiara Indriyani",
		Angkatan: "2022",
		Foto:     "img/gambar3.jpg",
	},
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	if r.Method == http.MethodGet {

		result, err := json.Marshal(data_mahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
}
