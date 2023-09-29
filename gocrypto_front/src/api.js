import axios from 'axios';

const baseurl = 'http://localhost:8080';


const api = axios.create({
    baseURL: baseurl,
})


export default api;