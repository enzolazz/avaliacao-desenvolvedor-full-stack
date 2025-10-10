import axios from "axios";

export const api = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_SERVER + "/api",
  headers: {
    "Content-Type": "application/json",
  },
});

export const redirect = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_SERVER + "/redirect",
  headers: {
    "Content-Type": "application/json",
  },
});
