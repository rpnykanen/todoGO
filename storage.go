package main

import (
	"errors"
	"regexp"
	"time"
)

type Storage struct {
	store map[int]Item
}

type Item struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Value     string    `json:"value"`
}

func NewStorage() *Storage {
	return &Storage{make(map[int]Item)}
}

func (s *Storage) Write(value string) Item {
	item := Item{len(s.store), time.Now(), s.clean(value)}
	length := len(s.store)
	(s.store)[length] = item
	return item
}

func (s *Storage) Update(id int, value string) (Item, error) {
	item, ok := s.store[id]
	if ok == false {
		return Item{}, errors.New("not found")
	}

	item.Value = s.clean(value)
	s.store[id] = item
	return item, nil
}

func (s *Storage) ReadAll() *map[int]Item {
	return &s.store
}

func (s *Storage) Read(id int) (*Item, error) {
	x, ok := s.store[id]

	if ok == false {
		return nil, errors.New("not found")
	}

	return &x, nil
}

func (s *Storage) Remove(id int) {
	delete(s.store, id)
}

func (s *Storage) clean(value string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9 !?()/,.-]+")
	return regex.ReplaceAllString(value, "")
}
