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

const ScopingDashboard = () => import( /* webpackChunkName: "ScopingDashboard" */ '@client/vue/orgs/engagements/scoping/ScopingDashboard.vue')

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
const GeneralLedgerAccountSubaccounts = () => import( /* webpackChunkName: "GeneralLedgerAccountSubaccounts" */ '@client/vue/orgs/engagements/scoping/gl/GeneralLedgerAccountSubaccounts.vue')
const GeneralLedgerAccountComments = () => import( /* webpackChunkName: "GeneralLedgerAccountComments" */ '@client/vue/orgs/engagements/scoping/gl/GeneralLedgerAccountComments.vue')

const ScopingVendors = () => import( /* webpackChunkName: "ScopingVendors" */ '@client/vue/orgs/engagements/scoping/ScopingVendors.vue')
const VendorPage = () => import( /* webpackChunkName: "VendorPage" */ '@client/vue/orgs/engagements/scoping/vendors/VendorPage.vue')
const VendorOverview = () => import( /* webpackChunkName: "VendorOverview" */ '@client/vue/orgs/engagements/scoping/vendors/VendorOverview.vue')
const VendorProducts = () => import( /* webpackChunkName: "VendorProducts" */ '@client/vue/orgs/engagements/scoping/vendors/VendorProducts.vue')
const VendorComments = () => import( /* webpackChunkName: "VendorComments" */ '@client/vue/orgs/engagements/scoping/vendors/VendorComments.vue')
const VendorProductOverview = () => import( /* webpackChunkName: "VendorProductOverview" */ '@client/vue/orgs/engagements/scoping/vendors/VendorProductOverview.vue')

const ScopingInventory = () => import( /* webpackChunkName: "ScopingInventory" */ '@client/vue/orgs/engagements/scoping/ScopingInventory.vue')
const ScopingServerList = () => import( /* webpackChunkName: "ScopingServerList" */ '@client/vue/orgs/engagements/scoping//inventory/ScopingServerList.vue')
const ScopingDesktopList = () => import( /* webpackChunkName: "ScopingDesktopList" */ '@client/vue/orgs/engagements/scoping//inventory/ScopingDesktopList.vue')
const ScopingLaptopList = () => import( /* webpackChunkName: "ScopingLaptopList" */ '@client/vue/orgs/engagements/scoping//inventory/ScopingLaptopList.vue')
const ScopingMobileList = () => import( /* webpackChunkName: "ScopingMobileList" */ '@client/vue/orgs/engagements/scoping//inventory/ScopingMobileList.vue')
const ScopingEmbeddedList = () => import( /* webpackChunkName: "ScopingEmbeddedList" */ '@client/vue/orgs/engagements/scoping//inventory/ScopingEmbeddedList.vue')

const ServerPage = () => import( /* webpackChunkName: "ServerPage" */ '@client/vue/orgs/engagements/scoping/inventory/servers/ServerPage.vue')
const ServerOverview = () => import( /* webpackChunkName: "ServerOverview" */ '@client/vue/orgs/engagements/scoping/inventory/servers/ServerOverview.vue')
const ServerComments = () => import( /* webpackChunkName: "ServerComments" */ '@client/vue/orgs/engagements/scoping/inventory/servers/ServerComments.vue')

const DesktopPage = () => import( /* webpackChunkName: "DesktopPage" */ '@client/vue/orgs/engagements/scoping/inventory/desktops/DesktopPage.vue')
const DesktopOverview = () => import( /* webpackChunkName: "DesktopOverview" */ '@client/vue/orgs/engagements/scoping/inventory/desktops/DesktopOverview.vue')
const DesktopComments = () => import( /* webpackChunkName: "DesktopComments" */ '@client/vue/orgs/engagements/scoping/inventory/desktops/DesktopComments.vue')

const LaptopPage = () => import( /* webpackChunkName: "LaptopPage" */ '@client/vue/orgs/engagements/scoping/inventory/laptops/LaptopPage.vue')
const LaptopOverview = () => import( /* webpackChunkName: "LaptopOverview" */ '@client/vue/orgs/engagements/scoping/inventory/laptops/LaptopOverview.vue')
const LaptopComments = () => import( /* webpackChunkName: "LaptopComments" */ '@client/vue/orgs/engagements/scoping/inventory/laptops/LaptopComments.vue')

const MobilePage = () => import( /* webpackChunkName: "MobilePage" */ '@client/vue/orgs/engagements/scoping/inventory/mobile/MobilePage.vue')
const MobileOverview = () => import( /* webpackChunkName: "MobileOverview" */ '@client/vue/orgs/engagements/scoping/inventory/mobile/MobileOverview.vue')
const MobileComments = () => import( /* webpackChunkName: "MobileComments" */ '@client/vue/orgs/engagements/scoping/inventory/mobile/MobileComments.vue')

const EmbeddedPage = () => import( /* webpackChunkName: "EmbeddedPage" */ '@client/vue/orgs/engagements/scoping/inventory/embedded/EmbeddedPage.vue')
const EmbeddedOverview = () => import( /* webpackChunkName: "EmbeddedOverview" */ '@client/vue/orgs/engagements/scoping/inventory/embedded/EmbeddedOverview.vue')
const EmbeddedComments = () => import( /* webpackChunkName: "EmbeddedComments" */ '@client/vue/orgs/engagements/scoping/inventory/embedded/EmbeddedComments.vue')

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
        { name: 'scopingHome', path: '/orgs/:orgId/engagements/:engId/scoping', redirect: { name : 'scopingDashboard' }},
        { name: 'scopingDashboard', path: '/orgs/:orgId/engagements/:engId/scoping/dashboard', component: ScopingDashboard },
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
                    name: 'glAccSubaccounts',
                    path: 'subaccounts',
                    component: GeneralLedgerAccountSubaccounts
                },
                {
                    name: 'glAccComments',
                    path: 'comments',
                    component: GeneralLedgerAccountComments
                }
            ],
        },
        { name: 'scopingVendors', path: '/orgs/:orgId/engagements/:engId/scoping/vendors', component: ScopingVendors },
        { 
            path: '/orgs/:orgId/engagements/:engId/scoping/vendors/:vendorId',
            component: VendorPage,
            children: [
                {
                    name: 'vendorHome',
                    path: '',
                    redirect: { name: 'vendorOverview' }
                },
                {
                    name: 'vendorOverview',
                    path: 'overview',
                    component: VendorOverview,
                },
                {
                    name: 'vendorProducts',
                    path: 'products',
                    component: VendorProducts,
                    children: [
                        {
                            name: 'vendorProductOverview',
                            path: ':productId',
                            component: VendorProductOverview
                        },
                    ]
                },
                {
                    name: 'vendorComments',
                    path: 'comments',
                    component: VendorComments,
                },
            ]
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/inventory',
            component: ScopingInventory,
            children : [
                {
                    name: 'scopingInventory',
                    path: '',
                },
                {
                    name: 'scopingServers',
                    path: 'servers',
                    component: ScopingServerList,
                },
                {
                    name: 'scopingDesktops',
                    path: 'desktops',
                    component: ScopingDesktopList,
                },
                {
                    name: 'scopingLaptops',
                    path: 'laptops',
                    component: ScopingLaptopList,
                },
                {
                    name: 'scopingMobile',
                    path: 'mobile',
                    component: ScopingMobileList,
                },
                {
                    name: 'scopingEmbedded',
                    path: 'embedded',
                    component: ScopingEmbeddedList,
                },
            ]
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/inventory/servers/:serverId',
            component: ServerPage,
            children : [
                {
                    name: 'serverHome',
                    path: '',
                    redirect : { name: 'serverOverview' },
                },
                {
                    name: 'serverOverview',
                    path: 'overview',
                    component: ServerOverview,
                },
                {
                    name: 'serverComments',
                    path: 'comments',
                    component: ServerComments,
                },
            ],
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/inventory/desktops/:desktopId',
            component: DesktopPage,
            children : [
                {
                    name: 'desktopHome',
                    path: '',
                    redirect : { name: 'desktopOverview' },
                },
                {
                    name: 'desktopOverview',
                    path: 'overview',
                    component: DesktopOverview,
                },
                {
                    name: 'desktopComments',
                    path: 'comments',
                    component: DesktopComments,
                },
            ],
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/inventory/laptops/:laptopId',
            component: LaptopPage,
            children : [
                {
                    name: 'laptopHome',
                    path: '',
                    redirect : { name: 'laptopOverview' },
                },
                {
                    name: 'laptopOverview',
                    path: 'overview',
                    component: LaptopOverview,
                },
                {
                    name: 'laptopComments',
                    path: 'comments',
                    component: LaptopComments,
                },
            ],
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/inventory/mobile/:mobileId',
            component: MobilePage,
            children : [
                {
                    name: 'mobileHome',
                    path: '',
                    redirect : { name: 'mobileOverview' },
                },
                {
                    name: 'mobileOverview',
                    path: 'overview',
                    component: MobileOverview,
                },
                {
                    name: 'mobileComments',
                    path: 'comments',
                    component: MobileComments,
                },
            ],
        },
        {
            path: '/orgs/:orgId/engagements/:engId/scoping/inventory/embedded/:embeddedId',
            component: EmbeddedPage,
            children : [
                {
                    name: 'embeddedHome',
                    path: '',
                    redirect : { name: 'embeddedOverview' },
                },
                {
                    name: 'embeddedOverview',
                    path: 'overview',
                    component: EmbeddedOverview,
                },
                {
                    name: 'embeddedComments',
                    path: 'comments',
                    component: EmbeddedComments,
                },
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
