export interface NavLink {
    title : string
    icon : string
    hidden?: boolean
    disabled? : boolean
    children?: NavLink[]
    path? : string
}
