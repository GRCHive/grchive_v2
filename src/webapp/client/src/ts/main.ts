import 'core-js/stable'
import 'regenerator-runtime/runtime'

import Vue from 'vue'
import Vuetify, { 
    VApp,
    VMain,
    VContainer,
    VRow,
    VCol
} from 'vuetify/lib'
import VueRouter from 'vue-router'
import Vuex from 'vuex'
Vue.use(Vuex)
Vue.use(Vuetify)
Vue.use(VueRouter)

import { RootStoreOptions } from '@client/ts/stores/store'
import '@client/sass/main.scss'

const UserHome = () => import( /* webpackChunkName: "UserHome" */ '@client/vue/user/UserHome.vue')
const UserProfile = () => import( /* webpackChunkName: "UserProfile" */ '@client/vue/user/UserProfile.vue')
const OrgProfile = () => import( /* webpackChunkName: "OrgProfile" */ '@client/vue/orgs/OrgProfile.vue')

const store = new Vuex.Store(RootStoreOptions)
import { ApiClient } from '@client/ts/api/client'
export const GrchiveApi = new ApiClient(store)

const router = new VueRouter({
    mode: 'history',
    base: '/app/',
    routes: [
        { name: 'appHome', path: '/', redirect: '/user' },
        { name: 'userHome', path: '/user', component: UserHome },
        { name: 'userProfile', path: '/user/profile', component: UserProfile },
        { name: 'orgHome', path: '/orgs/:orgId', component: OrgProfile },
    ],
})

new Vue({
    router,
    store,
    el: '#app',
    components: {
        VApp,
        VMain,
        VContainer,
        VRow,
        VCol,
    },
    vuetify: new Vuetify({})
}).$mount('#app')
