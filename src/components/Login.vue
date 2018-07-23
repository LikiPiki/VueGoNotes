<template lang="pug">
  .container
    h1 Login
    .form-group
      input.form-control(type="text" placeholder="username" v-model="username")
    .form-group
      input.form-control(type="password" placeholder="password" v-model="password")
    .btn-group.center-block
      .btn.btn-success(@click="login") Login
      router-link.btn.btn-primary(to="/register" tag="div") Register
</template>

<script>
export default {
  name: 'App',
  data: () => {
    return {
      username: '',
      password: '',
      token: ''
    }
  },
  methods: {
    async login () {
      let self = this
      let result = await this.$api.send('post', '/login', {
        username: self.username,
        password: self.password
      })

      console.log('status', result)
      if (result.status === 200 && result.data.status === 'success') {
        this.$store.dispatch('setToken', result.data.token)
        this.$store.dispatch('setUserId', result.data.id)

        localStorage.setItem('token', JSON.stringify(result.data.token))
        localStorage.setItem('user_id', JSON.stringify(result.data.id))

        this.$router.push('/notes')
        console.log(this.$store.state)
      }
    }
  }
}
</script>
