<template lang="pug">
  .container
    h1 Register
    .form-group
      input.form-control(type="text" placeholder="username" v-model="username")
    .form-group
      input.form-control(type="password" placeholder="password" v-model="password")
    .form-group
      input.form-control(type="text" placeholder="Reg code" v-model="regcode")
    .btn-group.float-right
      router-link.btn.btn-success(to="/" tag="div") Login
      .btn.btn-primary(@click="register") Register
</template>

<script>
export default {
  name: 'register',
  data: () => {
    return {
      username: '',
      password: '',
      regcode: ''
    }
  },
  methods: {
    async register () {
      if (this.regcode === 'whitecat') {
        await this.$api.send('post', '/register', {
          username: this.username,
          password: this.password
        })
      }
    }
  },
  mounted () {
    this.$v.$reset()
  }
}
</script>
