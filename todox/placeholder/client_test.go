package placeholder

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"todox/todo"
)

func TestClient_GetTodo(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/1", r.URL.Path)
		example := `{
		"userId": 1,
  		"id": 1,
  		"title": "delectus aut autem",
		"completed": true
		}`
		_, _ = w.Write([]byte(example))
	}))
	defer s.Close()

	pc := Client{
		httpClient: &http.Client{},
		baseURL:    s.URL + "/",
	}
	td, err := pc.GetTodo(1)
	assert.NoError(t, err)

	expected := todo.Todo{
		UserID:        1,
		ID:            1,
		Title:         "delectus aut autem",
		IsCompleted:   true,
		AnotherSecret: "",
	}
	assert.Equal(t, expected, td)
}
