<template lang="pug">
  .container
    h1 Register
    .form-group
      input.form-control(type="text" placeholder="username" v-model="username")
    .form-group
      input.form-control(type="password" placeholder="password" v-model="password")
    .form-group
      input.form-control(type="text" placeholder="Reg code" v-model="regcode")
    .btn-group.center-block
      .btn.btn-success(@click="register") Register
      router-link.btn.btn-primary(to="/" tag="div") Login
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
  }
}
</script>
