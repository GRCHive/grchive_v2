<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToVendor"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawVendor } from '@client/ts/types/vendors'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class VendorGrid extends Vue {
    @Prop({required: true})    
    vendors!: RawVendor[]

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
                field: 'Name',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'URL',
                field: 'Url',
                sortable: true,
                filter: true,
            },
        ]
    }

    get rowData() : any[] {
        return this.vendors
    }

    onFirstDataRender(params : any) {
        params.columnApi.autoSizeAllColumns()
    }

    goToVendor(e : RowEvent) {
        this.$router.push({
            name: 'vendorHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                vendorId: e.data.Id,
            },
        })
        this.$emit('change-vendor')
    }
}

</script>
