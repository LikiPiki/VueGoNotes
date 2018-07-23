<template lang="pug">
  .container
    .loader(v-if="loading")
      icon.green(name="sync" spin scale="2")
    slot(v-else="!loading")
</template>

<script>
export default {
  name: 'Loader',
  data: () => {
    return {}
  },
  props: {
    loading: {
      type: Boolean,
      default: false,
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
