export function parseJWT(jwt: string) {
    const jwtSections = jwt.split(".")
    const jwtBodyStringB64 = jwtSections[1]

    // base64 decode
    const jwtBodyString = atob(jwtBodyStringB64)

    // parse json
    const jwtBody = JSON.parse(jwtBodyString)

    return jwtBody
}