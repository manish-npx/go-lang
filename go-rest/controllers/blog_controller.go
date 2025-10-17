package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/manish-npx/go-lang/go-rest/models"
)

var blogs = []models.Blog{
	{ID: 1, Title: "New Blog Title-1"},
	{ID: 2, Title: "New Blog Title-2"},
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
	json.NewEncoder(w).Encode(blogs)

}
