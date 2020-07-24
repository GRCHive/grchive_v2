import { Permission } from '@client/ts/types/roles'
export interface NavLink {
    title : string
    icon : string
    hidden?: boolean
    disabled? : boolean
    children?: NavLink[]
    path? : string
    params? : any
    permissions?: Permission[]
}
