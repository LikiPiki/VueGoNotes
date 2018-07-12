<template lang="pug">
  .container
    Navbar
    .row.mt-3
      .offset-md-2.offset-sm-0
      .col-md-8.col-sm-12
        transition(enter-active-class="animated bouncelnDown")
          router-view
      .offset-md-2.offset-sm-0
</template>

<script>
import Navbar from './components/Navbar'

export default {
  name: 'App',
  data: () => {
    return {}
  },
  async created () {
    let token = JSON.parse(localStorage.getItem('token'))
    let userId = JSON.parse(localStorage.getItem('user_id'))
    console.log('starting app', token, userId)

    if (token && userId) {
      this.$store.dispatch('setToken', token)
      this.$store.dispatch('setUserId', userId)
    }

    if (token) {
      this.$router.push('/notes')
    }
  },
  components: {
    Navbar
  }
}
</script>
