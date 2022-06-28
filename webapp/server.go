package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store PlayerStore
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, s.store.GetPlayerScore(player))
}

type PlayerStore interface {
	GetPlayerScore(name string) int
}

//func GetPlayerScore(player string) string {
//	if player == "Pepper" {
//		return "20"
//	}
//
//	if player == "Floyd" {
//		return "10"
//	}
//
//	return ""
//}
