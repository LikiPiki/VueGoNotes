<template lang="pug">
  .container
    br
    .row
      .col-sm-3
        pre {{token}}
      .col-sm-6
        h1 Login
        .form-group
          input.form-control(v-model="username")
        .form-group
          input.form-control(v-model="password")
        .form-group
          .btn.btn-success(@click="login") Enter
          .btn.btn-primary(@click="show") Show
      .col-sm-3 {{dd}}
</template>

<script>
import axios from 'axios'

export default {
  name: 'App',
  data: () => {
    return {
      username: '',
      password: '',
      token: '',
      dd: ''
    }
  },
  methods: {
    async login () {
      let result = await axios.post('/api/login', {
        username: this.username,
        password: this.password
      })
      if (result.data.token) {
        this.token = result.data.token
      }
    },
    async show () {
      let result = await axios.get('/api/private', {
        headers: {
          'Authorization': 'bearer ' + this.token
        }
      })
      console.log(result.data)
      this.dd = result.data
    }
  }
}
</script>
