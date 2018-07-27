<template lang="pug">
  .container
    Loader(:loading="loading")
      .form-group
        input.form-control(type="text" v-model="search" placeholder="Поищем?//!?...")
      .row(v-if="filteredNotes.length > 0")
        .col-sm-12
          b-card.mt-2(
            v-for="(note, index) in filteredNotes"
            :key="index"
          )
            span(slot="header")
              p {{note.title}}
                button.close(
                  aria-label="Удалить"
                  @click="deleteNote(note.id)"
                )
                  span(aria-hidden="true") &times;
            p(slot="footer") {{formatDate(note.created_at)}}
            p.card-text {{note.content}}
            .button-group
              router-link.btn.btn-success.btn-sm(:to="'/edit/' + note.id" tag="div") Edit
      Banner(v-else text="Ничего не найдено :/ ..." icon="search")
    br
    br
</template>

<script>
import Loader from '@/components/Utils/Loader'
import Banner from '@/components/Utils/Banner'

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
    formatDate (date) {
      return this.moment(date).format('hh:mm DD:MM:YYYY')
    },
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
    try {
      await this.update()
    } catch (err) {
      console.log(err)
      this.$store.dispatch('setToken', '')
      this.$router.push('/')
    }
    this.loading = false
  },
  components: {
    Loader, Banner
  }
}
</script>
