package book

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) GetBook(w http.ResponseWriter, r *http.Request) {
	book := c.Repository.GetBookByTitle()
	data, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
