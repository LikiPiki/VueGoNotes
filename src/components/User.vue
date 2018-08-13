<template lang="pug">
  b-container
    b-card(no-body)
      b-tabs(card)
        b-tab(title="Info")
          h3 Tags list
          p
            b-list-group
              b-list-group-item.d-flex.justify-content-between.align-items-center(
                v-for="(tag, index) in tab1.tags"
                :key="index"
              )
                | {{tag.tag_name}}
                h4(aria-hidden="true" @click="deleteTag(tag.id)") &times;
          .input-group-append
            b-form-input(v-model="tagText" placeholder="Add new tag").mr-1
            b-btn.btn-outline-secondary(@click="addTag")
              icon(name="plus")

        b-tab(title="Settings")
          h3 Update password
          b-form
            b-form-group
              b-form-input(placehlder="New password" type="password")
            b-form-group
              b-form-input(placeholder="Repeat new password" type="password")
</template>

<script>
export default {
  name: 'User',
  data: () => {
    return {
      tagText: '',
      tab1: {
        tags: ['lol', 'kek', 'cheburek']
      }
    }
  },
  methods: {
    async deleteTag (id) {
      let result = await this.$api.send('delete', `/tags/${id}`)
      if (result) {
        console.log('result')
      }
    },
    async addTag () {
      let result = await this.$api.send('post', '/tags', {
        user_id: this.$store.getters.userId,
        tag_name: this.tagText,
        color: 'default'
      })
      console.log(result)
      if (result) {
        this.tagText = ''
      }
    }
  },
  async mounted () {
    let tags = await this.$api.send('get', '/tags')
    if (tags) {
      this.tab1.tags = tags.data.tags
    }
    console.log(tags)
  }
}
</script>
