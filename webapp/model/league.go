package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}

	return nil
}

func NewLeague(r io.Reader) (League, error) {
	var league League

	err := json.NewDecoder(r).Decode(&league)
	if err != nil {
		err = fmt.Errorf("could not parse league, %v", err)
	}

	return league, err
}
