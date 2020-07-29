<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToDatabase"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawDatabase } from '@client/ts/types/databases'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'
import DatabaseTypeRenderer from '@client/vue/shared/grid/DatabaseTypeRenderer.vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class DatabaseGrid extends Vue {
    @Prop({required: true})    
    databases!: RawDatabase[]

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
            DatabaseTypeRenderer,
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
                headerName: 'Type',
                sortable: true,
                filter: true,
                cellRenderer: 'DatabaseTypeRenderer',
            },
            {
                headerName: 'Version',
                field: 'Version',
                sortable: true,
                filter: true,
            },
        ]
    }

    get rowData() : any[] {
        return this.databases
    }

    onFirstDataRender(params : any) {
        params.columnApi.autoSizeAllColumns()
    }

    goToDatabase(e : RowEvent) {
        this.$router.push({
            name: 'databaseHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                databaseId: e.data.Id,
            },
        })
        this.$emit('change-database')
    }
}

</script>
