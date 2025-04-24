package server

import (
	"encoding/json"
	"fmt"
	"gohttp101/data"
	"io"
	"net/http"
)

func Get(s *Server) {
	s.mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		var todos []data.Todo
		s.db.Find(&todos)

		ret, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(ret)
	})
}

func Post(s *Server) {
	s.mux.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		body, _ := io.ReadAll(request.Body)
		fmt.Println(body)
		todo := data.Todo{}
		err := json.Unmarshal(body, &todo)
		if err != nil {
			return
		}
		s.db.Create(&todo)

	})
}

func Update(s *Server) {
	s.mux.HandleFunc("/update", func(writer http.ResponseWriter, request *http.Request) {
		body, _ := io.ReadAll(request.Body)
		fmt.Println(body)
		todo := data.Todo{}
		err := json.Unmarshal(body, &todo)
		if err != nil {
			return
		}
		s.db.Model(&todo).Updates(todo)

	})

}
