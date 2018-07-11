// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import Api from './plugins/api'
import store from './store/index'

import BootstrapVue from 'bootstrap-vue'
import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/sync'
import 'vue-awesome/icons/search'

// globally (in your main .js file)

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import 'animate.css/animate.min.css'

Vue.config.productionTip = false

Vue.use(Api, {store})

Vue.use(BootstrapVue)
Vue.component('icon', Icon)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
