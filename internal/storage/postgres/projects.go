package postgres

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"creative_at"`
}

func (s *Storagepostgres) CreateProject(title, description string) error {
	_, err := s.db.Exec(`INSERT INTO projects (title, description, created_at) VALUES ($1, $2, now())`, title, description)
	return err
}

func (s *Storagepostgres) GetProjects() ([]Project, error) {
	rows, err := s.db.Query(`SELECT id, title, description, created_at FROM projects`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []Project
	for rows.Next() {
		var p Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.CreatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}
