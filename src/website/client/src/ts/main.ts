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
const GettingStartedPage = () => import( /* webpackChunkName: "GettingStartedPage" */ '@client/vue/GettingStartedPage.vue')
const ContactUsPage = () => import( /* webpackChunkName: "ContactUsPage" */ '@client/vue/ContactUsPage.vue')
const LearnMorePage = () => import( /* webpackChunkName: "LearnMorePage" */ '@client/vue/LearnMorePage.vue')

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
        LandingPageAppBar,
        LandingFooter,
        LandingPage,
        GettingStartedPage,
        ContactUsPage,
        LearnMorePage,
    },
    vuetify: new Vuetify({})
}).$mount('#app')
