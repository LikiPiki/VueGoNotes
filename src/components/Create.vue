<template lang="pug">
  .form
    .form-group
        input.form-control(type="text" v-model="form.title" placeholder="Заголовок")
    .form-group
        textarea.form-control(v-model="form.content" rows="10" placeholder="Ну что, пиши заметку, раз захотел...")
    .form-inline
      .form-group
        b-dropdown.m-2(variant="primary" text="Выбери цвет")
          b-dropdown-item-button Orange
          b-dropdown-item-button Blue
          b-dropdown-item-button Green
      .form-group
        input.form-control(placeholder="тэги")
    .btn-group
      .btn.btn-success(@click="save") Сохранить
      .btn.btn-danger(@click="clear") Очистить
    b-alert.mt-4(:show="alert.show" dismissible :variant="alert.type") {{alert.text}}
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
      },
      alert: {
        show: false,
        type: 'success',
        text: 'Something text',
        time: 3,
        dismissCountDown: 10
      }
    }
  },
  methods: {
    async save () {
      let self = this
      let sending = {
        ...self.form,
        user_id: this.$store.getters.getUserId,
        created_at: this.moment(Date.now())
      }
      let result = await this.$api.send('post', '/notes', sending)

      if (result.status === 200) {
        this.alert.text = 'Сохранено успешно'
        this.alert.type = 'success'
      } else {
        this.alert.text = 'Произошла неведомая ошиПка...'
        this.alert.type = 'danger'
      }
      this.alert.dismissCountDown = this.alert.time
      this.alert.show = true
    },
    clear () {
      this.form.title = ''
      this.form.content = ''
      this.form.tags = []
    }
  }
}
</script>
