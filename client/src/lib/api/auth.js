import axios from "axios"
import {PUBLIC_BASE_API_URL} from '$env/static/public'

export const login = async (username, password) => {
    try {
        const res  = await axios.post(`${PUBLIC_BASE_API_URL}/login`, {username, password})
        localStorage.setItem("token", res.data.token)
        localStorage.setItem("username", username)

        console.log(res.data.message)
        window.location.href = "/dashboard"
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Error in logging in.")
    }
}

export const register = async(username, password) => {
    try {
        const res = await axios.post(`${PUBLIC_BASE_API_URL}/register`, {username, password})
        console.log(res.data.message)
        alert(res.data.message || "User registered successfully")
        window.location.href = "/login"
    } catch (error) {
        console.error(error)
        alert(error.response?.data.message || "Error in registering.")
    }
}