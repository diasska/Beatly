import './assets/main.css'
import 'v-calendar/dist/style.css';

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import VCalendar from 'v-calendar';
import axios from 'axios'; 

axios.defaults.baseURL = 'http://localhost:8080';

const app = createApp(App)

app.use(router)
app.use(VCalendar, {})

app.mount('#app')
