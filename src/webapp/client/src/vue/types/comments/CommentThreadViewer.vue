<template>
    <v-list two-line>
        <template v-for="(item, index) in comments">
            <single-comment-viewer
                :comment="item"
                :thread-id="threadId"
                :key="`comment-${item.Id}`"
                @do-delete="deleteComment(item)"
            >
            </single-comment-viewer>

            <v-divider
                :key="`divider-${item.Id}`"
                v-if="index != comments.length - 1"
            >
            </v-divider>
        </template>

    </v-list>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawComment, CommentThreadId } from '@client/ts/types/comments'
import SingleCommentViewer from '@client/vue/types/comments/SingleCommentViewer.vue'

@Component({
    components: {
        SingleCommentViewer
    }
})
export default class CommentThreadViewer extends Vue {
    @Prop({ type: Array, default: []})
    comments!: RawComment[]

    @Prop({ required: true })
    threadId! : CommentThreadId

    deleteComment(comment : RawComment) {
        this.$emit('update:comments', this.comments.filter((ele : RawComment) => ele.Id != comment.Id))
    }
}

</script>
