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

export const getKittens = (page = 1, limit = 6, breed = "", sort = "") => {
  let url = `/kittens?page=${page}&limit=${limit}`;
  
  if (breed) {
    url += `&breed=${encodeURIComponent(breed)}`; 
  }
  if (sort) {
    url += `&sort=${sort}`;
  }
  
  return api.get(url).then((r) => r.data);
};

export const createKitten = (data) => api.post("/kittens", data).then((r) => r.data);
export const uploadImage = (file) => {
    const formData = new FormData();
    formData.append("image", file);
    return api.post("/upload", formData).then((r) => r.data);
};
export const login = (username, password) => api.post("/login", { username, password }).then((r) => r.data);
export const getAccessToken = () => accessToken;
export const logout = async () => {
   try{
       await api.post("/logout");
   }catch (e){
       console.error("Logout error", e);
   }

   accessToken = null;
   localStorage.removeItem("role");
};
export const refreshToken = () => api.post("/refresh", {}).then((r) => {
    setAccessToken(r.data.access_token); 
    return r.data;
});
export const getKitten = (id) => api.get(`/kittens/${id}`).then((r) => r.data);

export const registerUser = (username, email, password) => api.post("/register", {username, email, password}).then((r) => r.data);
export const loginWithCode = (username, code) => api.post("/login-with-code", {username, code}).then((r) => r.data);