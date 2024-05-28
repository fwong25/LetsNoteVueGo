package interfaces_v2

import (
	"db_mgmt"
	"encoding/json"
	"fmt"
	"net/http"
)

const templatesDirPath = "templates"

var funcMap = map[string]interface{}{
	"Iterate": func(count int) []uint {
		var i uint
		var Items []uint
		for i = 0; i < uint(count); i++ {
			Items = append(Items, i)
		}
		return Items
	},
	"IsSubnote": func(note_level int) bool {
		if note_level == 0 {
			return false
		}
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case ("/list_note"):
		listNote(w, r)
	case ("/insert_note"):
		insertNote(w, r)
	case ("/get_note"):
		getNote(w, r)
	case ("/update_note"):
		updateNote(w, r)
	case ("/delete_note"):
		deleteNote(w, r)
	default:
		fmt.Fprintf(w, "bad address 404")
	}
}

func listNote(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "GET")

	if r.Method == "GET" {
		note_list := db_mgmt.GetNoteListAllSubtable(db_mgmt.Main_tbl_id)
		var result, err = json.Marshal(note_list)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	note_id := r.FormValue("note_id")
	tbl_id := r.FormValue("tbl_id")

	if tbl_id == "none_none" {
		tbl_id = db_mgmt.Main_tbl_id
	}

	fmt.Println("View note with table ID: " + tbl_id + ", note ID: " + note_id)
	selected_note := db_mgmt.GetNote(tbl_id, note_id)

	setHeader(w, "GET")

	if r.Method == "GET" {
		var result, err = json.Marshal(selected_note)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

type insert_note_body_t struct {
	Title   string
	Content string
}

func insertNote(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "POST")

	contentType := r.Header.Get("Content-type")
	var title, content string
	if contentType == "application/json" {
		var t insert_note_body_t
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		title = t.Title
		content = t.Content
	} else {
		title = r.FormValue("title")
		content = r.FormValue("content")
	}

	parent_tbl_id := r.FormValue("Parent_tbl_id")
	parent_note_id := r.FormValue("Parent_note_id")

	fmt.Println("Insert note with the following parameters:")
	fmt.Println("Parent_tbl_id: ", parent_tbl_id)
	fmt.Println("Parent_note_id: ", parent_note_id)
	fmt.Println("Title: ", title)
	fmt.Println("Content: ", content)

	if r.Method == "POST" {
		if parent_tbl_id == "none_none" {
			parent_tbl_id = db_mgmt.Main_tbl_id
		}
		note_id, table_id := db_mgmt.InsertNote(parent_tbl_id, parent_note_id, title, content)
		// note_id := 5
		// table_id := 10
		fmt.Print("New record note ID is:", note_id)
		fmt.Println(", Table ID: ", table_id)

		var result, err = json.Marshal(map[string]interface{}{"note_id": note_id, "table_id": table_id})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "DELETE")

	tbl_id := r.FormValue("tbl_id")
	note_id := r.FormValue("note_id")
	fmt.Println("Delete note with table ID: ", tbl_id, "note id:", note_id)

	if tbl_id == "none_none" {
		tbl_id = db_mgmt.Main_tbl_id
	}

	if r.Method == "DELETE" {
		db_mgmt.DeleteNote(tbl_id, note_id)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	setHeader(w, "PUT")

	note_id := r.FormValue("note_id")
	tbl_id := r.FormValue("tbl_id")
	fmt.Print("Update note with ID:", note_id)
	fmt.Println(", Table ID:", tbl_id)

	contentType := r.Header.Get("Content-type")
	var title, content string
	if contentType == "application/json" {
		var t insert_note_body_t
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		title = t.Title
		content = t.Content
	} else {
		title = r.FormValue("title")
		content = r.FormValue("content")
	}

	if r.Method == "PUT" {
		if tbl_id == "none_none" {
			tbl_id = db_mgmt.Main_tbl_id
		}
		db_mgmt.UpdateNote(tbl_id, note_id, title, content)
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}

func setHeader(w http.ResponseWriter, allow_methods string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", allow_methods)
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	w.WriteHeader(http.StatusOK)
}
