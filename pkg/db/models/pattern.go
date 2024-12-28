package models

import (
	"github.com/tech-sam/GoPatternLab/pkg/db"
	"time"
)

type Pattern struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Problem struct {
	ID          int64     `json:"id"`
	PatternID   int64     `json:"pattern_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Solution    string    `json:"solution"`
	CreatedAt   time.Time `json:"created_at"`
}

// PatternModel wraps the database connection
type PatternModel struct {
	DB *db.DB
}

func NewPatternModel(db *db.DB) *PatternModel {
	return &PatternModel{DB: db}
}

func (m *PatternModel) Create(name, description string) error {
	query := `
        INSERT INTO patterns (name, description)
        VALUES (?, ?)
    `
	_, err := m.DB.Exec(query, name, description)
	return err
}

func (m *PatternModel) GetPatterns() ([]Pattern, error) {
	query := `
        SELECT id, name, description, created_at
        FROM patterns
        ORDER BY created_at DESC
    `
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patterns []Pattern
	for rows.Next() {
		var p Pattern
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		patterns = append(patterns, p)
	}
	return patterns, nil
}
