import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import VueAxios from 'vue-axios'
Vue.config.productionTip = false
Vue.use(VueAxios)
new Vue({
  vuetify,
  render: h => h(App)
}).$mount('#app')
