import ky from 'ky'
import { v4 as uuidv4 } from 'uuid'

const apiPrefix : string = '/api/v1'

const baseOptions : any = {
    retry: {
        limit: 1,
    },
    hooks: {
        beforeRequest: [
            (request : any) => {
                request.headers.set('Grchive-Request-Id', uuidv4())
            }
        ]
    }

}

export class ApiHttpHandler {
    get<T>(endpoint : string, options : any) : Promise<T> {
        return ky.get(`${apiPrefix}${endpoint}`, {
            ...baseOptions,
            ...options,
        }).json()
    }

    post<T>(endpoint : string, options: any) : Promise<T> {
        return ky.post(`${apiPrefix}${endpoint}`, {
            ...baseOptions,
            ...options,
        }).json()
    }

    put<T>(endpoint : string, options: any) : Promise<T> {
        return ky.put(`${apiPrefix}${endpoint}`, {
            ...baseOptions,
            ...options,
        }).json()
    }

    delete<T>(endpoint : string, options: any) : Promise<T> {
        return ky.delete(`${apiPrefix}${endpoint}`, {
            ...baseOptions,
            ...options,
        }).json()
    }
}
