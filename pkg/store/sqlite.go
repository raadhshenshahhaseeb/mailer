package store

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
