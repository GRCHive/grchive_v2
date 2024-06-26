<template>
    <div>
        <landing-carousel
            :imageUrl="bannerImageUrl"
            tag-line="Automated features for modern GRC."
            :description="`See how GRCHive can enhance your audit workflow and document handoff process today with secure database and cloud IPE integrations.`"
            position="center center"
            hide-learn-more
        >
        </landing-carousel>

        <p class="text-center text-uppercase display-1 font-weight-thin mt-12 text-learn-more-header">
            Key Features
        </p>

        <feature-highlight
            headline="Process Flow Visualization"
            description="Accelerate the learning curve of your business processes for all key stakeholders. Lengthy, confusing business narratives are transformed into visually pleasing and easy-to-make process flows that everyone can understand."
            id="flow"
        >
            <template v-slot:image>
                <v-img
                    src="/static/assets/icons/graph.svg"
                    width="250px"
                    max-width="250px"
                >
                </v-img>
            </template>
        </feature-highlight>

        <v-divider class="feature-divider"></v-divider>

        <feature-highlight
            headline="Risk and Control Management"
            description="Gain back time spent on maintaining risk matrices and other governance documentation by letting GRCHive do it for you. Changes to controls, risks, systems, vendors, and more update related downstream and upstream components to keep process relationships always up to date."
            flip
            id="riskcontrol"
        >
            <template v-slot:image>
                <v-img
                    src="/static/assets/icons/checklist.svg"
                    width="250px"
                    max-width="250px"
                >
                </v-img>
            </template>
        </feature-highlight>

        <v-divider class="feature-divider"></v-divider>

        <feature-highlight
            headline="Document Repository"
            description="Prevent common audit pains such as versioning issues and knowledge loss by having a single source of truth in GRCHive. Automatic versioning and backups allow all key stakeholders to rely on GRCHive for their documentation needs."
            id="repo"
        >
            <template v-slot:image>
                <v-img
                    src="/static/assets/icons/documents.svg"
                    width="250px"
                    max-width="250px"
                >
                </v-img>
            </template>
        </feature-highlight>

        <v-divider class="feature-divider"></v-divider>

        <feature-highlight
            headline="IT and Vendor Management"
            description="Increase clarity into your operational infrastructure by documenting the roles your audit infrastructure components play and how they connect to your organization's controls, risks, and processes."
            flip
            id="it"
        >
            <template v-slot:image>
                <v-img
                    src="/static/assets/icons/servers.svg"
                    width="250px"
                    max-width="250px"
                >
                </v-img>
            </template>
        </feature-highlight>

        <v-divider class="feature-divider"></v-divider>

        <feature-highlight
            headline="Database Integration"
            description="Realize potential audit efficiencies by giving auditors the ability to create and run approvable SQL requests. Furthermore, track database schema changes over time to take advantage of IT audit benchmarking. We will work with you to ensure your database of choice can be fully integrated into GRCHive."
            id="db"
        >
            <template v-slot:image>
                <v-img
                    src="/static/assets/icons/database.svg"
                    width="250px"
                    max-width="250px"
                >
                </v-img>
            </template>
        </feature-highlight>

        <v-divider class="feature-divider"></v-divider>

        <feature-highlight
            headline="Audit Requests"
            description="Minimize the time spent on creating updates for key stakeholders in status meetings. GRCHive tracks audit requests in real-time to make updating document request lists a thing of the past."
            flip
            id="req"
        >
            <template v-slot:image>
                <v-img
                    src="/static/assets/icons/negotiating.svg"
                    width="250px"
                    max-width="250px"
                >
                </v-img>
            </template>
        </feature-highlight>

        <v-divider class="feature-divider"></v-divider>

        <p class="text-center text-uppercase display-1 font-weight-thin mt-12 text-learn-more-header">
            Additional Features
        </p>

        <v-row justify="space-around" class="py-12">
            <v-col cols="3">
                <feature-summary
                    headline="Notifications"
                    description="Remove the hassle and stress of manually reminding key stakeholders of document request progress and other important events."
                    id="notification"
                >
                    <template v-slot:image>
                        <v-img
                            src="/static/assets/icons/bell.svg"
                            width="128px"
                            max-width="128px"
                        >
                        </v-img>
                    </template>
                </feature-summary>
            </v-col>

            <v-col cols="3">
                <feature-summary
                    headline="Access Controls"
                    description="Tighten data security by defining set roles for users and controlling what resources they can view, update, and create on GRCHive. Extensive audit trails are provided for all resources on GRCHive."
                    id="access"
                >
                    <template v-slot:image>
                        <v-img
                            src="/static/assets/icons/keywords.svg"
                            width="128px"
                            max-width="128px"
                        >
                        </v-img>
                    </template>
                </feature-summary>
            </v-col>

            <v-col cols="3">
                <feature-summary
                    headline="Security"
                    description="All the data transmitted to, from, and within GRCHive is encrypted using TLS to ensure your organization's confidential information is protected. Additionally, your documentation is are encrypted and signed before being stored on Google Cloud to ensure confidentiality and integrity."
                    id="security"
                >
                    <template v-slot:image>
                        <v-img
                            src="/static/assets/icons/security.svg"
                            width="128px"
                            max-width="128px"
                        >
                        </v-img>
                    </template>
                </feature-summary>
            </v-col>
        </v-row>

        <v-divider class="feature-divider"></v-divider>

        <p class="text-center text-uppercase display-1 font-weight-thin mt-12 text-learn-more-header">
            Coming Soon
        </p>

        <v-row justify="space-around" class="py-12">
            <v-col cols="3">
                <feature-summary
                    headline="Automated Scripts and Tests"
                    description="Increase confidence in your controls by eliminating human error and guaranteeing consistent, complete, and accurate results using schedulable custom Kotlin programs with full GRCHive integration."
                    id="automate"
                >
                    <template v-slot:image>
                        <v-img
                            src="/static/assets/icons/automate.svg"
                            width="128px"
                            max-width="128px"
                        >
                        </v-img>
                    </template>
                </feature-summary>
            </v-col>

            <v-col cols="3">
                <feature-summary
                    headline="Customizable Audit Workflows"
                    description="Reduce audit planning efforts by standardizing and automating the creation of your test plan templates."
                    id="workflow"
                >
                    <template v-slot:image>
                        <v-img
                            src="/static/assets/icons/flow.svg"
                            width="128px"
                            max-width="128px"
                        >
                        </v-img>
                    </template>
                </feature-summary>
            </v-col>

            <v-col cols="3">
                <feature-summary
                    headline="Cloud IPE Integrations"
                    description="Automatically pull data from your cloud IPE vendors such as Workday, SAP, Black Line, etc. and connect your data to requests and automated scripts. We will work with you to ensure your vendors are supported."
                    id="ipe"
                >
                    <template v-slot:image>
                        <v-img
                            src="/static/assets/icons/cloud-computing.svg"
                            width="128px"
                            max-width="128px"
                        >
                        </v-img>
                    </template>
                </feature-summary>
            </v-col>
        </v-row>

        <v-divider class="feature-divider"></v-divider>

        <standalone-cta class="max-height" id="bottomCta">
        </standalone-cta>
    </div>
</template>

<script lang="ts">

import LandingCarousel from '@client/vue/components/LandingCarousel.vue'
import FeatureHighlight from '@client/vue/components/FeatureHighlight.vue'
import FeatureSummary from '@client/vue/components/FeatureSummary.vue'
import StandaloneCta from '@client/vue/components/StandaloneCta.vue'
import { createAssetUrl } from '@client/ts/url'

import Vue from 'vue'
import Component from 'vue-class-component'

@Component({
    components: {
        LandingCarousel,
        FeatureHighlight,
        FeatureSummary,
        StandaloneCta
    },
})
export default class LearnMorePage extends Vue {
    readonly bannerImageUrl : string = createAssetUrl('generic-banner4.jpg')
}

</script>

<style scoped>

.text-learn-more-header {
    color: #1976D2;
}

</style>
