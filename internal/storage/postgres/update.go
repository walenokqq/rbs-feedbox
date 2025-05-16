package postgres

import "database/sql"

func UpdateProject(db *sql.DB, id int, title, description string) error {
	_, err := db.Exec(`
	 UPDATE projects 
	 SET title =$1, description = $2
	 WHERE ID =$3`, title, description, id)
	return err

}
