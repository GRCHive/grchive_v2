<template>
    <div>
        <hero-image
            :src="bannerImageUrl"
            :max-height=300
            position="center center"
            text="Let's Talk"
        ></hero-image>

        <v-row justify="center">
            <v-card width="30%" class="ma-4">
                <v-card-title>
                    Email Us
                </v-card-title>
                <v-card-text class="long-text">
                    <div class="long-text">
                        Send us an email and we will respond as soon as we can!
                    </div>

                    <div class="text--primary headline my-2">
                        Interested in our product?
                        <a :href="salesEmail.mailto">{{ salesEmail.email }}</a>
                    </div>

                    <div class="text--primary headline my-2">
                        Interested in working with us?
                        <a :href="careerEmail.mailto">{{ careerEmail.email }}</a>
                    </div>

                    <div class="text--primary headline my-2">
                        Need technical or audit support?
                        <a :href="supportEmail.mailto">{{ supportEmail.email }}</a>
                    </div>

                    <div class="text--primary headline my-2">
                        All other inquiries:
                        <a :href="contactEmail.mailto">{{ contactEmail.email }}</a>
                    </div>
                </v-card-text>
            </v-card>
        </v-row>
    </div>
</template>

<script lang="ts">

import HeroImage from '@client/vue/components/HeroImage.vue'
import { domainName, createAssetUrl, createMailtoUrl } from '@client/ts/url'

import Vue from 'vue'
import Component from 'vue-class-component'

@Component({
    components: {
        HeroImage
    },
})
export default class ContactUsPage extends Vue {
    readonly bannerImageUrl: string = createAssetUrl('generic-banner2.jpg')
    readonly salesEmail : Object  = createMailtoUrl('sales', domainName)
    readonly careerEmail : Object  = createMailtoUrl('careers', domainName)
    readonly supportEmail : Object  = createMailtoUrl('support', domainName)
    readonly contactEmail : Object  = createMailtoUrl('contact', domainName)
}

</script>
