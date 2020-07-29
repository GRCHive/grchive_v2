// Rules meant to be used with Vuetify form elements.
export function createMaxLength(len : number) : (_: string) => boolean | string {
    return (v : string) => !!v && v.length <= len || `Invalid input length, must have less than ${len} characters.`;
}

export function createMinLength(len : number) : (_: string) => boolean | string {
    return (v : string) => !!v && v.length >= len || `Invalid input length, must have more than ${len} characters.`;
}

export function nonZero(v : any) : boolean | string {
    return (v != 0) || "Invalid choice.";
}

export function required(v : any) : boolean | string {
    const tst : boolean = ((v !== null) && (
            (Array.isArray(v) && v.length > 0) || 
            (!Array.isArray(v) && 
                (v !== 0) &&
                (!v.trim || v.trim() !== ""))
        ))
    return tst || "Input required.";
}

export function email(v : string) : boolean | string {
    if (!v) {
        return true;
    }

    const regex : RegExp = /(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/;
    const match = v.match(regex);
    return !!match || "A valid email is required.";
}

export function numeric(v : string) : boolean | string {
    return (v !== "" && !isNaN(Number(v))) || "Must be numeric."
}

export function hasLowerCase(v : string) : boolean | string {
    return (/[a-z]/.test(v)) || "Must have lower case characters."
}

export function hasUpperCase(v : string) : boolean | string {
    return (/[A-Z]/.test(v)) || "Must have upper case characters."
}

export function hasNumeric(v : string) : boolean | string {
    return (/[0-9]/.test(v)) || "Must have numeric characters."
}

export function geq(cmp : number) : (_: string) => boolean | string {
    return (v : string) => Number(v) >= cmp || `Invalid value, must be greater than or equal to ${cmp}.`
}

export function createPerElement(fn : (arg0 : string) => boolean | string) : (arg0 : Array<string>) => boolean | string {
    return function(v: Array<string>) : boolean | string  {
        for (const ele of v) {
            const res = fn(ele)
            if (!res || typeof res == 'string') {
                return res + ` [${ele}]`
            }
        }
        return true
    }
}
