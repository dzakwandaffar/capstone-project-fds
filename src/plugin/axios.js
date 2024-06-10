import axios from 'axios'

const myAxios = axios.create({
    baseURL: import.meta.env.VITE_BASE_API
})

export default myAxios;