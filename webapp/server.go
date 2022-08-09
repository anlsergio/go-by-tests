package poker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

type PlayerServer struct {
	Store PlayerStore
	http.Handler
	template *template.Template
	Game
}

const jsonContentType = "application/json"

func (s *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(s.Store.GetLeague())
}

func (s *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		s.showScore(w, player)
	case http.MethodPost:
		s.processWin(w, player)
	}
}

func (s *PlayerServer) playGame(w http.ResponseWriter, r *http.Request) {
	s.template.Execute(w, nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	ws := newPlayerServerWS(w, r)

	numberOfPlayersMsg := ws.WaitForMsg()
	numberOfPlayers, _ := strconv.Atoi(numberOfPlayersMsg)

	s.Game.Start(numberOfPlayers, ws) // todo: don't discard the blind messages!

	winnerMsg := ws.WaitForMsg()

	s.Game.Finish(winnerMsg)
}

func (s *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := s.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *PlayerServer) processWin(w http.ResponseWriter, player string) {
	s.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

const htmlTemplatePath = "game.html"

func NewPlayerServer(store PlayerStore, game Game) (*PlayerServer, error) {
	s := new(PlayerServer)

	tpl, err := template.ParseFiles(htmlTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("problem loading template %s", err.Error())
	}

	s.template = tpl
	s.Store = store
	s.Game = game

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(s.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(s.playersHandler))
	router.Handle("/game", http.HandlerFunc(s.playGame))
	router.Handle("/ws", http.HandlerFunc(s.webSocket))

	s.Handler = router

	return s, nil
}

type playerServerWS struct {
	*websocket.Conn
}

func (w *playerServerWS) WaitForMsg() string {
	_, msg, err := w.ReadMessage()
	if err != nil {
		log.Printf("failed to read from websocket %v\n", err)
	}

	return string(msg)
}

func (w *playerServerWS) Write(p []byte) (n int, err error) {
	if err = w.WriteMessage(websocket.TextMessage, p); err != nil {
		return 0, err
	}

	return len(p), nil
}

func newPlayerServerWS(w http.ResponseWriter, r *http.Request) *playerServerWS {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("failed to upgrade connection to WebSocket: ", err)
	}

	return &playerServerWS{conn}
}
