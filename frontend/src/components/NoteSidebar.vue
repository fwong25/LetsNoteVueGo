<template>
<div class="col-auto col-md-3 col-lg-3 col-xl-3 px-sm-2 px-0 bg-dark">
    <div class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
        <form class="float-right">
            <button type="submit" class="astext mb-4 text-primary"  @click.stop.prevent="backToNoteList()">
                <span style="font-size:16.0pt" class="mt-2 mb-2 fs-5 d-none d-sm-inline">LetsNote Home</span>
            </button>
        </form>
        <div class="input-group mb-2">
            <form>
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
            viewNote(parent_tbl_id, parent_note_id, note_id) {
                let tbl_id = parent_tbl_id.concat('_', parent_note_id)
                this.$router.push({ path: 'note_view', query: { tbl_id: tbl_id, note_id: note_id } });
            }
        },
        beforeMount() {
            this.fetchNoteList()
        }
}
</script>

<style>
</style>