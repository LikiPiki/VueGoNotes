// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

import Api from './plugins/api'

import store from './store/index'

import BootstrapVue from 'bootstrap-vue'
import Vuelidate from 'vuelidate'

import Icon from 'vue-awesome/components/Icon'
import 'vue-awesome/icons/sync'
import 'vue-awesome/icons/search'
import 'vue-awesome/icons/user'
import 'vue-awesome/icons/plus'
import 'vue-awesome/icons/sign-out-alt'
import 'vue-awesome/icons/trash-alt'

// custom bootstrap theme color
import './assets/scss/theme.scss'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import 'animate.css/animate.min.css'

Vue.config.productionTip = false

// @TODO create utils mixin for date
// and state mixin for TOKEN, state, bla, bla bla....

Vue.use(BootstrapVue)
Vue.use(Vuelidate)

Vue.use(Api, {store})

Vue.component('icon', Icon)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
