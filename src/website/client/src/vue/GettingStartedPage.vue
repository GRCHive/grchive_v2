<template>
    <div>
        <hero-image
            :src="bannerImageUrl"
            :max-height=300
            position="center center"
            text="Get Started"
        ></hero-image>

        <div class="ma-4" id="hubspotForm"></div>
    </div>
</template>

<script lang="ts">

import HeroImage from '@client/vue/components/HeroImage.vue'
import { createAssetUrl } from '@client/ts/url'

import Vue from 'vue'
import Component from 'vue-class-component'

@Component({
    components: {
        HeroImage
    },
})
export default class GettingStartedPage extends Vue {
    readonly bannerImageUrl: string = createAssetUrl('generic-banner1.jpg')
}

</script>
