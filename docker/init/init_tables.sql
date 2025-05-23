DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'feedback_status') THEN
        CREATE TYPE feedback_status AS ENUM ('new', 'processed', 'archived');
    END IF;
END$$;

		--VARCHAR предпочтительнее CHAR, тк мы не знаем заранее длину title
		--CHAR займет в памяти всю свою заранее определнную длину вне зависимости от фактической, в отличие от VARCHAR 
CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			title VARCHAR(64) NOT NULL,
			description TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS forms (
			id SERIAL PRIMARY KEY,
			project_id INTEGER NOT NULL,
			title VARCHAR(64) NOT NULL,
			description TEXT NOT NULL,
			schema JSONB NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (project_id) REFERENCES projects(id) ON DELETE CASCADE
		);

		CREATE TABLE IF NOT EXISTS feedback (
			id SERIAL PRIMARY KEY,
			form_id INTEGER NOT NULL,
			data JSONB NOT NULL,
			status feedback_status NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (form_id) REFERENCES forms(id) ON DELETE CASCADE
		);