<template>
    <v-list-item>
        <v-list-item-content>
            <v-list-item-title>
                {{ userName }}
                <span class="caption">
                     {{ standardFormatTime(comment.PostTime) }}
                </span>
            </v-list-item-title>

            <v-list-item-subtitle>
                <pre>{{ comment.Content }}</pre>
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

@Component
export default class SingleCommentViewer extends Vue {
    @Prop({ required: true })
    comment!: RawComment

    @Prop({ required: true })
    threadId! : CommentThreadId

    standardFormatTime  = standardFormatTime

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
}

</script>
