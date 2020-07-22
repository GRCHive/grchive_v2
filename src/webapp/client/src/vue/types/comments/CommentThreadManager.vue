<template>
    <loading-container :loading="!allComments">
        <template v-slot:default="{show}">
            <v-row v-if="show">
                <v-col cols="8">
                    <comment-thread-viewer
                        :comments="allComments"
                        :thread-id="threadId"
                    >
                    </comment-thread-viewer>
                </v-col>

                <v-col cols="4">
                    <span class="h5 font-weight-bold">Submit a Comment</span>
                    <v-divider class="mb-4"></v-divider>
                    <comment-creator
                        :thread-id="threadId"
                        @new-comment="onNewComment"
                    >
                    </comment-creator>
                </v-col>
            </v-row>
        </template>
    </loading-container>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import CommentThreadViewer from '@client/vue/types/comments/CommentThreadViewer.vue'
import CommentCreator from '@client/vue/types/comments/CommentCreator.vue'
import { Watch, Prop } from 'vue-property-decorator'
import { RawComment, CommentThreadId } from '@client/ts/types/comments'
import { GrchiveApi } from '@client/ts/main'

@Component({
    components: {
        LoadingContainer,
        CommentThreadViewer,
        CommentCreator
    }
})
export default class CommentThreadManager extends Vue {
    @Prop({ required: true })
    threadId! : CommentThreadId

    allComments: RawComment[] | null = null

    @Watch('threadId')
    refreshThread() {
        GrchiveApi.comments.listComments(this.threadId).then((resp : RawComment[] | null) => {
            this.allComments = resp
        })
    }

    onNewComment(c : RawComment) {
        if (!this.allComments) {
            return
        }
        this.allComments!.push(c)
    }

    mounted() {
        this.refreshThread()
    }
}

</script>
