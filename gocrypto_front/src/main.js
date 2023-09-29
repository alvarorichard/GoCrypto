import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import store from './store.js'
import r from './router.js'
createApp(App).use(r).use(store).mount('#app')
