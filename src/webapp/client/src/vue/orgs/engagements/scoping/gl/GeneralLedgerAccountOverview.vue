<template>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { RawGLAccount } from '@client/ts/types/gl'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'

@Component({
    components: {
    }
})
export default class GeneralLedgerAccountOverview extends Vue {
    get currentGeneralLedgerAccount() : RawGLAccount | null {
        return this.$store.state.gl.rawGLAccount
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }
}

</script>
