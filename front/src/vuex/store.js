import { createStore } from 'vuex'

import mutations from './mutations/mutations'
import getters from './getters/getters'
import commonActions from './actions/actions'
import apiRequest from './actions/api-request'

const actions = { ...commonActions, ...apiRequest }
const store = createStore({
  state: {
    // сюда кладём игры после гет запроса
    games: [],
    // корзина
    cart: [],
    // лоадер
    isLoading: false,
    // форма для игры
    newGameForm: {},
    // карточка игры
    setCurrentGame: {},
    // game info
    gameId: {},
    userInfo: {
      login: '',
      password: '',
    },
    // Для проверки на наличие куки
    isCookie: null,
    // отображение ника в хеддере
    userNickname: '',
  },

  mutations,
  actions,
  getters,
})

export default store
