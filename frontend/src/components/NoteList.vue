<template>
    <div class="container">
        <div class="row mt-5">
            <div class="col-md-8 offset-md-2">
                <div class="card">
                    <div class="card-header shadow-sm bg-white">
                        <h1 class="display-5 header-font"><font-awesome-icon icon="check-double" /> LetsNote</h1>
                    </div>
                    <div class="card-body">
                        <ul class="list-group">
                            <li class="list-group-item">
                                <form class="float-left d-inline">
                                    <div class="input-group-append header-font" style="column-width: 60px;">
                                        <span class="input-group-text bg-white py-0">
                                            <button type="submit" class="btn btn-sm header-font" @click.stop.prevent="addNewNote()">
                                                <font-awesome-icon icon="plus-circle" size="lg"/>
                                            </button>
                                        </span>
                                    </div>
                                </form>
                            </li>
                            <li class="list-group-item">
                                <h6 class="mb-1" v-for="note in note_list">
                                    <form class="d-inline">
                                        <span v-for="val in note.Note_level">&nbsp;&nbsp;</span>
                                        <button type="submit" class="astext"  @click.stop.prevent="viewNote(note.Id)">
                                            <p v-if="note.Note_level > 0">&#x2022; {{ note.Title }}</p>
                                            <h5 v-else>{{ note.Title }}</h5>
                                        </button>
                                </form>
                                </h6>
                            </li>
                        </ul>
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
            note_list: []
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
            this.$router.push({ path: 'note_add_new', query: { Parent_note_id: 0 } });
        },
        viewNote(note_id) {
            this.$router.push({ path: 'note_view', query: { note_id: note_id } });
        }
    },
    beforeMount() {
        this.fetchNoteList()
    }
}
</script>

<style>
</style>