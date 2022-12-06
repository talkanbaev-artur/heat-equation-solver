const apiEndpoint = "http://localhost:8000/api" //leave empty for prod

import axios from "axios"


const getTypes = () => {
    return axios.get(apiEndpoint + "/numericals")
}

const getSchemas = () => {
    return axios.get(apiEndpoint + "/schemas")
}

const getInitials = (data) => {
    return axios.post(apiEndpoint + "/initial", data)
}

const getSolution = (params, timePoint, id) => {
    return axios.post(apiEndpoint + "/solution", { t: timePoint, params, id })
}

export default { getSchemas, getInitials, getSolution }