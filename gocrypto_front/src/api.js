import axios from "axios";

const baseurl = "http://localhost:3000";

const api = axios.create({
  baseURL: baseurl,
});

export default api;
