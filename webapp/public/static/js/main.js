Vue.component('note', {
    props: ['note'],
    data() {
        return {

        }
    },
    template: `
    <div class="note">
        <div v-html="note.text" @input="$emit('update-note', [note.id, $event.target.outerHTML])"></div>   

        <div class="actions">
            <span class="red" @click="$emit('delete-note', note.id)">Удалить</span>
        </div>
    </div>`
})

let app = new Vue({
    el: "#app",
    data: {
        notes: null
    },
    created(){
        this.getAllNotes();
    },
    methods: {
        getAllNotes(){
            this.notes = [];
            axios.get("/api/notes/").then((response) => {
                for (let i in response.data){
                    this.notes.push(response.data[i]);
                }
            });
        },
        createNote(){
            $.ajax({
                url: "/api/notes/",
                method: "post",
                data: JSON.stringify({text: `<div contenteditable="true">Новая заметка</div>`}),
                success: this.getAllNotes()
            });
        },
        updateNote(note){
            axios.post(`/api/notes/${note[0]}`, {"text": note[1]});
        },
        deleteNote(id){
            $.ajax({
                url: `/api/notes/${id}`,
                method: "delete",
                success: this.getAllNotes()
            });
        },
        deleteAllNotes(){
            for (let i in this.notes){
                var notes_id = this.notes[i]["id"];
                $.ajax({
                    url: `/api/notes/${notes_id}`,
                    method: "delete"
                });
            }
            this.getAllNotes();
        }
    }
});