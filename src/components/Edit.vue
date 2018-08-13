<template lang="pug">
  .container
    Loader(:loading="loading")
      .form
        .form-group
            input.form-control(type="text" v-model="form.title")
        .form-group
            textarea.form-control(v-model="form.content" rows="10")
        .form-inline
          .form-group
            b-dropdown.m-2(variant="primary" text="Выбери цвет")
              b-dropdown-item-button Orange
              b-dropdown-item-button Blue
              b-dropdown-item-button Green
          .form-group
            input.form-control(placeholder="тэги")
          .form-group {{form.created_at}}
        .btn-group
          .btn.btn-success(@click="") Обновить
</template>

<script>
import Loader from '@/components/Utils/Loader'
export default {
  name: 'Edit',
  data: () => {
    return {
      loading: false,
      form: {
        title: '',
        content: '',
        created_at: '',
        tags: []
      }
    }
  },
  async mounted () {
    this.loading = true
    let id = this.$route.params.id
    if (id.match(/\d+/)) {
      let result = await this.$api.send('get', `/note/${id}`)
      if (result) {
        console.log(result)
        let note = result.data.note
        this.form.title = note.title
        this.form.content = note.content
        this.form.create_at = note.create_at
      }
    } else {
      this.$router.push('/notes')
    }
    this.loading = false
  },
  components: {
    Loader
  }
}
</script>
