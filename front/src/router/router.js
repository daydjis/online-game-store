import { createRouter, createWebHashHistory } from 'vue-router'
import gMainWrapper from '../components/g-main-wrapper.vue'
import gCart from '../components/g-cart.vue'
import gAddNewGame from '../components/g-add-new-game.vue'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: '/', component: gMainWrapper, alias: '/' },
    { path: '/cart', component: gCart },
    { path: '/create/game', component: gAddNewGame },
  ],
})
