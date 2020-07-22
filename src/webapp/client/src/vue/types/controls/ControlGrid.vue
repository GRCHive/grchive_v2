<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToControl"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawControl } from '@client/ts/types/controls'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'
import RatingRenderer from '@client/vue/shared/grid/RatingRenderer.vue'
import TrueFalseRenderer from '@client/vue/shared/grid/TrueFalseRenderer.vue'
import ControlTypeRenderer from '@client/vue/shared/grid/ControlTypeRenderer.vue'
import UserIdRenderer from '@client/vue/shared/grid/UserIdRenderer.vue'
import ControlFrequencyRenderer from '@client/vue/shared/grid/ControlFrequencyRenderer.vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class ControlGrid extends Vue {
    @Prop({required: true})    
    controls!: RawControl[]

    get gridOptions() : any {
        return {
            enableCellTextSelection: true,
        }
    }

    get frameworkComponents() : any {
        return {
            RatingRenderer,
            TrueFalseRenderer,
            ControlTypeRenderer,
            UserIdRenderer,
            ControlFrequencyRenderer
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
                headerName: 'Type',
                field: 'ControlType',
                sortable: true,
                filter: true,
                cellRenderer: 'ControlTypeRenderer',
            },
            {
                headerName: 'Control Owner',
                field: 'OwnerId',
                sortable: true,
                filter: true,
                cellRenderer: 'UserIdRenderer',
            },
            {
                headerName: 'Failure Likelihood',
                field: 'Likelihood',
                sortable: true,
                filter: true,
                cellRenderer: 'RatingRenderer',
            },
            {
                headerName: 'Frequency',
                sortable: true,
                filter: true,
                cellRenderer: 'ControlFrequencyRenderer',
            },
            {
                headerName: 'Is Manual',
                field: 'IsManual',
                sortable: true,
                filter: true,
                cellRenderer: 'TrueFalseRenderer',
            },
        ]
    }

    get rowData() : any[] {
        return this.controls
    }

    onFirstDataRender(params : any) {
        params.api.sizeColumnsToFit()
    }

    goToControl(e : RowEvent) {
        this.$router.push({
            name: 'controlHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                controlId: e.data.Id,
            },
        })
        this.$emit('change-control')
    }
}

</script>
