import axios from 'axios'
import { createStore } from 'vuex'

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
    }
  },

  actions: {
    async GET_CURRENT_GAME({ commit }, gameid) {
      try {
        commit('ISLOADING', true)
        const game = await axios('http://localhost:5000/api/games', { params: { id: gameid } }, {
          method: 'GET',
        })
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
        await axios
          .post('http://localhost:5000/api/games/new', this.state.newGameForm, {
            method: 'POST',
          })
          .then(function (response) {
            console.log('УРА', response)
          })
      } catch (error) {
        console.log('Ошибка пост запроса', error)
        console.log('NEW_GAME')
      }
    },
    async LOGIN_USER({ commit }, userInfo) {
      try {
        commit('CREATE_NEW_GAME', 'NEW_GAME')
        await axios
          .post('http://localhost:5000/api/login', userInfo)
          .then(function (response) {
            console.log(response)
          })
      } catch (error) {
        console.log('Ошибка пост запроса', error)
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
    }
  },
})

export default store
