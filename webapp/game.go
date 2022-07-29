package poker

import "time"

type Game struct {
	store   PlayerStore
	alerter BlindAlerter
}

func (g *Game) Start(numberOfPlayers int) {
	const minimumBlindMinutes = 5

	blindIncrement := time.Duration(minimumBlindMinutes+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}

}

func (g *Game) Finish(winner string) {
	g.store.RecordWin(extractWinner(winner))
}

func NewGame(store PlayerStore, alerter BlindAlerter) *Game {
	return &Game{
		store,
		alerter,
	}
}
