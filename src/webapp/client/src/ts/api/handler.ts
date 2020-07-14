import ky from 'ky'

const apiPrefix : string = '/api/v1'

export class ApiHttpHandler {
    get<T>(endpoint : string, options : any) : Promise<T> {
        return ky.get(`${apiPrefix}${endpoint}`, options).json()
    }

    post<T>(endpoint : string, options: any) : Promise<T> {
        return ky.post(`${apiPrefix}${endpoint}`, options).json()
    }
}
