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
                headerName: 'Likelihood',
                field: 'Likelihood',
                sortable: true,
                filter: true,
                cellRenderer: 'RatingRenderer',
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
