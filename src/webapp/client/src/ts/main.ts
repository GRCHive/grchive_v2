import Vue from 'vue'
import Vuetify, { 
    VApp,
    VMain,
    VContainer,
    VRow,
    VCol
} from 'vuetify/lib'
Vue.use(Vuetify)

new Vue({
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
