<template>
    <ag-grid-vue
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        :framework-components="frameworkComponents"
        :grid-options="gridOptions"
        @first-data-rendered="onFirstDataRender"
        @row-clicked="goToSystem"
    >
    </ag-grid-vue>
</template>

<script lang="ts">

import Vue from 'vue'
import Component from 'vue-class-component'
import { Prop } from 'vue-property-decorator'
import { RawSystem } from '@client/ts/types/systems'
import { RowEvent } from 'ag-grid-community'
import { AgGridVue } from 'ag-grid-vue'

@Component({
    components: {
        AgGridVue,
    }
})
export default class SystemGrid extends Vue {
    @Prop({required: true})    
    systems!: RawSystem[]

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
                field: 'Name',
                sortable: true,
                filter: true,
            },
            {
                headerName: 'Purpose',
                field: 'Purpose',
                sortable: true,
                filter: true,
            },
        ]
    }

    get rowData() : any[] {
        return this.systems
    }

    onFirstDataRender(params : any) {
        params.api.sizeColumnsToFit()
    }

    goToSystem(e : RowEvent) {
        this.$router.push({
            name: 'systemHome',
            params: {
                orgId: this.$route.params.orgId,
                engId: this.$route.params.engId,
                systemId: e.data.Id,
            },
        })
        this.$emit('change-system')
    }
}

</script>
