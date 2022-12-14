import { createRouter, createWebHashHistory } from 'vue-router'
import gMainWrapper from '../components/g-main-wrapper.vue'
import gCart from '../components/g-cart.vue'
import gAddNewGame from '../components/g-add-new-game.vue'
import gAddNewInfo from '../components/g-game-info.vue'
import gAuthLogin from '../components/auth/g-auth-login.vue'
import gAuthRegister from '../components/auth/g-auth-register.vue'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', component: gMainWrapper, alias: '/' },
    { path: '/cart', component: gCart },
    { path: '/create/game', component: gAddNewGame },
    { path: '/auth/register', component: gAuthRegister },
    { path: '/auth/login', component: gAuthLogin },
    { path: '/game:title?', component: gAddNewInfo },
  ],
})
