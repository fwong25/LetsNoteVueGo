<template>
<note-edit-template @confirm-action="insertNote()" @cancel-action="cancelAddNote()" />
</template>

<script>
export default {
    data() {
            return {
                note_title: '',
                note_content: '',
            }  
        },
        methods: {
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
                    this.$router.push({ path: 'note_view',
                        query: { tbl_id: this.$route.query.Parent_tbl_id, note_id: this.$route.query.Parent_note_id }
                    });
                }
            }
        }
}
</script>

<style>
    p:empty:not(:focus)::before {
        content: attr(data-placeholder);
        color: grey;
    }
    h5:empty:not(:focus)::before {
        content: attr(data-placeholder);
        color: grey;
    }
</style>