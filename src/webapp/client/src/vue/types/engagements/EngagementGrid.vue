<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToEngagement"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawEngagement } from '@client/ts/types/engagements'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'
import SimpleDateTimeRenderer from '@client/vue/shared/grid/SimpleDateTimeRenderer.vue'
import TrueFalseRenderer from '@client/vue/shared/grid/TrueFalseRenderer.vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class EngagementGrid extends Vue {
    @Prop({required: true})    
    engagements!: RawEngagement[]

    get gridOptions() : any {
        return {
            enableCellTextSelection: true,
            defaultColDef: {
                resizable: true,
            },
        }
    }

    get frameworkComponents() : any {
        return {
            SimpleDateTimeRenderer,
            TrueFalseRenderer,
        }
    }

    get columnDefs() : any[] {
        return [
            {
                headerName: 'Name',
                field: 'Name',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'Created Time',
                field: 'CreatedTime',
                sortable: true,
                filter: true,
                cellRenderer: 'SimpleDateTimeRenderer',
            },
            {
                headerName: 'Start Time',
                field: 'StartTime',
                sortable: true,
                filter: true,
                cellRenderer: 'SimpleDateTimeRenderer',
            },
            {
                headerName: 'End Time',
                field: 'EndTime',
                sortable: true,
                filter: true,
                cellRenderer: 'SimpleDateTimeRenderer',
            },
            {
                headerName: 'Closed',
                field: 'IsClosed',
                sortable: true,
                filter: true,
                cellRenderer: 'TrueFalseRenderer',
            },
        ]
    }

    get rowData() : any[] {
        return this.engagements
    }

    onFirstDataRender(params : any) {
        params.columnApi.autoSizeAllColumns()
    }

    goToEngagement(e : RowEvent) {
        this.$router.push({
            name: 'orgSingleEngagement',
            params: {
                orgId: this.$route.params.orgId,
                engId: e.data.Id,
            },
        })
        this.$emit('change-engagement')
    }
}

</script>
