<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToServer"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawServer } from '@client/ts/types/inventory'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class ServerGrid extends Vue {
    @Prop({required: true})    
    servers!: RawServer[]

    get gridOptions() : any {
        return {
            enableCellTextSelection: true,
        }
    }

    get frameworkComponents() : any {
        return {
        }
    }

    get columnDefs() : any[] {
        return [
            {
                headerName: 'Name',
                field: 'Inventory.Name',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'Brand',
                field: 'Inventory.Brand',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'Model',
                field: 'Inventory.Model',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'Location',
                field: 'PhysicalLocation',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'OS',
                field: 'OperatingSystem',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'Hypervisor',
                field: 'Hypervisor',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'External IP',
                field: 'StaticExternalIp',
                sortable: true,
                filter: true,
            },
        ]
    }

    get rowData() : any[] {
        return this.servers
    }

    onFirstDataRender(params : any) {
        params.api.sizeColumnsToFit()
    }

    goToServer(e : RowEvent) {
        this.$emit('change-server')
    }
}

</script>
