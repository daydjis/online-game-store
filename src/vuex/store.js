import axios from 'axios'

import { createStore } from 'vuex'

const store = createStore({
  state: {
    games: [],
  },
  mutations: {
    SET_GAMES_TO_STATE: (state, games) => {
      state.games = games
    },
  },
  actions: {
    GET_GAMES_FROM_API({ commit }) {
      return axios('http://localhost:5000/api/games', {
        method: 'GET',
      })
        .then((games) => {
          commit('SET_GAMES_TO_STATE', games.data)
          return games
        })
        .catch((e) => {
          console.log(e)
          return e
        })
    },
  },
  getters: {
    GAMES(state) {
      return state.games
    },
  },
})

export default store
