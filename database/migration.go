package database

import "database/sql"

// InitializeDatabaseObjects initializes database objects like tables
func InitializeDatabaseObjects() error {
	db, err := GetConnection()
	defer db.Close()

	// Create Tables
	err = createSnippetTable(db)
	if err != nil {
		return err
	}

	err = createTagTable(db)
	if err != nil {
		return err
	}

	err = createSnippetTagTable(db)
	if err != nil {
		return err
	}

	err = createSnippetSearchTable(db)
	if err != nil {
		return err
	}

	return nil
}

// createSnippetTable creates snippet table
func createSnippetTable(connection *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS snippets (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		code TEXT NOT NULL,
		language TEXT,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := connection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// createTagTable creates tag table
func createTagTable(connection *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);`

	_, err := connection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// createSnippetTagTable creates snippet-tag table
func createSnippetTagTable(connection *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS snippet_tags(
    	snippet_id INTEGER,
		tag_id INTEGER,
		FOREIGN KEY (snippet_id) REFERENCES snippets(id) ON DELETE CASCADE,
		FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE,
		PRIMARY KEY (snippet_id, tag_id)
	);`

	_, err := connection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// createSnippetSearchTable creates snippet-search table
func createSnippetSearchTable(connection *sql.DB) error {
	query := `CREATE VIRTUAL TABLE IF NOT EXISTS snippet_search USING fts5( 
	title, code, description, snippet_id, tokenize="porter ascii");`
	_, err := connection.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
