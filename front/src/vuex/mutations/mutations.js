export default {
  SET_GAMES_TO_STATE: (state, games) => {
    state.games = games
  },
  SET_NICKNAME: (state, login) => {
    state.userNickname = login
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
}
