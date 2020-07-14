export class GrchiveError {
}

export class GrchiveApiError extends GrchiveError {
    readonly rawErr : any

    constructor(err : any) {
        super()
        this.rawErr = err
    }
}
