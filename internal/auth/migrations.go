package auth

func (s *SqlDB) Migrate() error {
	// Drop database
	_, err := s.DB.Exec("Drop DATABASE IF EXISTS auth")
	if err != nil {
		return err
	}

	// Create the database if it doesn't exist
	_, err = s.DB.Exec("CREATE DATABASE IF NOT EXISTS auth")
	if err != nil {
		return err
	}

	// Switch to the 'auth' database
	_, err = s.DB.Exec("USE auth")
	if err != nil {
		return err
	}

	_, err = s.DB.Exec(`
		CREATE TABLE IF NOT EXISTS  user (
			uuid int PRIMARY KEY,
			email varchar(32) NOT NULL,
			phone int NOT NULL,
			hashed_password varchar(32) NOT NULL,
			country varchar(16) NOT NULL,
			is_verified boolean DEFAULT false
		);
	`)
	if err != nil {
		return err
	}
	return nil
}
