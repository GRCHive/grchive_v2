<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToRisk"
        @selection-changed="onChangeSelection"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component, { mixins } from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawRisk } from '@client/ts/types/risks'
import { RowEvent, AgGridEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'
import RatingRenderer from '@client/vue/shared/grid/RatingRenderer.vue'
import RelationshipWrapperGridMixin from '@client/vue/types/relationships/RelationshipWrapperGridMixin.vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class RiskGrid extends mixins(RelationshipWrapperGridMixin) {
    @Prop()
    value!: RawRisk[]

    @Prop({required: true})    
    risks!: RawRisk[]

    @Prop({type: Boolean, default: false})
    selectable!: boolean

    get gridOptions() : any {
        return {
            enableCellTextSelection: true,
            defaultColDef: {
                resizable: true,
            },
        }
    }

    get frameworkComponents() : any {
        return this.convertFrameworkComponents({
            RatingRenderer,
        })
    }

    get columnDefs() : any[] {
        return this.convertColumnDefs([
            {
                headerName: 'ID',
                field: 'HumanId',
                sortable: true,
                filter: true,
                checkboxSelection: this.selectable,
            },
            {
                headerName: 'Name',
                field: 'Name',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'Severity',
                field: 'Severity',
                sortable: true,
                filter: true,
                cellRenderer: 'RatingRenderer',
            },
        ])
    }

    get rowData() : any[] {
        return this.convertRowData(this.risks)
    }

    onFirstDataRender(params : any) {
        params.columnApi.autoSizeAllColumns()
    }

    goToRisk(e : RowEvent) {
        if (this.selectable) {
            return
        }

        this.$router.push({
            name: 'riskHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                riskId: e.data.Id || e.data.Data.Id,
            },
        })
        this.$emit('change-risk')
    }

    onChangeSelection(event : AgGridEvent) {
        this.$emit('input', event.api.getSelectedRows().map((ele : any) => {
            return ele.Data || ele
        })
    }
}

</script>
