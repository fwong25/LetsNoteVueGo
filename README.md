# LetsNote App

### App Features
- Create notes with title and content
- Add subnotes
- Notes are editable
- Records created date and last modified date

### Prerequisites
- Backend: Golang, Postgresql
- Frontend: Vue.js, bootstrap

Frontend UI design reference: [Todo App with Python+Django](https://www.youtube.com/watch?v=Nnoxz9JGdLU&ab_channel=CodAffection)

Before running project, make sure you have created the DB table in postgresql with following cmd
```
CREATE TABLE note_table_0 (
    id SERIAL PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_date TEXT NOT NULL,
    last_modified_date TEXT NOT NULL,
    subnote_exist BOOLEAN NOT NULL DEFAULT FALSE,
    note_level INT NOT NULL DEFAULT 0,
    parent_note_id INT NOT NULL DEFAULT 0
);
```

Modify the DB connection settings in src/db_mgmt/db_mgmt.go

<img src="assets/db_setting.png" alt="drawing" width="200"/>

### Running the project
```
cd backend
go run main.go

cd frontend
npm run dev
```

In your browser, e.g. Chrome, open the following address:
http://127.0.0.1:8000/list_note
or
http://localhost:8000/list_note

### UI Introduction

#### Main page
<img src="assets/main_page.png" alt="drawing" width="500"/>

#### Add note page
<img src="assets/add_note_page.png" alt="drawing" width="500"/>

#### View note page
<img src="assets/view_note_page.png" alt="drawing" width="500"/>

#### Edit note page
<img src="assets/edit_note_page.png" alt="drawing" width="500"/>

#### Sidebar introduction
<img src="assets/sidebar_intro.png" alt="drawing" width="500"/>

### Database Design
<img src="assets/db_design.png" alt="drawing" width="500"/>

### Function Implementations
#### Add note
<img src="assets/add_note_algo.png" alt="drawing" width="380"/>

#### Delete note
<img src="assets/delete_note_algo.png" alt="drawing" width="380"/>