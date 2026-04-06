package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Todo struct {
	storage *Storage
}

func (t Todo) Todo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	all := t.storage.ReadAll()
	buf := new(bytes.Buffer)
	for _, item := range *all {
		v, err := json.Marshal(item)
		if err != nil {
			continue
		}

		buf.Write(v)
	}

	_, err2 := w.Write(buf.Bytes())
	if err2 != nil {
	}

	return
}

func (t Todo) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	v, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createRequest := CreateRequest{}
	err2 := json.Unmarshal(v, &createRequest)
	if err2 != nil {
		// log.Fatal(err)
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	if createRequest.Value == "" {
		http.Error(w, "Value is empty", http.StatusBadRequest)
		return
	}

	item := t.storage.Write(createRequest.Value)
	val, err3 := json.Marshal(item)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err4 := w.Write(val)
	if err4 != nil {
		http.Error(w, err4.Error(), http.StatusBadRequest)
		return
	}
}

func (t Todo) Read(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	item, err2 := t.storage.Read(id)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusNotFound)
		return
	}

	val, err3 := json.Marshal(item)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}

	_, err4 := w.Write(val)
	if err4 != nil {
		// http.Error(w, err4.Error(), http.StatusBadRequest)
		return
	}
}

func (t Todo) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	v, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateRequest := UpdateRequest{}
	err2 := json.Unmarshal(v, &updateRequest)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

	id, err3 := strconv.Atoi(r.PathValue("id"))
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
	}

	item, err4 := t.storage.Update(id, updateRequest.Value)
	if err4 != nil {
		http.Error(w, err4.Error(), http.StatusBadRequest)
		return
	}

	res, err5 := json.Marshal(item)
	if err5 != nil {
		http.Error(w, err5.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err6 := w.Write(res)
	if err6 != nil {
	}
}

func (t Todo) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	t.storage.Remove(id)

	w.WriteHeader(http.StatusOK)
	_, err2 := w.Write([]byte("{}"))
	if err2 != nil {
	}
	return
}

type CreateRequest struct {
	Value string `json:"value"`
}

type UpdateRequest struct {
	Value string `json:"value"`
}
