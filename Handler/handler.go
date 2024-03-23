package handler


// import packages  
import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// item struct  
type Item struct {

	ID	int 	`json:"id"`
	Name	string 	`json:"name"`
	Price	int 	`json:"price"`
}

var items []Item 



// handler to get all items  
func GetAllItems(w http.ResponseWriter, r *http.Request) { 
	// set content header  
	w.Header().Set("Content-Type","application/json")

	// return items in json format 
	json.NewEncoder(w).Encode(items)
}



// handler to get a single item by ID  
func GetItem(w http.ResponseWriter, r *http.Request) {
	// set content header
	w.Header().Set("Content-Type","application/json")

	// get item id with go-chi urlParam 
	idStr := chi.URLParam(r, "id")
	id , err := strconv.Atoi(idStr)
	// check for error 
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 
	}

	// if there is no error, iterate items to find specific item 
	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return 
		}
	}

	// if item is not found 
	http.NotFound(w, r)
}


// handler to create new item 
func CreateItem(w http.ResponseWriter, r *http.Request) {
	// set content header 
	w.Header().Set("Content-Type","application/json")

	var newItem Item 

	// create new item 
	if err := json.NewDecoder(r.Body).Decode(&newItem);  err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return 
	}

	// add newItem to items  
	items = append(items, newItem)
	json.NewEncoder(w).Encode(newItem)
}



// handler to update item 
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// set content header 
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 
	}

	var updatedItem Item 

	// get and process request data  
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return 
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	http.NotFound(w, r)

}



// handler to delete item 
func DeleteItem(w http.ResponseWriter, r  *http.Request){
	// set header 
	w.Header().Set("Content-Type", "application/json")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
