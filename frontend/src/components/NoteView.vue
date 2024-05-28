<template>
    <subpage-template>
        <small class="text-muted">Created on {{ selected_note.Created_date }}, </small>
        <small class="text-muted">last modified on {{ selected_note.Last_modified_date}}</small>
        <br>
        <h5 class="mb-1 d-inline"  id="title_p">title_to_replace</h5>
        <form class="float-right d-inline" style="float: right;">
            <button type="submit" class="btn" title="Delete note" @click.stop.prevent="deleteNote()">
                <i class="text-danger float-right">
                    <font-awesome-icon icon="trash-alt" size="lg" class="text-danger float-right"/>
                </i>
            </button>
        </form>
        <form class="float-right d-inline" style="float: right;">
            <button type="submit" class="btn"title="Edit note" @click.stop.prevent="editNote()">
                <i class="text-danger float-right">
                    <font-awesome-icon icon="pencil-alt" size="lg" class="text-danger float-right"/>
                </i>
            </button>
        </form>
        <form class="float-right d-inline" style="float: right;">
            <button type="submit" class="btn" @click.stop.prevent="addNewSubnote()">
                <i class="text-danger float-right">
                    <font-awesome-icon icon="plus" size="lg" />
                </i>
            </button>
        </form>
        <p class="mb-1" style="white-space: pre-line" id="content_p">content_to_replace</p>
    </subpage-template>
</template>

<script>
export default {
    data() {
            return {
                selected_note: {
                    Id: 0,
                    Title: '',
                    Content: '',
                    Created_date: '',
                    Last_modified_date: '',
                    Subnote_exist: false,
                    Note_level: 0,
                    Parent_note_id: 0
                }
            }  
        },
        methods: {
            fetchSelectedNote(note_id) {
                fetch('http://localhost:8000/get_note?' + 
                    new URLSearchParams({
                        note_id: note_id
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
            deleteNote() {
                fetch('http://localhost:8000/delete_note?' + 
                    new URLSearchParams({
                        note_id: this.$route.query.note_id
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

                    if (this.selected_note.Parent_note_id === 0) {
                        this.$router.push({ path: '/'});
                    } else {
                        this.$router.push({ path: 'note_view',
                            query: { note_id: this.selected_note.Parent_note_id }
                        });
                    }
                })
            },
            editNote() {
                this.$router.push({ path: 'note_modify', query: { note_id: this.$route.query.note_id } });
            },
            addNewSubnote() {
                let parent_note_id = this.selected_note.Id
                console.log(parent_note_id)
                this.$router.push({ path: 'note_add_new', query: { Parent_note_id: parent_note_id } });
            }
        },
        beforeMount() {
            this.fetchSelectedNote(this.$route.query.note_id)
        }
}
</script>

<style>
</style>