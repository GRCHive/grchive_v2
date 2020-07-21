<template>
    <v-navigation-drawer
        app
        clipped
        :mini-variant="mini"
        :expand-on-hover="mini"
        mini-variant-width="50"
        width="256"
        permanent
    >
        <slot></slot>
        <v-list-item
            v-if="!!backLink"
            dense
            :to="{name: `${backLink.path}`, params: `${backLink.params}` }"
            link
            color="secondary"
            id="navBarHeader"
        >
            <v-list-item-icon>
                <v-icon>mdi-keyboard-return</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
                <v-list-item-title>
                    Back to {{ backLink.title }}
                </v-list-item-title>
            </v-list-item-content>
        </v-list-item>

        <v-list class="py-0" expand>
            <div
                v-for="(item, i) in finalNavLinks" 
                :key="i"
                :style="!!item.hidden ? `display: none;` : ``"
            >
                <v-list-group
                    v-if="!!item.children && item.children.length > 0"
                    :prepend-icon="item.icon"
                    no-action
                    value="true"
                    color="rgba(0, 0, 0, 0.87) !important"
                >

                    <template v-slot:activator>
                        <v-list-item-title>{{ item.title }} </v-list-item-title>
                    </template>

                    <generic-nav-bar-item
                        v-for="(child, ci) in item.children"
                        :key="ci"
                        :item="child"
                    >
                    </generic-nav-bar-item>
                </v-list-group>

                <generic-nav-bar-item
                    :item="item"
                    v-else
                >
                </generic-nav-bar-item>
            </div>
        </v-list>
    </v-navigation-drawer>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Watch, Prop } from 'vue-property-decorator'
import { NavLink } from '@client/ts/frontend/navLink'
import GenericNavBarItem from '@client/vue/navbar/GenericNavBarItem.vue'

@Component({
    components: {
        GenericNavBarItem,
    }
})
export default class GenericNavBar extends Vue {
    @Prop({ required: true })
    navLinks!: NavLink[]

    @Prop({ default: null })
    backLink!: NavLink | null

    get mini() : boolean {
        return this.$store.state.appLayout.miniNav
    }

    get finalNavLinks() : Array<NavLink> {
        if (!this.mini) {
            return this.navLinks
        } else {
            let fullArray = new Array<NavLink>()

            let expandFn = (ele : any) => {
                if (!!ele.children && ele.children.length > 0) {
                    let newParent = Object.assign({}, ele)
                    newParent.children = []
                    fullArray.push(newParent)
                    ele.children.forEach(expandFn)
                } else {
                    fullArray.push(ele)
                }
            }

            this.navLinks.forEach(expandFn)
            return fullArray
        }
    }
}

</script>

<style scoped>

>>>.v-list-group__header .v-list-group__header__prepend-icon .v-icon {
    color: rgba(0, 0, 0, 0.54) !important;
}

>>>.v-navigation-drawer__content {
    overflow-y: hidden;
}

>>>.v-navigation-drawer__content:hover {
    overflow-y: auto;
}

#navBarHeader {
    background-color: #DCDCDC;
}

</style>
