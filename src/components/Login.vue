<template lang="pug">
  .container
    h1 Login
    b-form-group
      b-form-input(
        type="text"
        placeholder="username"
        v-model.lazy="$v.username.$model"
        :state="!$v.username.$invalid"
      )
      b-form-invalid-feedback(v-show="!$v.username.required") Required username
      b-form-invalid-feedback(v-show="!$v.username.minLength") Minimum 4 symbols
    b-form-group
      b-form-input(
        type="password"
        placeholder="password"
        v-model.lazy="$v.password.$model"
        :state="!$v.password.$invalid"
      )
      b-form-invalid-feedback(v-show="!$v.password.required") Required password
      b-form-invalid-feedback(v-show="!$v.password.minLength") Minimum 5 symbols
    .btn-group.float-right
      router-link.btn.btn-primary(to="/register" tag="div") Register
      .btn.btn-success(@click="login") Login
</template>

<script>
import { validationMixin } from 'vuelidate'
import { required, minLength } from 'vuelidate/lib/validators'

export default {
  name: 'App',
  data: () => {
    return {
      username: '',
      password: '',
      token: ''
    }
  },
  mixins: [validationMixin],
  validations: {
    username: {
      required,
      minLength: minLength(5)
    },
    password: {
      required,
      minLength: minLength(5)
    }
  },
  methods: {
    async login () {
      this.$v.$touch()
      if (!this.$v.$invalid) {
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
  },
  mounted () {
    // not stable variant
    // todo rewrite
    this.$v.$reset()
  }
}
</script>
