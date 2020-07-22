<template>
    <loading-container :loading="!allComments">
        <template v-slot:default="{show}">
            <v-row v-if="show">
                <v-col cols="8">
                </v-col>

                <v-col cols="4">
                </v-col>
            </v-row>
        </template>
    </loading-container>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import LoadingContainer from '@client/vue/loading/LoadingContainer.vue'
import { Watch, Prop } from 'vue-property-decorator'
import { RawComment, CommentThreadId } from '@client/ts/types/comments'
import { GrchiveApi } from '@client/ts/main'

@Component({
    components: {
        LoadingContainer
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

    mounted() {
        this.refreshThread()
    }
}

</script>
