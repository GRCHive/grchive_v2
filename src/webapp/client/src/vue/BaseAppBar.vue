<template>
    <v-app-bar
        app
        dense
        clipped-left
        clipped-right
        :extension-height="extensionHeight"
    >
        <v-app-bar-nav-icon @click.stop="clickNav"></v-app-bar-nav-icon>
        <v-toolbar-title color="primary" class="pl-1">
            <router-link
                :to="{ name: 'appHome'}"
            >
                <v-img
                    src="/static/assets/logos/grchive.png"
                    height="36px"
                    max-height="36px"
                    max-width="254px"
                    contain
                >
                </v-img>
            </router-link>
        </v-toolbar-title>

        <slot name="left-items"></slot>

        <v-spacer></v-spacer>

        <slot name="right-items"></slot>

        <v-toolbar-items>
            <v-btn
                text
                color="primary"
                href="mailto:support@grchive.com"
            >
                Support
                <v-icon color="primary" small>mdi-email</v-icon>
            </v-btn>

            <v-menu offset-y>
                <template v-slot:activator="{ on }">
                    <v-btn
                        text
                        color="primary"
                        v-on="on"
                    >
                        {{ fullName }}
                        <v-icon color="primary">mdi-menu-down</v-icon>
                    </v-btn>
                </template>
                <v-list dense>
                    <v-list-item dense :to="{ name: 'userProfile' }">
                        <v-list-item-title>My Account</v-list-item-title>
                    </v-list-item>
                    <v-list-item dense href="/logout">
                        <v-list-item-title>Logout</v-list-item-title>
                    </v-list-item>
                </v-list>

            </v-menu>

        </v-toolbar-items>

        <template v-slot:extension>
            <verify-email-banner></verify-email-banner>
        </template>
    </v-app-bar>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import VerifyEmailBanner from '@client/vue/verification/VerifyEmailBanner.vue'

@Component({
    components: {
        VerifyEmailBanner
    }
})
export default class BaseAppBar extends Vue {
    get fullName() : string {
        return this.$store.state.user.rawUser.FullName
    }

    get extensionHeight() : number {
        if (this.$store.getters['appLayout/showEmailVerificationBanner']) {
            return 44
        } else {
            return 0
        }
    }

    clickNav() {
        this.$store.commit('appLayout/toggleMini')
    }
}

</script>

<style scoped>

.v-menu__content {
    border-radius: 0px !important;
}

>>>.v-toolbar__extension {
    padding: 0;
    margin: 0;
}

</style>
