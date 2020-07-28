<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToMobile"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawMobile } from '@client/ts/types/inventory'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class MobileGrid extends Vue {
    @Prop({required: true})    
    mobile!: RawMobile[]

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
                headerName: 'OS',
                field: 'OperatingSystem',
                sortable: true,
                filter: true,
            },
        ]
    }

    get rowData() : any[] {
        return this.mobile
    }

    onFirstDataRender(params : any) {
        params.api.sizeColumnsToFit()
    }

    goToMobile(e : RowEvent) {
        this.$router.push({
            name: 'mobileHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                mobileId: e.data.Id,
            },
        })
        this.$emit('change-mobile')
    }
}

</script>
