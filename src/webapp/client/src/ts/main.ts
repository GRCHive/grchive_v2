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

import 'ag-grid-community/dist/styles/ag-grid.css'
import 'ag-grid-community/dist/styles/ag-theme-alpine.css'

const UserHome = () => import( /* webpackChunkName: "UserHome" */ '@client/vue/user/UserHome.vue')
const UserProfile = () => import( /* webpackChunkName: "UserProfile" */ '@client/vue/user/UserProfile.vue')
const OrgProfile = () => import( /* webpackChunkName: "OrgProfile" */ '@client/vue/orgs/OrgProfile.vue')
const OrgOverview = () => import( /* webpackChunkName: "OrgOverview" */ '@client/vue/orgs/profile/OrgOverview.vue')
const OrgTree = () => import( /* webpackChunkName: "OrgTree" */ '@client/vue/orgs/profile/OrgTree.vue')
const OrgEngagementList = () => import( /* webpackChunkName: "OrgEngagementList" */ '@client/vue/orgs/engagements/OrgEngagementList.vue')

const store = new Vuex.Store(RootStoreOptions)
import { ApiClient } from '@client/ts/api/client'
export const GrchiveApi = new ApiClient(store)

const router = new VueRouter({
    mode: 'history',
    base: '/app/',
    routes: [
        { name: 'appHome', path: '/', redirect: '/user' },
        { name: 'userHome', path: '/user', redirect : '/user/orgs' },
        { name: 'userOrgs', path: '/user/orgs', component: UserHome },
        { name: 'userProfile', path: '/user/profile', component: UserProfile },
        {
            path: '/orgs/:orgId',
            redirect:'/orgs/:orgId/profile'
        },
        {
            path: '/orgs/:orgId/profile',
            component: OrgProfile,
            children: [
                {
                    name: 'orgHome',
                    path: '',
                    redirect: 'overview',
                },
                {
                    path: 'overview',
                    component: OrgOverview,
                },
                {
                    path: 'tree',
                    component: OrgTree,
                },
            ],
        },
        { name: 'orgEngagements', path: '/orgs/:orgId/engagements', component: OrgEngagementList },
        { name: 'orgSingleEngagement', path: '/orgs/:orgId/engagements/:engId', component: OrgEngagementList },
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
