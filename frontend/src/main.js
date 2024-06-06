import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue/dist/bootstrap-vue.esm';
import Buefy from 'buefy'
import VueApexCharts from 'vue-apexcharts'
import VueRouter from 'vue-router'
import Home from './components/Home.vue'
import Commitsdata from './components/Commitsdata.vue'
import Commits1data from './components/Commits1data.vue'
import Chartsdata from './components/Chartsdata.vue'
import Charts1data from './components/Charts1data.vue'
import { BootstrapVueIcons } from 'bootstrap-vue'
import {BadgerAccordion, BadgerAccordionItem} from 'vue-badger-accordion'

import 'bootstrap-vue/dist/bootstrap-vue-icons.min.css'
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import 'buefy/dist/buefy.css'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faCoffee } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import { MdButton, MdContent, MdTabs } from 'vue-material/dist/components'
import VueMaterial from 'vue-material'
import 'vue-material/dist/vue-material.min.css'
import 'vue-material/dist/theme/default.css'

Vue.use(VueMaterial)

Vue.use(MdButton)
Vue.use(MdContent)
Vue.use(MdTabs)
Vue.use(Vuetify)

library.add(faCoffee)


Vue.use(VueRouter)
Vue.use(VueApexCharts)
Vue.use(Buefy)
Vue.use(BootstrapVue);
Vue.use(BootstrapVueIcons)
Vue.component('apexchart', VueApexCharts)
Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.component('BadgerAccordion', BadgerAccordion)
Vue.component('BadgerAccordionItem', BadgerAccordionItem)
const routes = [
  {path: '/',component:Home },
  {path: '/commits',component:Commitsdata },
  {path: '/charts',component:Chartsdata },
  {path: '/commits/:id',component:Commits1data },
  {path: '/charts/:id',component:Charts1data },
];

const router = new VueRouter({
  routes: routes,
  mode:'history'
});

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  router:router
}).$mount('#app')
