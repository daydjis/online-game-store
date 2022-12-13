import axios from 'axios'
import { createStore } from 'vuex'

const store = createStore({
  state: {
    games: [],
    cart: [],
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
  },

  actions: {
    async GET_GAMES_FROM_API({ commit }) {
      try {
        const games = await axios('http://localhost:5000/api/games', {
          method: 'GET',
        })
        commit('SET_GAMES_TO_STATE', games.data)
        return games
      } catch (e) {
        console.log(e)
        return e
      }
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
  },
})

export default store
