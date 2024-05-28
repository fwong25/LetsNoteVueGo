package db_mgmt

import (
	"database/sql"
	"fmt"
	"models"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "12345678"
	dbname   = "LetsNoteDB_v2"
)

var Main_tbl_id string = "note_table"

var db *sql.DB
var err error

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error when connecting to database", err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error when pinging database", err)
	}

	fmt.Println("Successfully connected to database!")
}

func CloseDB() {
	db.Close()
}

func setSubnoteExist(note_id int, val bool) (affected_rows int64) {
	sqlStatement := `
		UPDATE ` + Main_tbl_id + `
		SET subnote_exist = $1
		WHERE id = $2;`
	res, err := db.Exec(sqlStatement, strconv.FormatBool(val), strconv.Itoa(note_id))
	if err != nil {
		fmt.Println("Error when updating note", err)
	}

	affected_rows, _ = res.RowsAffected()
	return
}

func GetNoteListAllSubtable(note_id int) []models.Note {
	place := models.Note{} // Initialize a User struct to hold retrieved data
	queryStat := "SELECT id, title, content, created_date, last_modified_date, subnote_exist, note_level, parent_note_id FROM " + Main_tbl_id + " WHERE parent_note_id = " + strconv.Itoa(note_id) + " ORDER BY last_modified_date"

	rows, _ := db.Query(queryStat)

	note_list := []models.Note{}

	for rows.Next() {
		err := rows.Scan(&place.Id, &place.Title, &place.Content, &place.Created_date, &place.Last_modified_date, &place.Subnote_exist, &place.Note_level, &place.Parent_note_id) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
		if place.Subnote_exist {
			subnote_list := GetNoteListAllSubtable(place.Id)
			note_list = append(subnote_list, note_list...)
		}
		note_list = append([]models.Note{place}, note_list...) // ... unpack into note_list[0], note_list[1], ...
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return note_list
}

func GetDirectSubnotes(note_id int) []models.Note {
	place := models.Note{} // Initialize a User struct to hold retrieved data

	rows, _ := db.Query("SELECT id, title, content, created_date, last_modified_date, subnote_exist, note_level, parent_note_id FROM " + Main_tbl_id +
		" WHERE parent_note_id = " + strconv.Itoa(note_id))
	note_list := []models.Note{}

	for rows.Next() {
		err := rows.Scan(&place.Id, &place.Title, &place.Content, &place.Created_date, &place.Last_modified_date, &place.Subnote_exist, &place.Note_level, &place.Parent_note_id) // Scan the current row into the "place" variable
		if err != nil {
			fmt.Println("Error when scanning rows", err)
		}
		// fmt.Printf("%#v\n", place) // Log the content of the "place" struct for each row
		note_list = append([]models.Note{place}, note_list...) // ... unpack into note_list[0], note_list[1], ...
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		fmt.Println("Error during note iteration", err)
	}

	return note_list
}

func GetNote(note_id int) models.Note {
	note := models.Note{} // Initialize struct to hold retrieved data

	queryStat := "SELECT id, title, content, created_date, last_modified_date, subnote_exist, note_level, parent_note_id FROM " + Main_tbl_id + " WHERE id = " + strconv.Itoa(note_id)

	rows, _ := db.Query(queryStat)

	if rows == nil {
		fmt.Println("Can't retrieve note with ID: ", note_id)
	}

	rows.Next()
	err := rows.Scan(&note.Id, &note.Title, &note.Content, &note.Created_date, &note.Last_modified_date, &note.Subnote_exist, &note.Note_level, &note.Parent_note_id) // Scan the current row into the "place" variable
	if err != nil {
		fmt.Println("Error when scanning rows", err)
	}

	return note
}

func insertNotePrecheck(parent_note_id int) (note_level int) {
	if parent_note_id == 0 {
		return 0
	}
	parent_note := GetNote(parent_note_id)
	if !parent_note.Subnote_exist {
		setSubnoteExist(parent_note_id, true)
	}
	return parent_note.Note_level + 1
}

func InsertNote(parent_note_id int, title string, content string) (note_id int) {
	currentTime := time.Now()
	created_date := currentTime.Format("2006-01-02")
	last_modified_date := currentTime.Format("2006-01-02")
	fmt.Println("Insert note:")
	fmt.Println(title, content, created_date, last_modified_date)
	fmt.Println("Parent_note_id: " + strconv.Itoa(parent_note_id))

	note_level := insertNotePrecheck(parent_note_id)

	sqlStatement := `
		INSERT INTO ` + Main_tbl_id + ` (title, content, created_date, last_modified_date, note_level, parent_note_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err = db.QueryRow(sqlStatement, title, content, created_date, last_modified_date, note_level, parent_note_id).Scan(&note_id)
	if err != nil {
		fmt.Println("Error when inserting note", err)
	}

	return
}

func DeleteNote(note_id int) (affected_rows int64) {
	note := GetNote(note_id)
	fmt.Println("Delete note with ID: ", note_id)

	if note.Subnote_exist {
		DeleteSubNotes(note_id)
	}

	sqlStatement := `
		DELETE FROM ` + Main_tbl_id + `
		WHERE id = $1;`
	res, err := db.Exec(sqlStatement, note_id)

	if err != nil {
		fmt.Println("Error when deleting note", err)
	}

	affected_rows, _ = res.RowsAffected()

	is_only_subnote := isOnlySubnoteOfParent(note.Parent_note_id)

	if is_only_subnote {
		setSubnoteExist(note.Parent_note_id, false)
	}
	return
}

func DeleteSubNotes(note_id int) {
	fmt.Println("Delete subnote of note ID: ", note_id)

	note_list := GetDirectSubnotes(note_id)

	for _, note := range note_list {
		if note.Subnote_exist {
			DeleteSubNotes(note.Id)
		}
	}
	return
}

func isOnlySubnoteOfParent(parent_note_id int) (is_only_subnote bool) {

	if parent_note_id == 0 {
		return false
	}

	rows_cnt := 0
	queryStat := "SELECT COUNT(*) FROM " + Main_tbl_id + " WHERE parent_note_id = " + strconv.Itoa(parent_note_id)

	rows, _ := db.Query(queryStat)

	if rows == nil {
		fmt.Println("Can't query COUNT of entries with parent_note_id: ", parent_note_id)
	}

	rows.Next()
	err := rows.Scan(&rows_cnt)
	if err != nil {
		fmt.Println("Error when scanning rows in isOnlySubnoteOfParent", err)
	}

	if rows_cnt == 0 {
		return true
	}
	return false
}

func UpdateNote(note_id int, title string, content string) (affected_rows int64) {
	fmt.Println("Update note with note ID:", note_id)

	currentTime := time.Now()
	last_modified_date := currentTime.Format("2006-01-02")

	sqlStatement := `
		UPDATE ` + Main_tbl_id + `
		SET title = $2, content = $3, last_modified_date = $4
		WHERE id = $1
		RETURNING title, content;`
	res, err := db.Exec(sqlStatement, note_id, title, content, last_modified_date)

	if err != nil {
		fmt.Println("Error when updating note", err)
	}

	affected_rows, _ = res.RowsAffected()
	return
}
