<template>
    <div>
        <v-textarea
            label="Comment"
            v-model="commentStr"
            filled
            hide-details
        >
        </v-textarea>

        <div class="d-flex mt-4">
            <v-spacer></v-spacer>

            <restrict-role-permission-button
                :permissions="submitPermissions"
                :org-id="currentOrgId"
                :engagement-id="currentEngagementId"
                :loading="inProgress"
                color="success"
                @click="submitComment"
            >
                Submit
            </restrict-role-permission-button>
        </div>
    </div>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop, Watch } from 'vue-property-decorator'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import { RawComment, CommentThreadId } from '@client/ts/types/comments'
import { Permission } from '@client/ts/types/roles'
import { GrchiveApi } from '@client/ts/main'

@Component({
    components: {
        RestrictRolePermissionButton,
    }
})
export default class CommentCreator extends Vue {
    @Prop({ required: true })
    threadId! : CommentThreadId

    @Prop()
    value! : RawComment

    commentStr : string = ''
    inProgress: boolean = false

    get currentOrgId() : number {
        const org = this.$store.state.org.rawOrg
        if (!org) {
            return -1
        }
        return org.Id
    }

    get currentEngagementId() : number {
        const eng = this.$store.state.engagements.rawEngagement
        if (!eng) {
            return -1
        }
        return eng.Id
    }

    get submitPermissions(): Permission[] {
        if (!!this.value) {
            return [Permission.PCommentsCreate]
        } else {
            return [Permission.PCommentsUpdate]
        }
    }

    onSuccess(resp : RawComment) {
        this.$emit('new-comment', resp)
        this.commentStr = ''
    }

    submitComment() {
        this.inProgress = true
        if (!!this.value) {
            GrchiveApi.comments.updateComment(this.threadId, this.value.Id, this.commentStr).then(this.onSuccess).finally(() => {
                this.inProgress = false
            })
        } else {
            GrchiveApi.comments.createComment(this.threadId, this.commentStr).then(this.onSuccess).finally(() => {
                this.inProgress = false
            })
        }
    }

    @Watch('value')
    syncFromValue() {
        if (!!this.value) {
            this.commentStr = this.value.Content
        } else {
            this.commentStr = ''
        }
    }

    mounted() {
        this.syncFromValue()
    }
}

</script>
