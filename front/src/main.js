import { createApp } from 'vue'
import App from './App.vue'
import router from './router/router'
import store from './vuex/store'
import VueIframe from 'vue-iframes'

const app = createApp(App)

app.use(VueIframe)
app.use(store)
app.use(router)
app.mount('#app')
