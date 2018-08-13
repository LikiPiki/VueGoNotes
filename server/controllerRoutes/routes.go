package controllerRoutes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetByIDRoute create route, which get model by id
func GetByIDRoute(c GetController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		result, err := c.GetById(id)

		if err != nil {
			errorJSON(w, err.Error())
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": "success",
			"data":    result,
		})
	}
}

// PostRoute create route, which create model
func PostRoute(c PostController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			errorJSON(w, err.Error())
			return
		}

		err = c.Create()

		if err != nil {
			errorJSON(w, err.Error())
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"created": true,
		})
	}
}

// PutRoute update route, which update model
func PutRoute(c PutController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var err error
		c, err = c.Decode(r.Body)
		if err != nil {
			errorJSON(w, err.Error())
			return
		}

		err = c.Update(id)

		if err != nil {
			errorJSON(w, err.Error())
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"updated": true,
		})
	}
}

// PutRoute update route, which update model
func PutRoute(c DeleteController) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		err := c.Delete(id)
		if err != nil {
			errorJSON(w, err.Error())
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "success",
			"deleted": true,
		})
	}
}

func errorJSON(w http.ResponseWriter, err string) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "error",
		"error":  err,
	})
}
