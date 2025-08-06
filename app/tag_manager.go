package app

import (
	"code-fox/tag"
)

type TagManager struct {
}

func NewTagManager() *TagManager {
	return &TagManager{}
}

func (m *TagManager) GetAllTags() ([]*tag.Tag, error) {
	list, err := tag.GetAllTags()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (m *TagManager) CreateTag(model tag.Tag) (*tag.Tag, error) {
	err := model.Create()
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *TagManager) UpdateTag(model tag.Tag) (*tag.Tag, error) {
	err := model.Update()
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *TagManager) DeleteTag(model tag.Tag) (*tag.Tag, error) {
	err := model.Delete()
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *TagManager) AddSnippetTagRelation(snippetId, tagId int64) error {
	err := tag.AddSnippetTag(snippetId, tagId)
	if err != nil {
		return err
	}

	return nil
}

func (m *TagManager) DeleteAllTagRelation(snippetId int64) error {
	err := tag.DeleteAllTagRelations(snippetId)
	if err != nil {
		return err
	}

	return nil
}
