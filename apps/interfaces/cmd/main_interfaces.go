package main

import "lark/apps/interfaces/internal/server"

func main() {
	s := server.NewServer()
	s.Run()
}
