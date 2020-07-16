<template>
    <org-template
        page-name="Profile"
    >
        <template v-slot:content>
            <v-list-item class="px-0">
                <v-list-item-content>
                    <v-list-item-title class="text-h4">
                        Organization Profile
                    </v-list-item-title>
                </v-list-item-content>
            </v-list-item>
            <v-divider class="mb-4"></v-divider>
        </template>
    </org-template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import OrgTemplate from '@client/vue/orgs/OrgTemplate.vue'

@Component({
    components: {
        OrgTemplate
    }
})
export default class OrgProfile extends Vue {
}

</script>
