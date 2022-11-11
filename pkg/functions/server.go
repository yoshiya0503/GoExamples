package functions

import "fmt"

type Handler interface {
	Apply()
}

type Server struct {
	h map[string]Handler
}

func NewServer() Server {
	h := make(map[string]Handler)
	return Server{h: h}
}

func (s *Server) Handle(url string, h Handler) {
	s.h[url] = h
}

func (s *Server) Listen() {
	for url, handler := range s.h {
		fmt.Printf("[%s]: ", url)
		handler.Apply()
	}
}

func RunServer() {
	s := NewServer()
	var get Function = func() { fmt.Println("call get") }
	var post Function = func() { fmt.Println("call post") }
	s.Handle("/get", get)
	s.Handle("/post", post)
	s.Handle("/put", Function(func() { fmt.Println("call put") }))
	s.Listen()
}
