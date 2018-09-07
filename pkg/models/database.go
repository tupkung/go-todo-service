package models

import "database/sql"

//Database type
type Database struct {
	*sql.DB
}

//LatestTasks for getting Task objects
func (db *Database) LatestTasks() (Tasks, error) {
	return Tasks{
		&Task{},
	}, nil
}

//Migrate data
func (db *Database) Migrate() error {
	for _, stmt := range getMigrationCommands() {
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func getMigrationCommands() []string {
	return []string{
		// migration_1
		`
			CREATE DATABASE IF NOT EXISTS todo_db DEFAULT CHARACTER SET = 'utf8' COLLATE 'utf8_general_ci';
		`,
		`
			USE todo_db;
		`,
		`
			CREATE TABLE IF NOT EXISTS task (
				id varchar(50) NOT NULL,
				title varchar(255),
				complete tinyint(1),
				create_time int(13),
				update_time int(13),
				PRIMARY KEY (id)
			) ENGINE=InnoDB;
		`,
	}
}
