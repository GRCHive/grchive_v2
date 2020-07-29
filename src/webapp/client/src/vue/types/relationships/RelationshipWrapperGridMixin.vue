<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RelationshipWrapper } from '@client/ts/types/relationships'
import TrueFalseRenderer from '@client/vue/shared/grid/TrueFalseRenderer.vue'

@Component
export default class RelationshipWrapperGridMixin extends Vue {
    @Prop({default : null})
    rels!: RelationshipWrapper<any>[] | null

    convertFrameworkComponents(comp : any) : any {
        comp.TrueFalseRenderer = TrueFalseRenderer
        return comp
    }

    convertColumnDefs(defs : any[]) : any[] {
        if (!this.rels) {
            return defs
        }

        return defs.map((ele : any) => {
            if (!!ele.field) {
                ele.field = 'Data.' + ele.field
            }
            return ele
        }).concat([{
            headerName: 'Explicit Relationship',
            field: 'Explicit',
            sortable: true,
            filter: true,
            cellRenderer: 'TrueFalseRenderer',
        }])
    }

    convertRowData(data : any[]) : any[] {
        if (!this.rels) {
            return data
        }
        return this.rels
    }
}

</script>
