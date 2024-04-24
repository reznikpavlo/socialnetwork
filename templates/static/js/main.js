import Vue from 'vue'
import Vuetify from 'vuetify'
import App from 'pages/App.vue'
import 'vuetify/dist/vuetify.min.css'
//import axios from "axios";

//Vue.use(axios)
Vue.use(Vuetify)

new Vue({
    el: '#app',
    render: a => a(App)
})




