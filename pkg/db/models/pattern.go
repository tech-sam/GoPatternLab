package models

import (
	"database/sql"
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

func (p *Pattern) Save(db *sql.DB) error {
	query := `
        INSERT INTO patterns (name, description)
        VALUES (?, ?)
        RETURNING id, created_at`

	return db.QueryRow(query, p.Name, p.Description).Scan(&p.ID, &p.CreatedAt)
}

func GetPattern(db *sql.DB, id int64) (*Pattern, error) {
	p := &Pattern{}
	query := `SELECT id, name, description, created_at FROM patterns WHERE id = ?`

	err := db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Problem methods
func (p *Problem) Save(db *sql.DB) error {
	query := `
        INSERT INTO problems (pattern_id, name, description, solution)
        VALUES (?, ?, ?, ?)
        RETURNING id, created_at`

	return db.QueryRow(query, p.PatternID, p.Name, p.Description, p.Solution).
		Scan(&p.ID, &p.CreatedAt)
}

func GetProblem(db *sql.DB, id int64) (*Problem, error) {
	p := &Problem{}
	query := `SELECT id, pattern_id, name, description, solution, created_at 
              FROM problems WHERE id = ?`

	err := db.QueryRow(query, id).Scan(
		&p.ID, &p.PatternID, &p.Name, &p.Description, &p.Solution, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}
