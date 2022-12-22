import axios from 'axios'
import router from '../../router/router'
import { parseJwt } from '../../utils/parserJwt'
import { getCookie } from '../../utils/getCookie'

export default {
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
      const ourCookie = getCookie('jwt')
      await axios
        .post('http://localhost:5000/api/games/new', this.state.newGameForm, {
          method: 'POST',
          headers: {
            Authorization: `${ourCookie}`,
          },
        })
        .then(function (response) {
          console.log('игра успешно добавлена', response)
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
          const newCokkie = parseJwt(response.data.jwt)
          localStorage.setItem('login', newCokkie.login)
          commit('SET_NICKNAME', newCokkie.login)
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
          const newCokkie = parseJwt(response.data.jwt)
          commit('SET_NICKNAME', newCokkie.login)
          localStorage.setItem('login', newCokkie.login)
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
}
