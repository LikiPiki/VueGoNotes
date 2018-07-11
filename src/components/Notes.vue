<template lang="pug">
  .container
    .loader(v-if="loading")
      icon.green(name="sync" spin scale="2")
    .content(v-else)
      .form-group
        input.form-control(type="text" v-model="search" placeholder="Поищем?//!?...")
      .row(v-if="filteredNotes.length > 0")
        .col-sm-12
          b-card.mt-2(v-for="(note, index) in filteredNotes", :title="note.title")
            p.card-text {{note.content}}
      .row(v-else)
        .content.mx-auto
          h4.mx-auto Ничего не найдено :/
        .wrapper
          .loader
            icon.green(name="search" scale="2")
</template>

<script>
export default {
  name: 'notes',
  data: () => {
    return {
      loading: true,
      search: '',
      notes: []
    }
  },
  computed: {
    filteredNotes () {
      let self = this
      return this.notes.filter(el => {
        return el.title.toLowerCase().includes(self.search.toLowerCase()) ||
               el.content.toLowerCase().includes(self.search.toLowerCase())
      })
    }
  },
  async mounted () {
    this.loading = true
    console.log('getter is ', this.$store.getters.getUserId)
    let id = await this.$store.getters.getUserId
    let result = await this.$api.getNotes(id)
    if ('data' in result) {
      this.notes = result.data.notes
    }
    console.log('res is', result)
    this.loading = false
  }
}
</script>

<style scoped>
  .loader {
    width: 70px;
    height: 70px;
    background-color: #f8f9fa;
    border-radius: 50%;
    display: block;
    position: relative;
    margin: 10px auto;
  }
  .logo {
    display: block;
    width: 100%;
    margin: 10px auto;
    text-align: center;
  }
  .loader .fa-icon {
    display: block;
    position: absolute;
    top: 19px;
    left: 19px;
  }
  .wrapper {
    width: 100%;
  }
  .green {
    color: #41b883;
  }
</style>
