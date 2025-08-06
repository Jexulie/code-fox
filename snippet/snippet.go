package snippet

import (
	"code-fox/database"
	"fmt"
	"time"
)

type Snippet struct {
	Id          int64
	Title       string
	Code        string
	Language    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	Tags        string
}

// NewSnippet returns snippet struct
func NewSnippet(title, code, language, description string) *Snippet {
	return &Snippet{
		Title:       title,
		Code:        code,
		Language:    language,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
}

// Create saves snippet
func (s *Snippet) Create() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	query := `INSERT INTO snippets (title, code, language, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	res, err := db.Exec(query, s.Title, s.Code, s.Language, s.Description, s.CreatedAt, s.UpdatedAt)
	if err != nil {
		return err
	}

	s.Id, _ = res.LastInsertId()

	defer db.Close()
	return nil
}

// Update updates snippet
func (s *Snippet) Update() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	now := time.Now()
	s.UpdatedAt = &now

	query := `UPDATE snippets SET title=?, code=?, language=?, description=?, updated_at=? WHERE id=?`

	_, err = db.Exec(query, s.Title, s.Code, s.Language, s.Description, s.UpdatedAt, s.Id)
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}

// Delete deletes snippet
func (s *Snippet) Delete() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	query := `DELETE FROM snippets WHERE id = ?`

	_, err = db.Exec(query, s.Id)
	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

// GetById gets snippet by id
func GetById(id int64) (*Snippet, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var snippet Snippet

	query := `SELECT s.*, IFNULL(GROUP_CONCAT(t.name, ','), '') Tags FROM snippets as s
         LEFT JOIN snippet_tags as st on st.snippet_id = s.id
         LEFT JOIN tags t on t.id = st.tag_id WHERE s.id = ?
		GROUP BY s.id, s.title, s.code, s.language, s.description, s.created_at, s.updated_at`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Code, &snippet.Language, &snippet.Description, &snippet.CreatedAt, &snippet.UpdatedAt, &snippet.Tags)
		if err != nil {
			return nil, err
		}
	}

	return &snippet, nil
}

// GetByTitle gets snippet by title
func GetByTitle(title string) (*Snippet, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var snippet Snippet

	query := `SELECT s.*, IFNULL(GROUP_CONCAT(t.name, ','), '') Tags FROM snippets as s
         LEFT JOIN snippet_tags as st on st.snippet_id = s.id
         LEFT JOIN tags t on t.id = st.tag_id 
    	WHERE s.title like ?
		GROUP BY s.id, s.title, s.code, s.language, s.description, s.created_at, s.updated_at`

	rows, err := db.Query(query, "%"+title+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Code, &snippet.Language, &snippet.Description, &snippet.CreatedAt, &snippet.UpdatedAt, &snippet.Tags)
		if err != nil {
			return nil, err
		}
	}

	return &snippet, nil
}

// GetSnippets gets all snippets
func GetSnippets(tag string) ([]*Snippet, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var snippets []*Snippet

	tagFilter := ``

	if tag != "" {
		tagFilter = ` WHERE t.name LIKE ?`
	}

	query := fmt.Sprintf(`WITH inner AS (
			SELECT s.* FROM snippets as s
				LEFT JOIN snippet_tags as st on st.snippet_id = s.id
				LEFT JOIN tags t on t.id = st.tag_id
				%s
				GROUP BY s.id, s.title, s.code, s.language, s.description, s.created_at, s.updated_at		
		)
			SELECT i.*, IFNULL(GROUP_CONCAT(t.name, ','), '') Tags FROM inner as i
				LEFT JOIN snippet_tags as st on st.snippet_id = i.id
				LEFT JOIN tags t on t.id = st.tag_id
		GROUP BY i.id, i.title, i.code, i.language, i.description, i.created_at, i.updated_at
		ORDER BY Created_at DESC;`, tagFilter)

	rows, err := db.Query(query, "%"+tag+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var snippet Snippet
		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Code, &snippet.Language, &snippet.Description, &snippet.CreatedAt, &snippet.UpdatedAt, &snippet.Tags)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, &snippet)
	}

	return snippets, nil
}

// SearchByTitle gets snippets by searching title
func SearchByTitle(title string) ([]*Snippet, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	var snippets []*Snippet

	query := `SELECT * FROM snippet_search WHERE title LIKE ? ;`

	rows, err := db.Query(query, "%"+title+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var snippet Snippet
		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Code, &snippet.Language, &snippet.Description, &snippet.CreatedAt, &snippet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, &snippet)
	}

	return snippets, nil
}

// SearchByLanguage gets snippets by searching language
func SearchByLanguage(language string) ([]*Snippet, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var snippets []*Snippet

	query := `SELECT * FROM snippets WHERE language LIKE ? ORDER BY created_at DESC`

	rows, err := db.Query(query, "%"+language+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var snippet Snippet
		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Code, &snippet.Language, &snippet.Description, &snippet.CreatedAt, &snippet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, &snippet)
	}

	defer db.Close()
	return snippets, nil
}

// SearchByTag gets snippets by searching tag
func SearchByTag(tag string) ([]*Snippet, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	var snippets []*Snippet

	query := `SELECT * FROM snippet_tags as st
    LEFT JOIN snippets s on s.id = st.snippet_id
    LEFT JOIN tags t on t.id = st.tag_id
	WHERE t.name like ? ORDER BY s.created_at DESC`

	rows, err := db.Query(query, "%"+tag+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var snippet Snippet
		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Code, &snippet.Language, &snippet.Description, &snippet.CreatedAt, &snippet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, &snippet)
	}

	defer db.Close()
	return snippets, nil
}
