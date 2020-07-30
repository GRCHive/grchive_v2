<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToLaptop"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawLaptop } from '@client/ts/types/inventory'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'
import MachineStateInfoRenderer from '@client/vue/shared/grid/MachineStateInfoRenderer.vue'
import MachineStateTypeRenderer from '@client/vue/shared/grid/MachineStateTypeRenderer.vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class LaptopGrid extends Vue {
    @Prop({required: true})    
    laptops!: RawLaptop[]

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
            MachineStateInfoRenderer,
            MachineStateTypeRenderer
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
                headerName: 'State Type',
                field: 'State',
                sortable: true,
                filter: true,
                cellRenderer: 'MachineStateTypeRenderer'
            },
            {
                headerName: 'State Info',
                field: 'State',
                sortable: true,
                filter: true,
                cellRenderer: 'MachineStateInfoRenderer'
            },

        ]
    }

    get rowData() : any[] {
        return this.laptops
    }

    onFirstDataRender(params : any) {
        params.columnApi.autoSizeAllColumns()
    }

    goToLaptop(e : RowEvent) {
        this.$router.push({
            name: 'laptopHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                laptopId: e.data.Id,
            },
        })
        this.$emit('change-laptop')
    }
}

</script>
