export const BASE_URL = "http://68.183.226.223:20005/"

function http(url, body, opts = {}) {
    const headers = {}

    if (opts.token) {
        headers.Authorization = "Bearer " + opts.token
    }

    const config = {
        method: opts.method || "GET",
        headers
    }

    if (body && (opts.method === "POST" || opts.method === "PUT" || opts.method === "PATCH")) {
        config.body = typeof body === "string" ? body : JSON.stringify(body)
        headers["Content-Type"] = "application/json"
    }

    return fetch(BASE_URL + url, config)
}

export default http