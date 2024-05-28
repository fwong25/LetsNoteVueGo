<template>
    <note-edit-template @confirm-action="confirmUpdate()" @cancel-action="cancelUpdate()" />
</template>

<script>
export default {
    data() {
            return {
                note_to_edit: {
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
                        this.note_to_edit = data

                        const title_p = document.getElementById("title_p");
                        const content_p = document.getElementById("content_p");
                        title_p.innerHTML = this.note_to_edit.Title;
                        content_p.innerHTML = this.note_to_edit.Content;
                    })
                })
            },
            confirmUpdate() {
                const title_p = document.getElementById("title_p");
                const content_p = document.getElementById("content_p");
                this.note_title = title_p.innerHTML;
                this.note_content = content_p.innerHTML;

                fetch('http://localhost:8000/update_note?' + 
                    new URLSearchParams({
                        tbl_id: this.$route.query.tbl_id,
                        note_id: this.$route.query.note_id,
                    }), 
                    {
                    method: 'PUT',
                    mode: 'cors',
                    headers: {
                        "Content-Type": "application/json",
                        "Accept": "application/json",
                        "Origin": "http://localhost:5173"
                    },
                    body: JSON.stringify({
                        title: this.note_title,
                        content: this.note_content
                    })
                })
                .then((response) => {
                    console.log(response)
                    this.$router.push({ path: 'note_view',
                        query: { tbl_id: this.$route.query.tbl_id, note_id: this.$route.query.note_id }
                    });
                })
            },
            cancelUpdate() {
                this.$router.push({ path: 'note_view',
                    query: { tbl_id: this.$route.query.tbl_id, note_id: this.$route.query.note_id }
                });
            }
        },
        beforeMount() {
            this.fetchSelectedNote(this.$route.query.note_id, this.$route.query.tbl_id)
        }
}
</script>

<style>
</style>