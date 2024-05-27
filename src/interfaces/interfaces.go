package interfaces_v2

import (
	"db_mgmt"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	w.WriteHeader(http.StatusOK)

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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	w.WriteHeader(http.StatusOK)

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

func insertNoteAction(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	if action != "Insert" {
		parent_tbl_id := r.FormValue("Parent_tbl_id")
		parent_note_id := r.FormValue("Parent_note_id")

		if parent_tbl_id == "none" {
			http.Redirect(w, r, "/list_note", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/view_note?tbl_id="+parent_tbl_id+"&note_id="+parent_note_id, http.StatusSeeOther)
		}
	} else {
		insertNote(w, r)
	}
}

type insert_note_body_t struct {
	Title   string
	Content string
}

func insertNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	w.WriteHeader(http.StatusOK)

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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	w.WriteHeader(http.StatusOK)

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

func modifyNote(w http.ResponseWriter, r *http.Request) {
	note_id := r.FormValue("note_id")
	table_id := r.FormValue("tbl_id")
	note_list := db_mgmt.GetNoteListAllSubtable(db_mgmt.Main_tbl_id)

	if table_id == "none_none" {
		table_id = db_mgmt.Main_tbl_id
	}
	note_to_edit := db_mgmt.GetNote(table_id, note_id)

	var fileName = "note_modify.html"
	t, err := template.New(fileName).Funcs(funcMap).ParseFiles(filepath.Join(templatesDirPath, fileName))
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}

	note_id_int, _ := strconv.Atoi(note_id)

	modify_note_info := map[string]interface{}{"Note_id": note_id_int, "Note_list": note_list, "Note_to_edit": note_to_edit}

	err = t.ExecuteTemplate(w, fileName, modify_note_info)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")

	w.WriteHeader(http.StatusOK)

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

func modifyNoteConfirm(w http.ResponseWriter, r *http.Request) {
	note_id := r.FormValue("note_id")
	tbl_id := r.FormValue("tbl_id")
	fmt.Print("Confirm modification on note with ID:", note_id)
	fmt.Println(", Table ID:", tbl_id)
	title := r.FormValue("title")
	content := r.FormValue("content")

	if tbl_id == "none_none" {
		tbl_id = db_mgmt.Main_tbl_id
	}
	db_mgmt.UpdateNote(tbl_id, note_id, title, content)

	http.Redirect(w, r, "/view_note?tbl_id="+tbl_id+"&note_id="+note_id, http.StatusSeeOther)
}

func modifyNoteAction(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("action")
	note_id := r.FormValue("note_id")
	tbl_id := r.FormValue("tbl_id")
	if action != "Update" {
		http.Redirect(w, r, "/view_note?tbl_id="+tbl_id+"&note_id="+note_id, http.StatusSeeOther)
	} else {
		modifyNoteConfirm(w, r)
	}
}
