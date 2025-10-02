package main

import "github.com/vince-II/go-snippetcode.git/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
