import "bootstrap/dist/css/bootstrap.css"
/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'
import { faCheckDouble, faPlusCircle, faTrashAlt, faPencilAlt, faPlus, faCheck } from "@fortawesome/free-solid-svg-icons";
import { faPlusSquare } from "@fortawesome/free-regular-svg-icons";
library.add(faCheckDouble, faPlusCircle, faTrashAlt, faPencilAlt, faPlus, faCheck, faPlusSquare)

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import App from './App.vue'
import NoteList from './components/NoteList.vue'
import NoteAddNew from './components/NoteAddNew.vue'
import NoteView from './components/NoteView.vue'
import NoteModify from './components/NoteModify.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '', component: NoteList },
        { path: '/note_add_new', component: NoteAddNew },
        { path: '/note_view', component: NoteView },
        { path: '/note_modify', component: NoteModify },
    ]
});

const app = createApp(App)
app.use(router)

// app.component('note-list', NoteList)
app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')

import "bootstrap/dist/js/bootstrap.js"