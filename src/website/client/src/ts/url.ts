export const getStartedUrl : string = "/getting-started"
export const contactUsUrl : string = "/contact-us"
export const homePageUrl : string = "/"
export const blogUrl : string = "https://blog.grchive.com"
export const learnMoreUrl : string = "/learn"

export function createAssetUrl(asset : string) : string {
    return "/static/assets/" + asset
}
