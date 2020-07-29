<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToDesktop"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawDesktop } from '@client/ts/types/inventory'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class DesktopGrid extends Vue {
    @Prop({required: true})    
    desktops!: RawDesktop[]

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
        ]
    }

    get rowData() : any[] {
        return this.desktops
    }

    onFirstDataRender(params : any) {
        params.columnApi.autoSizeAllColumns()
    }

    goToDesktop(e : RowEvent) {
        this.$router.push({
            name: 'desktopHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                desktopId: e.data.Id,
            },
        })
        this.$emit('change-desktop')
    }
}

</script>
