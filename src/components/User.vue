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
              )
                | {{tag}}
                b-badge(variant="primary" pill) 2
          .input-group-append
            b-form-input(placeholder="Add new tag").mr-1
            b-btn.btn-outline-secondary
              icon(name="plus")

        b-tab(title="Settings")
          h3 Update password
          b-for
            b-form-group
              b-form-input(placeholder="New password" type="password")
            b-form-group
              b-form-input(placeholder="Repeat new password" type="password")
</template>

<script>
export default {
  name: 'User',
  data: () => {
    return {
      tab1: {
        tags: ['lol', 'kek', 'cheburek']
      }
    }
  }
}
</script>