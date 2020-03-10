package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/theantichris/todo-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AddTodoItemHandler adds a new TodoItem.
func AddTodoItemHandler(db *mgo.Collection) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		_ = db.Insert(models.TodoItem{
			ID:          bson.NewObjectId(),
			Date:        time.Now(),
			Description: r.FormValue("description"),
			Done:        false,
		})

		result := models.TodoItem{}
		_ = db.Find(bson.M{"description": r.FormValue("description")}).One(&result)
		json.NewEncoder(w).Encode(result)
	}

	return http.HandlerFunc(fn)
}
