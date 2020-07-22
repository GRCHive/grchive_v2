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
import { Prop } from 'vue-property-decorator'
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

    commentStr : string = ''

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
        return [Permission.PCommentsCreate]
    }

    submitComment() {
        GrchiveApi.comments.createComment(this.threadId, this.commentStr).then((resp : RawComment | null) => {
            if (!resp) {
                return
            }
            this.$emit('new-comment', resp)
            this.commentStr = ''
        })
    }
}

</script>
