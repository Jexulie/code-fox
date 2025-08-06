package tag

import "code-fox/database"

type Tag struct {
	Id   int64
	Name string
}

type SnippetTag struct {
	SnippetId int64
	TagId     int64
}

// NewTag returns tag struct
func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}

// Create saves tag
func (t *Tag) Create() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	query := `INSERT INTO tags (name) VALUES (?)`

	res, err := db.Exec(query, t.Name)
	if err != nil {
		return err
	}

	t.Id, _ = res.LastInsertId()

	defer db.Close()
	return nil
}

// Update updates tag
func (t *Tag) Update() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	query := `UPDATE tags SET name=? WHERE id=?`

	_, err = db.Exec(query, t.Name, t.Id)
	if err != nil {
		return err
	}

	defer db.Close()
	return nil
}

// Delete deletes tag
func (t *Tag) Delete() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	query := `DELETE FROM tags WHERE id = ?`

	_, err = db.Exec(query, t.Id)
	if err != nil {
		return err
	}

	defer db.Close()

	return nil
}

func DeleteAllTagRelations(snippetId int64) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	query := `DELETE FROM snippet_tags WHERE snippet_id = ?`

	_, err = db.Exec(query, snippetId)
	if err != nil {
		return err
	}

	return nil
}

// GetAllTags gets all tags
func GetAllTags() ([]*Tag, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := `SELECT id, name FROM tags;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []*Tag

	for rows.Next() {
		var tag Tag
		err = rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}

// GetTagByName gets tag by name
func GetTagByName(name string) (*Tag, error) {
	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := `SELECT id, name FROM tags WHERE name=?;`
	row := db.QueryRow(query, name)

	var tag Tag
	err = row.Scan(&tag.Id, &tag.Name)
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

// AddSnippetTag inserts snippet-tag relation
func AddSnippetTag(snippetId int64, tagId int64) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	defer db.Close()

	query := `INSERT INTO snippet_tags (snippet_id, tag_id) VALUES (?, ?);`
	_, err = db.Exec(query, snippetId, tagId)
	if err != nil {
		return err
	}

	return nil
}
