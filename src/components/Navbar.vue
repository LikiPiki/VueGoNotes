<template lang="pug">
  nav.navbar.navabar-expand-lg.navbar-light.bg-light.mt-4
    router-link.navbar-brand(to="/notes" tag="a") Notes
    .btn-group
      router-link.btn.btn-warning(to="/user", tag="div")
        icon(name="user")
      router-link.btn.btn-success(to="/create" tag="div")
        icon(name="plus")
      .btn.btn-primary(@click="logout" v-if="visible")
        icon(name="sign-out-alt")
</template>

<script>
export default {
  data: () => {
    return {}
  },
  computed: {
    visible () {
      if (this.$store.getters.getToken !== '') return true
      return false
    }
  },
  methods: {
    logout () {
      this.$store.dispatch('setToken', '')
      localStorage.clear('token')
      localStorage.clear('user')
      localStorage.clear('id')
      this.$router.push('/')
    }
  }
}
</script>
