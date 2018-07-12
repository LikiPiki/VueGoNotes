import axios from 'axios'

const Api = {}
const api = '/api'

Api.install = (Vue, {store}) => {
  Vue.prototype.$api = {

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
    },

    send: async (type, path, data) => {
      let token = store.getters.getToken
      let headers = {
        'Authorization': 'bearer ' + token
      }

      let result = {}
      try {
        if (data) {
          result = await axios[type](api + path, data, headers)
        } else {
          result = await axios[type](api + path, headers)
        }

        if (result.data.status !== 'success') {
          result = {'error': true}
        } else {
          console.log('OKKKK!')
          result = result.data
        }
      } catch (err) {
        result = {'error': true}
      }
      return result
    }

  }
}

export default Api
