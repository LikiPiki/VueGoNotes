<template lang="pug">
  .form
    .form-group
        input.form-control(type="text" v-model="form.title" placeholder="Заголовок")
    .form-group
        textarea.form-control(v-model="form.content" rows="10" placeholder="Ну что, пиши заметку раз захотел...")
    .btn-group
      .btn.btn-success(@click="save") Сохранить
      .btn.btn-danger(@click="clear") Очистить
</template>

<script>
export default {
  name: 'create',
  data: () => {
    return {
      form: {
        title: '',
        content: '',
        tags: []
      }
    }
  },
  methods: {
    async save () {
      let self = this
      let sending = {
        ...self.form,
        user: this.$store.getters.getUserId
      }
      console.log('sending', sending)
      let result = await this.$api.saveNote(sending)
      console.log(result)
    },
    clear () {
      this.form.title = ''
      this.form.content = ''
      this.form.tags = []
    }
  }
}
</script>

<style scoped>

</style>
