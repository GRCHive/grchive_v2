import Vue from 'vue'
import Vuetify, { 
    VApp,
    VMain,
    VContainer,
    VRow,
    VCol
} from 'vuetify/lib'
Vue.use(Vuetify)

const LandingPage = () => import( /* webpackChunkName: "LandingPage" */ '@client/vue/LandingPage.vue')

import LandingPageAppBar from '@client/vue/components/LandingPageAppBar.vue'
import LandingFooter from '@client/vue/components/LandingFooter.vue'

import '@client/sass/main.scss'

new Vue({
    el: '#app',
    components: {
        VApp,
        VMain,
        VContainer,
        VRow,
        VCol,
        LandingPage,
        LandingPageAppBar,
        LandingFooter,
    },
    vuetify: new Vuetify({})
}).$mount('#app')
