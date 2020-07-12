export const domainName : string = "grchive.com"
export const getStartedUrl : string = "/getting-started"
export const contactUsUrl : string = "/contact-us"
export const homePageUrl : string = "/"
export const blogUrl : string = "https://blog.grchive.com"
export const learnMoreUrl : string = "/learn"

export function createAssetUrl(asset : string) : string {
    return "/static/assets/" + asset
}

export function createEmailAddress(user : string, domain : string) : string {
    return user + "@" + domain
}

export function createMailtoUrl(user : string, domain : string) : Object {
    const email = createEmailAddress(user, domain)
    return Object.freeze({
        mailto: "mailto:" + email,
        email: email
    })
}
