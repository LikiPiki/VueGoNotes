<template lang="pug">
  .container
    .loader(v-if="loading")
      icon.green(name="sync" spin scale="2")
    .content(v-else)
      .form-group
        input.form-control(type="text" v-model="search" placeholder="Поищем?//!?...")
      .row(v-if="filteredNotes.length > 0")
        .col-sm-12
          b-card.mt-2(
            v-for="(note, index) in filteredNotes"
          )
            span(slot="header")
              p {{note.title}}
                button.close(
                  aria-label="Удалить"
                  @click="deleteNote(note.id)"
                )
                  span(aria-hidden="true") &times;
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
  methods: {
    async deleteNote (id) {
      let result = await this.$api.send('delete', '/notes/' + id)
      if (result.data.status === 'success') {
        this.update()
      }
    },
    async update () {
      let id = await this.$store.getters.getUserId
      let result = await this.$api.send('get', '/notes/' + id)
      if (result) {
        this.notes = result.data.notes
      } else {
        this.$router.push('/')
      }
    }
  },
  async mounted () {
    this.loading = true
    await this.update()
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
