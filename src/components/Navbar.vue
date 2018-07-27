<template lang="pug">
  nav.navbar.navabar-expand-lg.navbar-light.bg-light.mt-4
    router-link.navbar-brand(to="/notes" tag="a") Notes
    .btn-group
      router-link.btn.btn-success(to="/create" tag="div") Create
      .btn.btn-primary(@click="logout" v-if="visible") Logout
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
