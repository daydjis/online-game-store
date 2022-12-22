export default {
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
    localStorage.clear()
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
}
