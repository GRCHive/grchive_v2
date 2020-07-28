<template>
    <v-card flat>
        <div class="stat-container">
            <div class="d-flex justify-center icon-container">
                <v-icon x-large>
                    {{ icon }}
                </v-icon>
            </div>

            <div
                class="d-flex justify-center text-overline"
                v-for="stat in stats"
            >
                <span class="font-weight-bold">
                    {{ stat.value }}
                </span>
                &nbsp;{{ stat.text }}
            </div>
        </div>
    </v-card>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { StatInfo } from '@client/ts/stats'

@Component
export default class StatDisplay extends Vue {
    @Prop()
    icon!: string

    @Prop({ type: Array })
    stats! : StatInfo[]
}

</script>

<style scoped>

.stat-container {
    padding-top: 16px;
    padding-bottom: 16px;
}

.icon-container {
    margin-bottom: 16px;
}

>>>.text-overline {
    line-height: 1rem;
}

</style>
