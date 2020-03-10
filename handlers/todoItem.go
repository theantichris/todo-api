package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

		// TODO: this is always returning the first TodoItem that has that description.
		result := models.TodoItem{}
		_ = db.Find(bson.M{"description": r.FormValue("description")}).One(&result)
		json.NewEncoder(w).Encode(result)
	}

	return http.HandlerFunc(fn)
}

// GetTodoItemHandler gets a TodoItem.
func GetTodoItemHandler(db *mgo.Collection) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		var results []models.TodoItem

		vars := mux.Vars(r)
		id := vars["id"]

		if id != "" {
			results = GetByID(id, db)
		} else {
			_ = db.Find(nil).All(&results)
		}

		json.NewEncoder(w).Encode(results)
	}

	return http.HandlerFunc(fn)
}

// GetByID gets a TodoItem by ID.
func GetByID(id string, db *mgo.Collection) []models.TodoItem {
	var item models.TodoItem
	var results []models.TodoItem

	_ = db.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&item)

	results = append(results, item)

	return results
}
