package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/yanmoyy/go-go-go/internal/game"
)

func NewDB(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateGameRecord(db *sql.DB, sessionID string, records []game.Event) error {
	data, err := json.Marshal(records)
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO game_records (session_id, records) VALUES ($1, $2)", sessionID, data)
	if err != nil {
		return fmt.Errorf("failed to create game record: %v", err)
	}
	return nil
}
