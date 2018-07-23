<template lang="pug">
  .container
    .row
      .content.mx-auto
        h4.mx-auto {{ text }}
    .wrapper
      .loader
        icon.green(:name="icon" scale="2")
</template>

<script>
export default {
  name: 'Banner',
  data: () => {
    return {
    }
  },
  props: {
    text: {
      type: String,
      default: 'Wow, must be a string',
      required: true
    },
    icon: {
      type: String,
      required: true
    }
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
