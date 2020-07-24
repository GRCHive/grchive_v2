import { Store } from 'vuex'
import { RootState } from '@client/ts/stores/store'
import { v4 as uuidv4 } from 'uuid'

export enum GrchiveErrorCodes {
	GECNoError    = 0,
	GECBadRequest = 1,
	GECDbInternal = 2,
	GECDbExternal = 3,
	GECUnauthorized = 10000,
    Generic = 99999,
}

export class ErrorWrapper {
    displayId: string
    code : GrchiveErrorCodes
    message: string
    technical : string
    httpCode : number
    context: string = ''

    constructor(code : GrchiveErrorCodes, message : string, err : any, httpCode : number) {
        this.displayId = uuidv4()
        this.code = code
        this.message = message
        this.technical = btoa(JSON.stringify(err))
        this.httpCode = httpCode
    }
}

interface ServerErrorData {
    GrchiveCode : GrchiveErrorCodes
    GrchiveMessage: string
}

async function constructErrorWrapperFromGenericError(err : any) : Promise<ErrorWrapper> {
    let ret : ErrorWrapper | null = null
    if (!!err.response) {
        let data : ServerErrorData | null  = null
        try {
            data = await err.response.json()
        } catch {
            data = null
        }

        ret = new ErrorWrapper(
            !!data ? data.GrchiveCode : GrchiveErrorCodes.Generic,
            !!data ? data.GrchiveMessage : 'Unknown',
            {
                headers: Object.fromEntries(err.response.headers.entries()),
                url: err.response.url, 
            },
            err.response.status,
        )
    } else {
        ret = new ErrorWrapper(GrchiveErrorCodes.Generic, 'Generic Error.', err, 418)
    }

    return new Promise( (resolved) => {
        resolved(ret!)
    })
}

export interface PopupErrorOptions {
    callback? : () => void
    context?: string 
}

export class GrchiveErrorHandler {
    store : Store<RootState>
    router : any 

    constructor(store : Store<RootState>, router: any) {
        this.store = store
        this.router = router
    }

    failurePageOnError(err : any) {
        constructErrorWrapperFromGenericError(err).then((e : ErrorWrapper) => {
            this.router.replace({
                path: '/error',
                query: e,
            })
        })
    }

    failurePopupOnError(err :any, opts? : PopupErrorOptions) {
        constructErrorWrapperFromGenericError(err).then((e : ErrorWrapper) => {
            if (!!opts) {
                if (!!opts.callback) {
                    opts.callback()
                }

                if (!!opts.context) {
                    e.context = opts.context
                }
            }
            this.store.commit('errors/addError', e)
        })
    }
}
