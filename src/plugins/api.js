import axios from 'axios'

const Api = {}
const api = '/api'

Api.install = (Vue, {store}) => {
  Vue.prototype.$api = {

    login (username, password) {
      return axios.post(api + '/login', {
        username: username,
        password: password
      })
    },

    getNotes (id) {
      return axios.get(api + '/getNotes/' + id, {
        headers: {
          'Authorization': 'bearer ' + store.getters.getToken
        }
      })
    },

    saveNote (note) {
      return axios.post(api + '/createNote', note, {
        headers: {
          'Authorization': 'bearer ' + store.getters.getToken
        }
      })
    }

  }
}

export default Api
