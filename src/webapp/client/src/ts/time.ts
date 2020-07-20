import format from 'date-fns/format'

export function standardFormatTime(dt : Date | null) : string {
    if (!dt) {
        return "None"
    }
    return format(dt, 'MMM do uuuu, h:mm:ss a')
}
