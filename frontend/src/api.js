import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:3000/api/v1",
    withCredentials: true
});
let accessToken = null;

export const setAccessToken = (token) => {accessToken = token;}

api.interceptors.request.use((config) =>{
    if (accessToken){
        config.headers.Authorization = `Bearer ${accessToken}`;
    }
    return config;
})

api.interceptors.response.use(
    (response) => response,
    async (error) =>{
        const originalRequest = error.config;

        if (error.response?.status == 401 && !originalRequest._retry){
            originalRequest._retry = true;
            try{
                const res = await axios.post("http://localhost:3000/api/v1/refresh", {}, {withCredentials: true});
                setAccessToken(res.data.access_token);
                originalRequest.headers.Authorization = `Bearer ${res.data.access_token}`;
                return api(originalRequest);
            }catch (refreshError){
                setAccessToken(null);
                window.location.href = "/login";
                return Promise.reject(refreshError);
            }
        }

        return Promise.reject(error);
    }
)

export const getKittens = (page = 1, limit = 10) => api.get(`/kittens?page=${page}&limit=${limit}`).then((r) => r.data);
export const createKitten = (data) => api.post("/kittens", data).then((r) => r.data);
export const uploadImage = (file) => {
    const formData = new FormData();
    formData.append("image", file);
    return api.post("/upload", formData).then((r) => r.data);
};
export const login = (username, password) => api.post("/login", { username, password }).then((r) => r.data);
export const getAccessToken = () => accessToken;
export const logout = () => {
    accessToken = null;
    localStorage.removeItem("role");
};
export const refreshToken = () => api.post("/refresh", {}).then((r) => {
    setAccessToken(r.data.access_token); 
    return r.data;
});