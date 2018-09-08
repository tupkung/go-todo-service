package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

//Database type
type Database struct {
	*sql.DB
}

//LatestTasks for getting Task objects
func (db *Database) LatestTasks() (Tasks, error) {
	stmt := `SELECT id, title, complete, create_time, update_time 
			FROM task 
			ORDER BY create_time DESC;`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := Tasks{}

	for rows.Next() {
		t := &Task{}

		err := rows.Scan(&t.ID, &t.Title, &t.IsComplete, &t.Created, &t.Updated)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

//InsertTask
func (db *Database) InsertTask(title string) (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	epoch := getEpoch()
	stmt := `INSERT INTO task (id, title, complete, create_time, update_time) 
	VALUES(?, ?, 0, ?, ?)`

	result, err := db.Exec(stmt, uid.String(), title, epoch, epoch)
	if err != nil {
		return "", err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return "", err
	}

	return uid.String(), nil
}

//UpdateTask
func (db *Database) UpdateTask(uid string, title string, complete bool) error {
	epoch := getEpoch()
	stmt := `UPDATE task 
	SET title=?,complete=?,update_time=? 
	WHERE id=?`

	result, err := db.Exec(stmt, title, complete, epoch, uid)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

//DeleteTask
func (db *Database) DeleteTask(uid string) error {
	stmt := `DELETE FROM task WHERE id = ?`

	result, err := db.Exec(stmt, uid)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
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

func getEpoch() int64 {
	currentTime := time.Now()
	epoch := currentTime.UnixNano() / 1000000
	return epoch
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
				create_time bigint(13),
				update_time bigint(13),
				PRIMARY KEY (id)
			) ENGINE=InnoDB;
		`,
	}
}
