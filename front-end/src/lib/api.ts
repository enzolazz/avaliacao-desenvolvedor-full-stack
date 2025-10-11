import { apiClient, ApiError } from "@/api/client";
import { useAuth } from "@/hooks/use-auth";
import axios, { AxiosError } from "axios";
import { toast } from "sonner";

export const api = axios.create({
  baseURL: import.meta.env.VITE_BACKEND_SERVER + "/api",
  withCredentials: true,
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

let isRefreshing = false;
let failedQueue: { resolve: () => void; reject: (err: any) => void }[] = [];

function processQueue(error: any = null) {
  failedQueue.forEach((prom) => {
    if (error) prom.reject(error);
    else prom.resolve();
  });
  failedQueue = [];
}

api.interceptors.response.use(
  (response) => response,
  async (error: AxiosError) => {
    if (error.response?.status !== 401) {
      throw error;
    }

    const originalRequest = error.config!;
    if (originalRequest.url === "/auth/refresh") {
      await apiClient.auth.logout();
      localStorage.removeItem("user");

      window.location.href = "/auth";

      return Promise.reject(error);
    }

    if (isRefreshing) {
      return new Promise((resolve, reject) => {
        failedQueue.push({
          resolve: () => resolve(api(originalRequest)),
          reject,
        });
      });
    }

    isRefreshing = true;
    try {
      await api.post("/auth/refresh");

      processQueue();
      return api(originalRequest);
    } catch (refreshError) {
      processQueue(refreshError);
    } finally {
      isRefreshing = false;
    }
  },
);
