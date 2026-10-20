import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:3000/api/v1"
});

export const getKittens = () => api.get("/kittens").then((r) => r.data);
export const getKitten = (id) => api.get(`/kittens/${id}`).then((r) => r.data);
export const createKitten = (data) => api.post("/kittens", data).then((r) => r.data);