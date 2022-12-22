export default {
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
  USER_NICKNAME(state) {
    return state.userNickname
  },
}
