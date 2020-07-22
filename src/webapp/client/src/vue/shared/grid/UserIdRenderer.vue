<template>
    <span>
        {{ userStr }}
    </span>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { RawUser } from '@client/ts/types/users'
import { GrchiveApi } from '@client/ts/main'

@Component
export default class UserIdRenderer extends Vue {
    params : any = null

    relevantUser : RawUser | null = null

    get userStr() : string {
        if (!this.relevantUser) {
            return 'None'
        }
        return `${this.relevantUser.FullName} (${this.relevantUser.Email})`
    }

    syncUser() {
        if (!this.params || !this.params.value) {
            this.relevantUser = null
            return
        }

        const orgId : number = Number(this.$route.params.orgId)
        if (isNaN(orgId)) {
            this.relevantUser = null
            return
        }
        GrchiveApi.orgs.getOrgUser(orgId, this.params.value).then((resp : RawUser | null) => {
            this.relevantUser = resp
        })
    }

    refresh(params:  any) {
        this.params = params
        this.syncUser()
        return true
    }

    mounted() {
        this.syncUser()
    }
}

</script>
