import axios from "axios"
import { useAuth } from "../Auth/useAuth"


const axiosInstance = axios.create({
    baseURL: import.meta.env.VITE_URL_API ?? "http://localhost:5000"
})

axiosInstance.interceptors.request.use((config) => {
    const { me } = useAuth()
    const token = me?.token
    if (token){
        config.headers.Authorization = `Bearer ${token}`
    }
    return config
})

export default axiosInstance