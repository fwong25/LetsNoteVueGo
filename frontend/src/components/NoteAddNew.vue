<template>
    <div class="container-fluid">
        <div class="row flex-nowrap">
            <div class="col-auto col-md-3 col-lg-3 col-xl-3 px-sm-2 px-0 bg-dark">
                <div class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
                    <form class="float-right">
                        <button type="submit" class="astext mb-4 text-primary"  @click.stop.prevent="backToNoteList()">
                            <span style="font-size:16.0pt" class="mt-2 mb-2 fs-5 d-none d-sm-inline">LetsNote Home</span>
                        </button>
                    </form>
                    <div class="input-group mb-2">
                        <form>
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
                            <form  class="d-inline">
                                <input type="hidden" name="tbl_id" v-bind:value="note.Parent_table_id.concat('_', note.Parent_note_id)">
                                <input type="hidden" name="note_id" v-bind:value="note.Id">
                                <nobr>
                                    <i v-for="val in note.Note_level">-</i>
                                    <button type="submit" class="astext" @click.stop.prevent="viewNote(note.Parent_table_id, note.Parent_note_id, note.Id)">
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
                                    <form>
                                        <div class="input-group mb-2">
                                            <input type="hidden" id="title_input" class="form-control" name="title" placeholder="Title">
                                            <input type="text" class="form-control" placeholder="Edit mode" readonly>
                                            <div class="input-group-append text-info">
                                                <span class="input-group-text bg-white py-0">
                                                    <button type="submit" class="btn btn-sm text-info" @click.stop.prevent="insertNote()">
                                                        <h6>
                                                            <font-awesome-icon icon="plus-circle" size="lg"/>
                                                            <!-- <font-awesome-icon :icon="['far', 'plus-square']"  size="lg"/> -->
                                                        </h6>
                                                    </button>
                                                </span>
                                            </div>
                                            <div class="input-group-append text-info">
                                                <span class="input-group-text bg-white py-0">
                                                    <button type="submit" class="btn btn-sm text-info"  @click.stop.prevent="cancelAddNote()">
                                                        <h6 style="color:rgb(197, 5, 5) !important;"><b>X</b></h6>
                                                    </button>
                                                </span>
                                            </div>
                                        </div>
                                        <div class="input-group">
                                            <textarea style="display:none;" class="form-control" id="content_input" name="content" placeholder="Content" rows="14"></textarea>
                                        </div>
                                        <h5 data-placeholder="Insert note title here..."class="mb-1" id="title_p" contenteditable="true" ></h5>
                                        <p data-placeholder="Insert note content here..." class="mb-1" id="content_p" style="white-space: pre-line" contenteditable="true"></p>
                                        
                                    </form>
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
            addNewNote() {
                this.$router.push({ path: 'note_add_new', query: { Parent_tbl_id: 'none', Parent_note_id: 'none' } });
                // this.$router.push({ path: 'food', query: { Parent_tbl_id: 'none', Parent_note_id: 'none' } });
            },
            backToNoteList() {
                this.$router.push({ path: '/'});
            },
            insertNote() {
                const title_p = document.getElementById("title_p");
                const content_p = document.getElementById("content_p");
                this.note_title = title_p.innerHTML;
                this.note_content = content_p.innerHTML;

                fetch('http://localhost:8000/insert_note?' + 
                    new URLSearchParams({
                        Parent_tbl_id: this.$route.query.Parent_tbl_id,
                        Parent_note_id: this.$route.query.Parent_note_id,
                    }), 
                    {
                    method: 'POST',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    },
                    body: JSON.stringify({
                        // "title": 'test title',
                        // "content": 'test content'
                        title: this.note_title,
                        content: this.note_content
                    })
                })
                .then((response) => {
                    console.log(response)
                    response.json().then((data) => {
                        console.log(data)
                        console.log(data.note_id)
                        console.log(data.table_id)
                        // todo: redirect to note_view.html
                        this.$router.push({ path: 'note_view',
                            query: { tbl_id: data.table_id, note_id: data.note_id }
                        });
                    })
                })
            },
            cancelAddNote() {
                if (this.$route.query.Parent_tbl_id === 'none') {
                    this.$router.push({ path: '/'});
                } else {
                    // Todo:
                    // redirect to note_view.html
                    this.$router.push({ path: 'note_view',
                        query: { tbl_id: this.$route.query.Parent_tbl_id, note_id: this.$route.query.Parent_note_id }
                    });
                }
            },
            viewNote(parent_tbl_id, parent_note_id, note_id) {
                let tbl_id = parent_tbl_id.concat('_', parent_note_id)
                // this.$router.push({ path: 'food', query: { tbl_id: tbl_id, note_id: note_id } });
                // todo: redirect to note_view.html
                this.$router.push({ path: 'note_view', query: { tbl_id: tbl_id, note_id: note_id } });
            }
        },
        beforeMount() {
            this.fetchNoteList()
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
    input {background:none!important;}
    input:focus {outline:none!important;}

    p:empty:not(:focus)::before {
        content: attr(data-placeholder);
    }
    h5:empty:not(:focus)::before {
        content: attr(data-placeholder);
    }
</style>