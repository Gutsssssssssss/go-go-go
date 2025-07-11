package database

type Database interface {
	SaveGameRecord(id string, data any) error
	GetGameRecord(id string) (any, error)
}

type MemoryDatabase struct {
	GameRecord map[string]any
}

func NewMemoryDatabase() *MemoryDatabase {
	return &MemoryDatabase{
		GameRecord: make(map[string]any),
	}
}

func (db *MemoryDatabase) SaveGameRecord(id string, data any) error {
	db.GameRecord[id] = data
	return nil
}

func (db *MemoryDatabase) GetGameRecord(id string) (any, error) {
	if data, ok := db.GameRecord[id]; ok {
		return data, nil
	}
	return nil, nil
}
