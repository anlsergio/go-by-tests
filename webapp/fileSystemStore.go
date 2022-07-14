package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type FileSystemPlayerStore struct {
	Database *json.Encoder
	League   League
}

func (s *FileSystemPlayerStore) GetPlayerScore(name string) (wins int) {
	player := s.League.Find(name)
	if player != nil {
		return player.Wins
	}

	return wins
}

func (s *FileSystemPlayerStore) GetLeague() League {
	sort.Slice(s.League, func(i, j int) bool {
		return s.League[i].Wins > s.League[j].Wins
	})

	return s.League
}

func (s *FileSystemPlayerStore) RecordWin(name string) {
	player := s.League.Find(name)

	if player != nil {
		player.Wins++
	} else {
		s.League = append(s.League, Player{
			Name: name,
			Wins: 1,
		})
	}

	s.Database.Encode(s.League)
}

func FileSystemStoreFromFile(path string) (store *FileSystemPlayerStore, closeFunc func(), err error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open file %s, %w", path, err)
	}

	closeFunc = func() {
		db.Close()
	}

	store, err = NewFileSystemStore(db)
	if err != nil {
		return nil, nil, fmt.Errorf("problem creating player store, %w", err)
	}

	return store, closeFunc, nil
}

func NewFileSystemStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initPlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file, %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player playerStore from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{
		Database: json.NewEncoder(&tape{file}),
		League:   league,
	}, nil
}

func initPlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		initEmptyFile(file)
	}

	return nil
}

func initEmptyFile(file *os.File) {
	file.Write([]byte("[]"))
	file.Seek(0, 0)
}
