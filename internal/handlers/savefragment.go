package handlers

import (
	"github.com/go-chi/chi"
	"internshipApplicationTemplate/internal/imageservice"
	"io"
	"log"
	"net/http"
)

func SaveFragment(w http.ResponseWriter, r *http.Request) {
	log.Println("in SaveFragment handler")
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Println("id not found")
		http.Error(w, "id is empty", http.StatusNotFound)
		return
	}
	log.Println("id=", id)

	x := r.URL.Query().Get("x")
	if x == "" {
		log.Println("x is empty")
		http.Error(w, "x is empty", http.StatusBadRequest)
		return
	}
	log.Println("x=", x)

	y := r.URL.Query().Get("y")
	if y == "" {
		log.Println("y is empty")
		http.Error(w, "y is empty", http.StatusBadRequest)
		return
	}
	log.Println("y=", y)

	width := r.URL.Query().Get("width")
	if width == "" {
		log.Println("width is empty")
		http.Error(w, "width is empty", http.StatusBadRequest)
		return
	}
	log.Println("width=", width)

	height := r.URL.Query().Get("height")
	if height == "" {
		log.Println("height is empty")
		http.Error(w, "height is empty", http.StatusBadRequest)
		return
	}
	log.Println("height=", height)

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("reading request body error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(string(data))

	err = imageservice.SaveFragment(id, x, y, width, height, data)
	if err != nil {
		log.Println("save fragment error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

}
