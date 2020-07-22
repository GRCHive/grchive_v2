<template>
    <v-list-item>
        <v-list-item-content>
            <v-list-item-title>
                <div class="d-flex">
                    {{ userName }}
                    <span class="caption ml-4">
                         {{ standardFormatTime(comment.PostTime) }}
                    </span>
                    
                    <v-spacer></v-spacer>

                    <restrict-role-permission-button
                        color="success"
                        icon
                        x-small
                        :permissions="editPermissions"
                        :org-id="currentOrgId"
                        :engagement-id="currentEngagementId"
                        @click="editMode = true"
                    >
                        <v-icon>
                            mdi-pencil
                        </v-icon>
                    </restrict-role-permission-button>

                    <restrict-role-permission-button
                        color="error"
                        icon
                        x-small
                        class="ml-4"
                        :permissions="deletePermissions"
                        :org-id="currentOrgId"
                        :engagement-id="currentEngagementId"
                        @click="deleteComment"
                    >
                        <v-icon>
                            mdi-delete
                        </v-icon>
                    </restrict-role-permission-button>
                </div>
            </v-list-item-title>

            <v-list-item-subtitle>
                <pre v-if="!editMode">{{ comment.Content }}</pre>
                <comment-creator
                    v-else
                    :thread-id="threadId"
                    :value="comment"
                    @new-comment="onEditComment"
                >
                </comment-creator>
            </v-list-item-subtitle>
        </v-list-item-content>
    </v-list-item>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawComment, CommentThreadId } from '@client/ts/types/comments'
import { standardFormatTime } from '@client/ts/time'
import { GrchiveApi } from '@client/ts/main'
import { RawUser } from '@client/ts/types/users'
import { Permission } from '@client/ts/types/roles'
import RestrictRolePermissionButton from '@client/vue/loading/RestrictRolePermissionButton.vue'
import CommentCreator from '@client/vue/types/comments/CommentCreator.vue'

@Component({
    components: {
        RestrictRolePermissionButton,
        CommentCreator
    }
})
export default class SingleCommentViewer extends Vue {
    @Prop({ required: true })
    comment!: RawComment

    @Prop({ required: true })
    threadId! : CommentThreadId
    standardFormatTime  = standardFormatTime

    editMode: boolean = false

    cachedUserName : string | null = null
    get userName() : string {
        if (!this.cachedUserName) {
            GrchiveApi.orgs.getOrgUser(this.threadId.orgId, this.comment.UserId).then((resp : RawUser | null) => {
                if (!resp) {
                    this.cachedUserName = 'Unknown'
                    return
                }
                this.cachedUserName = `${resp.FullName} (${resp.Email})`
            })
            return 'Loading...'
        }
        return this.cachedUserName
    }

    get isSelf() : boolean {
        return this.comment.UserId === this.$store.state.user.rawUser.Id 
    }

    get editPermissions() : Permission[] {
        if (this.isSelf) {
            return []
        } else {
            return [Permission.PCommentsUpdate]
        }
    }

    get deletePermissions() : Permission[] {
        if (this.isSelf) {
            return []
        } else {
            return [Permission.PCommentsDelete]
        }
    }

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

    onEditComment(c : RawComment) {
        this.editMode = false

        // TODO: This is probably hacky and goes against what we're supposed to do.
        // In reality we need to sync the value back up the prop chain somehow.
        this.comment.Content = c.Content
    }

    deleteComment() {
        GrchiveApi.comments.deleteComment(this.threadId, this.comment.Id).then((resp : void | null) => {
            if (resp === null)  {
                return
            }
            this.$emit('do-delete')
        })
    }
}

</script>
