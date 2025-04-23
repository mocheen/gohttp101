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
		todos, _ := s.db.GetAll()
		ret, _ := json.Marshal(todos)
		w.Write([]byte(ret))
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
		err = s.db.Create(data.Todo{
			Id:    0,
			Name:  todo.Name,
			State: todo.State,
		})
		if err != nil {
			return
		}

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
		err = s.db.Update(todo)

	})

}
