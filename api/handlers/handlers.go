package handlers

import "net/http"

type Stack struct {
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Root(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Stack) HealtCheck(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Stack) Login(w http.ResponseWriter, r *http.Request) {
	return
}

func (s *Stack) Logout(w http.ResponseWriter, r *http.Request) {
	return
}
