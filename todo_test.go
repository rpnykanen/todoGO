package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTodo(t *testing.T) {
	storage := NewStorage()
	storage.Write("first item")
	storage.Write("second item")
	todo := Todo{storage}

	req := httptest.NewRequest(http.MethodGet, "/todo", nil)
	rr := httptest.NewRecorder()

	routeToTest("GET /todo", rr, req, todo.Todo)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code %v", status)
	}
}

func TestTodoCreatePostBadRequest(t *testing.T) {
	storage := NewStorage()
	todo := Todo{storage}

	req := httptest.NewRequest(http.MethodPost, "/todo/create", strings.NewReader(`{"valu": "bad value"}`))
	rr := httptest.NewRecorder()

	routeToTest("POST /todo/create", rr, req, todo.Create)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("handler: status code : %v", rr.Code)
	}
}

func TestTodoCreatePost(t *testing.T) {
	storage := NewStorage()
	todo := Todo{storage}

	req := httptest.NewRequest(http.MethodPost, "/todo/create", strings.NewReader(`{"value": "New TODO item"}`))
	rr := httptest.NewRecorder()

	routeToTest("POST /todo/create", rr, req, todo.Create)

	if rr.Code != http.StatusCreated {
		t.Errorf("handler: status code : %v", rr.Code)
		return
	}

	i := Item{}
	err := json.Unmarshal([]byte(rr.Body.String()), &i)
	if err != nil {
	}

	// check id & created_at ?
	if len(i.Value) == 0 {
		t.Errorf("Saved item missing value")
	}
}

func TestTodoRead(t *testing.T) {
	storage := NewStorage()
	storage.Write("New item should exist")
	todo := Todo{storage}

	req := httptest.NewRequest(http.MethodGet, "/todo/read/0", nil)
	rr := httptest.NewRecorder()

	routeToTest("GET /todo/read/{id}", rr, req, todo.Read)

	if rr.Code != http.StatusOK {
		t.Errorf("handler: status code : %v", rr.Code)
		t.Errorf("Error: %v", rr.Body.String())
		return
	}

	i := Item{}
	err := json.Unmarshal([]byte(rr.Body.String()), &i)
	if err != nil {
	}

	if len(i.Value) == 0 {
		t.Errorf("Saved item missing value")
	}
}

func TestTodoUpdate(t *testing.T) {
	storage := NewStorage()
	storage.Write("This will be overwritten")
	todo := Todo{storage}

	req := httptest.NewRequest(http.MethodPatch, "/todo/update/0", strings.NewReader(`{"id": 0, "value": "should have been updated"}`))
	rr := httptest.NewRecorder()

	routeToTest("PATCH /todo/update/{id}", rr, req, todo.Update)

	if rr.Code != http.StatusOK {
		t.Errorf("handler: status code : %v", rr.Code)
		t.Errorf("Error: %v", rr.Body.String())
		return
	}

	i := Item{}
	err := json.Unmarshal([]byte(rr.Body.String()), &i)
	if err != nil {
	}

	if len(i.Value) == 0 {
		t.Errorf("Saved item missing value")
	}
}

func TestTodoDelete(t *testing.T) {
	storage := NewStorage()
	storage.Write("This will be removed")
	todo := Todo{storage}

	req := httptest.NewRequest(http.MethodDelete, "/todo/delete/0", strings.NewReader(`{"id": 0}`))
	rr := httptest.NewRecorder()

	routeToTest("DELETE /todo/delete/{id}", rr, req, todo.Delete)

	if rr.Code != http.StatusOK {
		t.Errorf("handler: status code : %v", rr.Code)
		t.Errorf("Error: %v", rr.Body.String())
		return
	}

	items := storage.ReadAll()
	if len(*items) != 0 {
		t.Errorf("Deleted items not empty")
	}
}

func routeToTest(path string, writer *httptest.ResponseRecorder, request *http.Request, handler func(w http.ResponseWriter, r *http.Request)) {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	mux.ServeHTTP(writer, request)
}
