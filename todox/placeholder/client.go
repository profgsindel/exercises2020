package placeholder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"todox/todo"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	secret     string
}

func New(secret string) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    "https://jsonplaceholder.typicode.com/todos/",
		secret:     secret,
	}
}

func (c *Client) GetTodo(id int) (todo.Todo, error) {
	// "https://jsonplaceholder.typicode.com/todos/{id}"
	resp, err := c.httpClient.Get(fmt.Sprintf("%s%d", c.baseURL, id))
	if err != nil {
		return todo.Todo{}, err
	}
	defer resp.Body.Close()
	log.Printf("%v: %v", resp.Status, resp.StatusCode)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return todo.Todo{}, err
	}

	var td todo.Todo
	err = json.Unmarshal(b, &td)
	if err != nil {
		return todo.Todo{}, err
	}
	return td, nil
}
