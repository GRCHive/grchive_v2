<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToRisk"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawRisk } from '@client/ts/types/risks'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'
import RatingRenderer from '@client/vue/shared/grid/RatingRenderer.vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class EngagementGrid extends Vue {
    @Prop({required: true})    
    risks!: RawRisk[]

    get gridOptions() : any {
        return {
            enableCellTextSelection: true,
        }
    }

    get frameworkComponents() : any {
        return {
            RatingRenderer,
        }
    }

    get columnDefs() : any[] {
        return [
            {
                headerName: 'ID',
                field: 'HumanId',
                sortable: true,
                filter: true,
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
        ]
    }

    get rowData() : any[] {
        return this.risks
    }

    onFirstDataRender(params : any) {
        params.api.sizeColumnsToFit()
    }

    goToRisk(e : RowEvent) {
        this.$router.push({
            name: 'riskHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                riskId: e.data.Id,
            },
        })
        this.$emit('change-risk')
    }
}

</script>
