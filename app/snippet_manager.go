package app

import "code-fox/snippet"

type SnippetManager struct {
}

func NewSnippetManager() *SnippetManager {
	return &SnippetManager{}
}

func (m *SnippetManager) GetSnippetById(id int64) (*snippet.Snippet, error) {
	model, err := snippet.GetById(id)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (m *SnippetManager) GetAllSnippets(tag string) ([]*snippet.Snippet, error) {
	list, err := snippet.GetSnippets(tag)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (m *SnippetManager) CreateSnippet(model snippet.Snippet) (*snippet.Snippet, error) {
	err := model.Create()
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *SnippetManager) UpdateSnippet(model snippet.Snippet) (*snippet.Snippet, error) {
	err := model.Update()
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *SnippetManager) DeleteSnippet(model snippet.Snippet) (*snippet.Snippet, error) {
	err := model.Delete()
	if err != nil {
		return nil, err
	}

	return &model, nil
}
