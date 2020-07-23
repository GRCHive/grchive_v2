import ky from 'ky'

const apiPrefix : string = '/api/v1'

const baseOptions : any = {
    retry: {
        limit: 1
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
