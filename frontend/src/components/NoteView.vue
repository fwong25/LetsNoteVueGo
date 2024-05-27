<template>
    <div class="container-fluid">
        <div class="row flex-nowrap">
            <div class="col-auto col-md-3 col-lg-3 col-xl-3 px-sm-2 px-0 bg-dark">
                <div class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
                    <form>
                    <!-- <form action="/list_note" method="get" autocomplete="false" class="float-right"> -->
                        <button type="submit" class="astext mb-4 text-primary"  @click.stop.prevent="backToNoteList()">
                            <span style="font-size:16.0pt" class="mt-2 mb-2 fs-5 d-none d-sm-inline">LetsNote Home</span>
                        </button>
                    </form>
                    <div class="input-group mb-2">
                        <form>
                        <!-- <form action="/add_new_note" method="get" autocomplete="false"> -->
                            <input type="hidden" name="Parent_tbl_id" value="none">
                            <input type="hidden" name="Parent_note_id" value="none">
                            <div class="input-group-append text-info">
                                <span class="input-group-text bg-dark py-0">
                                    <button type="submit" class="btn btn-sm text-info" @click.stop.prevent="addNewNote()">
                                        <font-awesome-icon icon="plus-circle" size="lg"/>
                                    </button>
                                </span>
                            </div>
                        </form>
                    </div>
                    <ul class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start" id="menu">
                        <li class="nav-item" v-for="note in note_list">
                            <form>
                            <!-- <form action="/view_note" method="get" class="d-inline"> -->
                                <input type="hidden" name="tbl_id" v-bind:value="note.Parent_table_id.concat('_', note.Parent_note_id)">
                                <input type="hidden" name="note_id" v-bind:value="note.Id">
                                <nobr>
                                    <i v-for="val in note.Note_level">-</i>
                                    <button type="submit" class="astext"  @click.stop.prevent="viewNote(note.Parent_table_id, note.Parent_note_id, note.Id)">
                                        <i class="fs-4 bi-house"></i> <span class="ms-1 d-none d-sm-inline text-white">{{ note.Title }}</span>
                                    </button>
                                </nobr>
                            </form>
                        </li>
                    </ul>
                </div>
            </div>
            <div class="col py-3">
                <div class="container">
                    <div class="row mt-1">
                        <div class="col-md-10 offset-md-1">
                            <div class="card">
                                <div class="card-header shadow-sm bg-white">
                                    <h1 class="display-5 text-info"><font-awesome-icon icon="check-double" /> LetsNote</h1>
                                </div>
                                <div class="card-body">
                                    <small class="text-muted">Created on {{ selected_note.Created_date }}, </small>
                                    <small class="text-muted">last modified on {{ selected_note.Last_modified_date}}</small>
                                    <br>
                                    <h5 class="mb-1 d-inline"  id="title_p">title_to_replace</h5>
                                    <!-- <form action="/delete_note_item" method="delete" class="float-right d-inline"> -->
                                    <form class="float-right d-inline" style="float: right;">
                                        <input type="hidden" name="tbl_id" v-bind:value="selected_note.Parent_table_id.concat('_', selected_note.Parent_note_id)">
                                        <input type="hidden" name="note_id" v-bind:value="selected_note.Id">
                                        <button type="submit" class="btn" title="Delete note" @click.stop.prevent="deleteNote()">
                                            <!-- <i class="far fa-trash-alt fa-lg text-danger float-right"></i> -->
                                            <i class="text-danger float-right">
                                                <font-awesome-icon icon="trash-alt" size="lg" class="text-danger float-right"/>
                                            </i>
                                        </button>
                                    </form>
                                    <!-- <form action="/modify_note_item" method="get" class="float-right d-inline"> -->
                                    <form class="float-right d-inline" style="float: right;">
                                        <input type="hidden" name="tbl_id" v-bind:value="selected_note.Parent_table_id.concat('_', selected_note.Parent_note_id)">
                                        <input type="hidden" name="note_id" v-bind:value="selected_note.Id">
                                        <button type="submit" class="btn"title="Edit note" @click.stop.prevent="editNote()">
                                            <!-- <i class="fa fa-pencil-alt fa-lg text-danger float-right"></i> -->
                                            <i class="text-danger float-right">
                                                <font-awesome-icon icon="pencil-alt" size="lg" class="text-danger float-right"/>
                                            </i>
                                        </button>
                                    </form>
                                    <!-- <form action="/add_new_note" method="get" class="float-right d-inline"> -->
                                    <form class="float-right d-inline" style="float: right;">
                                        <input type="hidden" name="Parent_tbl_id" v-bind:value="selected_note.Parent_table_id.concat('_', selected_note.Parent_note_id)">
                                        <input type="hidden" name="Parent_note_id"  v-bind:value="selected_note.Id">
                                        <button type="submit" class="btn" @click.stop.prevent="addNewSubnote()">
                                            <!-- <i class="fa fa-plus fa-lg text-danger float-right"></i> -->
                                            <i class="text-danger float-right">
                                                <font-awesome-icon icon="plus" size="lg" />
                                            </i>
                                        </button>
                                    </form>
                                    <input type="hidden" id="title_input" class="form-control" name="title" v-bind:value="selected_note.Title">
                                    <textarea style="display:none;" class="form-control" id="content_input" name="content" placeholder="Content" rows="14">{{ selected_note.Content}}</textarea>
                                    <p class="mb-1" style="white-space: pre-line" id="content_p">content_to_replace</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    data() {
            return {
                note_list: [],
                note_title: '',
                note_content: '',
                selected_note: {
                    Id: 0,
                    Title: '',
                    Content: '',
                    Created_date: '',
                    Last_modified_date: '',
                    Subnote_exist: false,
                    Note_level: 0,
                    Parent_table_id: '',
                    Parent_note_id: ''
                }
            }  
        },
        methods: {
            fetchNoteList() {
                fetch('http://localhost:8000/list_note', {
                    method: 'GET',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    }
                })
                .then((response) => {
                    console.log(response)
                    response.json().then((data) => {
                        this.note_list = data
                        console.log(data)
                    })
                })
            },
            fetchSelectedNote(note_id, tbl_id) {
                fetch('http://localhost:8000/get_note?' + 
                    new URLSearchParams({
                        note_id: note_id,
                        tbl_id: tbl_id,
                    }), 
                    {
                    method: 'GET',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    }
                })
                .then((response) => {
                    console.log(response)
                    response.json().then((data) => {
                        console.log(data)
                        console.log(data.Title)
                        console.log(data.Content)
                        this.selected_note = data

                        const title_p = document.getElementById("title_p");
                        const content_p = document.getElementById("content_p");
                        title_p.innerHTML = title_p.innerHTML.replace(/title_to_replace/g, this.selected_note.Title);
                        content_p.innerHTML = content_p.innerHTML.replace(/content_to_replace/g,  this.selected_note.Content);
                    })
                })
            },
            addNewNote() {
                // with query, resulting in /register?plan=private
                // router.push({ path: 'register', query: { plan: 'private' } })
                this.$router.push({ path: 'note_add_new', query: { Parent_tbl_id: 'none', Parent_note_id: 'none' } });
                // this.$router.push({ path: 'food', query: { Parent_tbl_id: 'none', Parent_note_id: 'none' } });
            },
            backToNoteList() {
                this.$router.push({ path: '/'});
            },
            viewNote(parent_tbl_id, parent_note_id, note_id) {
                let tbl_id = parent_tbl_id.concat('_', parent_note_id)
                this.$router.push({ path: 'note_view', query: { tbl_id: tbl_id, note_id: note_id } });
            },
            deleteNote() {
                fetch('http://localhost:8000/delete_note?' + 
                    new URLSearchParams({
                        note_id: this.$route.query.note_id,
                        tbl_id: this.$route.query.tbl_id,
                    }), 
                    {
                    method: 'DELETE',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    }
                })
                .then((response) => {
                    console.log(response)
                    // response.json().then((data) => {
                    //     console.log(data)
                    // })

                    if (this.selected_note.Parent_table_id === 'none') {
                        this.$router.push({ path: '/'});
                    } else {
                        // Todo:
                        // redirect to note_view.html
                        // fetchSelectedNote(this.selected_note.Parent_table_id, this.selected_note.Parent_note_id)
                        this.$router.push({ path: 'note_view',
                            query: { tbl_id: this.selected_note.Parent_table_id, note_id: this.selected_note.Parent_note_id }
                        });
                        // this.$router.go()
                    }
                })
            },
            editNote() {
                this.$router.push({ path: 'note_modify', query: { tbl_id: this.$route.query.tbl_id, note_id: this.$route.query.note_id } });
            },
            addNewSubnote() {
                console.log(this.selected_note.Parent_table_id)
                let parent_tbl_id = this.selected_note.Parent_table_id.concat('_', this.selected_note.Parent_note_id)
                let parent_note_id = this.selected_note.Id
                console.log(parent_tbl_id)
                console.log(parent_note_id)
                this.$router.push({ path: 'note_add_new', query: { Parent_tbl_id: parent_tbl_id, Parent_note_id: parent_note_id } });
            }
        },
        beforeMount() {
            this.fetchNoteList()
            this.fetchSelectedNote(this.$route.query.note_id, this.$route.query.tbl_id)
        }
}
</script>

<style>
.astext {
    background:none;
    border:none;
    margin:0;
    padding:0;
    cursor: pointer;
}
</style>