import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import FontAwesomeIcon from './components/fontawesome'

import 'normalize.css'

const app = createApp(App)
app.use(store).use(router)
app.component('font-awesome-icon', FontAwesomeIcon)
app.mount('#app')
