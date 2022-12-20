import axios from 'axios'
import { createStore } from 'vuex'
import router from '../router/router'

const store = createStore({
  state: {
    // сюда кладём игры после гет запроса
    games: [],
    // корзина
    cart: [],
    // лоадер
    isLoading: false,
    // форма для игры
    newGameForm: {
      title: '',
      description: '',
      price: 0,
      genres: [],
      video: '',
      imageDescription: '',
      image: '',
    },
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
  },

  mutations: {
    SET_GAMES_TO_STATE: (state, games) => {
      state.games = games
    },
    SET_CART: (state, game) => {
      if (state.cart.length) {
        let isProductExists = false
        state.cart.map((item) => {
          if (item.id === game.id) {
            isProductExists = true
            item.quantity++
          }
        })
        if (!isProductExists) {
          state.cart.push(game)
        }
      } else {
        state.cart.push(game)
      }
    },
    REMOVE_FROM_CART: (state, gameIndex) => {
      state.cart.splice(gameIndex, 1)
    },
    ISLOADING: (state, loading) => {
      state.isLoading = loading
    },
    CREATE_NEW_GAME: (state, newGameInfo) => {
      state.newGameForm = newGameInfo
    },
    SET_CURRENT_GAME: (state, game) => {
      state.setCurrentGame = game
    },
    SET_GAME_ID: (state, gameId) => {
      state.gameId = gameId[0]
    },
    SET_USER_INFO: (state, authInfo) => {
      state.userInfo = authInfo
    },
    SET_COOKIE: (state, cookie) => {
      state.isCookie = cookie
    },
  },

  actions: {
    async GET_CURRENT_GAME({ commit }, gameid) {
      try {
        commit('ISLOADING', true)
        const game = await axios(
          'http://localhost:5000/api/games',
          { params: { id: gameid } },
          {
            method: 'GET',
          }
        )
        commit('SET_GAME_ID', game.data)
        return game
      } catch (e) {
        console.log(e)
        return e
      } finally {
        commit('ISLOADING', false)
      }
    },

    async GET_GAMES_FROM_API({ commit }) {
      try {
        commit('ISLOADING', true)
        const games = await axios('http://localhost:5000/api/games', {
          method: 'GET',
        })
        commit('SET_GAMES_TO_STATE', games.data)
        return games
      } catch (e) {
        console.log(e)
        return e
      } finally {
        commit('ISLOADING', false)
      }
    },

    async POST_NEW_GAME({ commit }, newGameInfo) {
      try {
        commit('CREATE_NEW_GAME', newGameInfo)
        const getCookie = (name) => {
          const value = `; ${document.cookie}`
          const parts = value.split(`; ${name}=`)
          if (parts.length === 2) return parts.pop().split(';').shift()
        }
        const ourCookie = getCookie('jwt')
        await axios
          .post('http://localhost:5000/api/games/new', this.state.newGameForm, {
            method: 'POST',
            headers: {
              Authorization: `${ourCookie}`,
            },
          })
          .then(function (response) {
            console.log('УРА', response)
          })
      } catch (error) {
        console.log('Ошибка пост запроса', error)
      } finally {
        console.log('Наш куки', document.cookie)
      }
    },

    async LOGIN_USER({ commit }, authInfo) {
      try {
        commit('SET_USER_INFO', authInfo)
        commit('ISLOADING', true)
        await axios
          .post('http://localhost:5000/api/login', this.state.userInfo, {
            withCredentials: true,
          })
          .then(function (response) {
            document.cookie = `jwt=${response.data.jwt}`
            commit('SET_COOKIE', true)
          })
      } catch (error) {
        console.log('Ошибка пост запроса', error)
      } finally {
        commit('ISLOADING', false)
        if (document.cookie) {
          router.push({ path: `/` })
        } else {
          console.log('попался')
        }
      }
    },

    async REGISTER_USER({ commit }, authInfo) {
      try {
        commit('SET_USER_INFO', authInfo)
        commit('ISLOADING', true)
        await axios
          .post('http://localhost:5000/api/register', this.state.userInfo, {
            withCredentials: true,
          })
          .then(function (response) {
            document.cookie = `jwt=${response.data.jwt}`
            commit('SET_COOKIE', true)
          })
      } catch (error) {
        console.log('Ошибка пост запроса', error)
        commit('SET_COOKIE', null)
      } finally {
        commit('ISLOADING', false)
        if (document.cookie) {
          router.push({ path: `/` })
        } else {
          console.log('попался')
        }
      }
    },

    SET_CURRENT_GAME: ({ commit }, game) => {
      commit('SET_CURRENT_GAME', game)
    },

    ADD_GAME_TO_CART({ commit }, game) {
      commit('SET_CART', game)
    },
    DELETE_FROM_CART({ commit }, game) {
      commit('REMOVE_FROM_CART', game)
    },
    CHECK_COOKIE({ commit }) {
      if (document.cookie) {
        commit('SET_COOKIE', true)
      } else {
        commit('SET_COOKIE', null)
      }
    },
    DELETE_COOKIE({ commit }) {
      const cookies = document.cookie.split(';')
      for (let i = 0; i < cookies.length; i++) {
        const cookie = cookies[i]
        const eqPos = cookie.indexOf('=')
        const name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie
        document.cookie = name + '=;expires=Thu, 01 Jan 1970 00:00:00 GMT;'
        document.cookie =
          name + '=; path=/; expires=Thu, 01 Jan 1970 00:00:01 GMT;'
      }
      if (document.cookie) {
        commit('SET_COOKIE', true)
      } else {
        commit('SET_COOKIE', null)
      }
    },
  },

  getters: {
    GAMES(state) {
      return state.games
    },
    CART(state) {
      return state.cart
    },
    LOADER(state) {
      return state.isLoading
    },
    NEW_GAME(state) {
      return state.newGameForm
    },
    CURRENT_GAME(state) {
      return state.setCurrentGame
    },
    GAME_ID(state) {
      return state.gameId
    },
    USER_INFO(state) {
      return state.userInfo
    },
    COOKIE_IS_EXIST(state) {
      return state.isCookie
    },
  },
})

export default store
