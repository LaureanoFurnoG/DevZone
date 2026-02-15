import axios from "axios"

const axiosInstance = axios.create({
    baseURL: import.meta.env.VITE_URL_API ?? "http://localhost:5000"
})

axiosInstance.interceptors.request.use((config) => {
    const token = sessionStorage.getItem('auth')
    if (token){
        config.headers.Authorization = `Bearer ${token}`
    }
    return config
})

export default axiosInstance