import ky from 'ky'
import { Store } from 'vuex'
import { RootState } from '@client/ts/stores/store'

const apiPrefix : string = '/api/v1'

export class ApiHttpHandler {
    store : Store<RootState>
    constructor(store : Store<RootState>) {
        this.store = store
    }

    get<T>(endpoint : string, options : any) : Promise<T | null> {
        return ky.get(`${apiPrefix}${endpoint}`, options).json().then(
            (resp : unknown) => {
                return <T>resp
            },
            (err : any) => {
                this.store.commit('errors/addApiError', err)
                return null
            },
        )
    }

    post<T>(endpoint : string, options: any) : Promise<T | null> {
        return ky.post(`${apiPrefix}${endpoint}`, options).json().then(
            (resp : unknown) => {
                return <T>resp
            },
            (err : any) => {
                console.log("BAD", err)
                this.store.commit('errors/addApiError', err)
                return null
            },
        )
    }

    put<T>(endpoint : string, options: any) : Promise<T | null> {
        return ky.put(`${apiPrefix}${endpoint}`, options).json().then(
            (resp : unknown) => {
                return <T>resp
            },
            (err : any) => {
                this.store.commit('errors/addApiError', err)
                return null
            },
        )
    }
}
