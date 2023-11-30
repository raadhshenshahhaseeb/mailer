package store

import "time"

type Emails struct {
	Email     string    `json:"email"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

func (s *store) InsertRecord(email string) error {
	// Prepare the insert statement
	s.sqlMutex.Lock()
	defer s.sqlMutex.Unlock()
	stmt, err := s.sqlDB.Prepare("INSERT INTO emails(email) VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the statement with the provided values
	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}

	return nil
}

func (s *store) Get() (*[]Emails, error) {
	emails := []Emails{}

	// Prepare the query
	rows, err := s.sqlDB.Query("SELECT id, email, created_at FROM emails")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var e Emails
		err := rows.Scan(&e.ID, &e.Email, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		emails = append(emails, e)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &emails, nil
}
