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
const OrgEngagement = () => import( /* webpackChunkName: "OrgEngagement" */ '@client/vue/orgs/engagements/OrgEngagement.vue')

const ScopingRisks = () => import( /* webpackChunkName: "ScopingRisks" */ '@client/vue/orgs/engagements/scoping/ScopingRisks.vue')
const RiskPage = () => import( /* webpackChunkName: "RiskPage" */ '@client/vue/orgs/engagements/scoping/risks/RiskPage.vue')
const RiskOverview = () => import( /* webpackChunkName: "RiskOverview" */ '@client/vue/orgs/engagements/scoping/risks/RiskOverview.vue')
const RiskComments = () => import( /* webpackChunkName: "RiskComments" */ '@client/vue/orgs/engagements/scoping/risks/RiskComments.vue')

const ScopingControls = () => import( /* webpackChunkName: "ScopingControls" */ '@client/vue/orgs/engagements/scoping/ScopingControls.vue')
const ControlPage = () => import( /* webpackChunkName: "ControlPage" */ '@client/vue/orgs/engagements/scoping/controls/ControlPage.vue')
const ControlOverview = () => import( /* webpackChunkName: "ControlOverview" */ '@client/vue/orgs/engagements/scoping/controls/ControlOverview.vue')
const ControlComments = () => import( /* webpackChunkName: "ControlComments" */ '@client/vue/orgs/engagements/scoping/controls/ControlComments.vue')

const ScopingGeneralLedger = () => import( /* webpackChunkName: "ScopingGeneralLedger" */ '@client/vue/orgs/engagements/scoping/ScopingGeneralLedger.vue')
const GeneralLedgerAccounts = () => import( /* webpackChunkName: "GeneralLedgerAccounts" */ '@client/vue/orgs/engagements/scoping/gl/GeneralLedgerAccounts.vue')
const GeneralLedgerComments = () => import( /* webpackChunkName: "GeneralLedgerComments" */ '@client/vue/orgs/engagements/scoping/gl/GeneralLedgerComments.vue')
const GeneralLedgerAccountPage = () => import( /* webpackChunkName: "GeneralLedgerAccountPage" */ '@client/vue/orgs/engagements/scoping/gl/GeneralLedgerAccountPage.vue')
const GeneralLedgerAccountOverview = () => import( /* webpackChunkName: "GeneralLedgerAccountOverview" */ '@client/vue/orgs/engagements/scoping/gl/GeneralLedgerAccountOverview.vue')
const GeneralLedgerAccountComments = () => import( /* webpackChunkName: "GeneralLedgerAccountComments" */ '@client/vue/orgs/engagements/scoping/gl/GeneralLedgerAccountComments.vue')

const ErrorPage = () => import( /* webpackChunkName: "ErrorPage" */ '@client/vue/ErrorPage.vue')

const store = new Vuex.Store(RootStoreOptions)
import { ApiClient } from '@client/ts/api/client'
export const GrchiveApi = new ApiClient()

const router = new VueRouter({
    mode: 'history',
    base: '/app/',
    routes: [
        { name: 'appHome', path: '/', redirect: {name : 'userHome'} },
        { name: 'userHome', path: '/user', redirect : {name: 'userOrgs'} },
        { name: 'userOrgs', path: '/user/orgs', component: UserHome },
        { name: 'userProfile', path: '/user/profile', component: UserProfile },
        {
            path: '/orgs/:orgId',
            redirect: {name : 'orgHome'},
        },
        {
            path: '/orgs/:orgId/profile',
            component: OrgProfile,
            children: [
                {
                    name: 'orgHome',
                    path: '',
                    redirect: {name: 'orgOverview'},
                },
                {
                    name: 'orgOverview',
                    path: 'overview',
                    component: OrgOverview,
                },
                {
                    name: 'orgTree',
                    path: 'tree',
                    component: OrgTree,
                },
            ],
        },
        { name: 'orgEngagements', path: '/orgs/:orgId/engagements', component: OrgEngagementList },
        { name: 'orgSingleEngagement', path: '/orgs/:orgId/engagements/:engId', component: OrgEngagement },
        { name: 'scopingHome', path: '/orgs/:orgId/engagements/:engId/scoping', redirect: { name : 'scopingRisks' }},
        { name: 'scopingRisks', path: '/orgs/:orgId/engagements/:engId/scoping/risks', component: ScopingRisks },
        { 
            path: '/orgs/:orgId/engagements/:engId/scoping/risks/:riskId',
            component: RiskPage,
            children: [
                {
                    name: 'riskHome',
                    path: '',
                    redirect: { name: 'riskOverview' }
                },
                {
                    name: 'riskOverview',
                    path: 'overview',
                    component: RiskOverview,
                },
                {
                    name: 'riskComments',
                    path: 'comments',
                    component: RiskComments,
                },
            ]
        },
        { name: 'scopingControls', path: '/orgs/:orgId/engagements/:engId/scoping/controls', component: ScopingControls },
        { 
            path: '/orgs/:orgId/engagements/:engId/scoping/controls/:controlId',
            component: ControlPage,
            children: [
                {
                    name: 'controlHome',
                    path: '',
                    redirect: { name: 'controlOverview' }
                },
                {
                    name: 'controlOverview',
                    path: 'overview',
                    component: ControlOverview,
                },
                {
                    name: 'controlComments',
                    path: 'comments',
                    component: ControlComments,
                }
            ]
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/gl',
            component: ScopingGeneralLedger,
            children: [
                {
                    name: 'glHome',
                    path: '',
                    redirect: { name: 'glAccounts' }
                },
                {
                    name: 'glAccounts',
                    path: 'accs',
                    component: GeneralLedgerAccounts
                },
                {
                    name: 'glComments',
                    path: 'comments',
                    component: GeneralLedgerComments
                }
            ],
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/gl/accs/:accId',
            component: GeneralLedgerAccountPage,
            children: [
                {
                    name: 'glAccHome',
                    path: '',
                    redirect: { name: 'glAccOverview' }
                },
                {
                    name: 'glAccOverview',
                    path: 'overview',
                    component: GeneralLedgerAccountOverview
                },
                {
                    name: 'glAccComments',
                    path: 'comments',
                    component: GeneralLedgerAccountComments
                }
            ],
        },
        {
            path: '*',
            name: 'errorPage',
            component: ErrorPage
        },
    ],
})

import { GrchiveErrorHandler } from '@client/ts/error'
export const ErrorHandler = new GrchiveErrorHandler(store, router)

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
