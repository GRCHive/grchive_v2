<template>
    <comment-thread-manager
        :thread-id="threadId"
    >
    </comment-thread-manager>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { RawServer } from '@client/ts/types/inventory'
import { RawOrganization } from '@client/ts/types/orgs'
import { RawEngagement } from '@client/ts/types/engagements'
import { CommentThreadId } from '@client/ts/types/comments'
import CommentThreadManager from '@client/vue/types/comments/CommentThreadManager.vue'

@Component({
    components: {
        CommentThreadManager
    }
})
export default class ServerComments extends Vue {
    get threadId() : CommentThreadId {
        return {
            orgId: this.currentOrg!.Id,
            engagementId: this.currentEngagement!.Id,
            serverId: this.currentServer!.Id,
        }
    }

    get currentServer() : RawServer | null {
        return this.$store.state.inventory.rawServer
    }

    get currentOrg() : RawOrganization | null {
        return this.$store.state.org.rawOrg
    }

    get currentEngagement() : RawEngagement | null {
        return this.$store.state.engagements.rawEngagement
    }
}

</script>
