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
      token: '',
      dd: ''
    }
  },
  methods: {
    async login () {
      let result = await this.$api.login(this.username, this.password)
      if ('data' in result) {
        console.log(result)
        if (result.data.success) {
          this.$store.dispatch('setToken', result.data.token)
          this.$store.dispatch('setUserId', result.data.user_id)

          localStorage.setItem('token', JSON.stringify(result.data.token))
          localStorage.setItem('user_id', JSON.stringify(result.data.user_id))

          this.$router.push('/notes')
          console.log(this.$store.state)
        }
      }
    }
  }
}
</script>
