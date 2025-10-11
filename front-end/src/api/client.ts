import { api, redirect } from "@/lib/api";
import { isAxiosError } from "axios";
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
} from "./types/auth";
import type {
  RedirectResponse,
  ShortenRequest,
  ShortenResponse,
} from "./types/url";
import type { MetricData } from "./types/metric";
import type { ShortURL } from "@/types/url";

export class ApiError extends Error {
  status: number;
  constructor(message: string, status: number) {
    super(message);
    this.name = "ApiError";
    this.status = status;
  }
}

export const apiClient = {
  auth: {
    async login(data: LoginRequest): Promise<LoginResponse> {
      try {
        const response = await api.post<LoginResponse>("/auth/login", data);
        return response.data;
      } catch (error: unknown) {
        returnErrors(error);

        throw new Error(
          error instanceof Error
            ? `Erro inesperado: ${error.message}`
            : "Erro inesperado ao realizar login.",
        );
      }
    },

    async register(data: RegisterRequest): Promise<RegisterResponse> {
      try {
        const response = await api.post<RegisterResponse>(
          "/auth/register",
          data,
        );
        return response.data;
      } catch (error: unknown) {
        returnErrors(error);

        throw new Error(
          error instanceof Error
            ? `Erro inesperado: ${error.message}`
            : "Erro inesperado ao realizar cadastro.",
        );
      }
    },

    async logout() {
      try {
        await api.post("/auth/logout");
      } catch (error: unknown) {
        returnErrors(error);

        throw new Error(
          error instanceof Error
            ? `Erro inesperado: ${error.message}`
            : "Erro inesperado ao realizar logout.",
        );
      }
    },
  },
  url: {
    async shorten(data: ShortenRequest): Promise<ShortenResponse> {
      try {
        const response = await api.post<ShortenResponse>("/shorten", data);

        return response.data;
      } catch (error: unknown) {
        returnErrors(error);

        throw new Error(
          error instanceof Error
            ? `Erro inesperado: ${error.message}`
            : "Erro inesperado ao encurtar URL.",
        );
      }
    },
    async getAllLinks(): Promise<ShortURL[]> {
      try {
        const response = await api.get<ShortURL[]>("/shorten");

        return response.data;
      } catch (error: unknown) {
        returnErrors(error);

        throw new Error(
          error instanceof Error
            ? `Erro inesperado: ${error.message}`
            : "Erro inesperado ao buscar URLs.",
        );
      }
    },

    async redirect(id: string): Promise<RedirectResponse> {
      try {
        const response = await redirect.get<RedirectResponse>("/" + id);

        return response.data;
      } catch (error: unknown) {
        returnErrors(error);

        throw new Error(
          error instanceof Error
            ? `Erro inesperado: ${error.message}`
            : "Erro inesperado ao buscar URLs.",
        );
      }
    },
  },
  metrics: {
    async getMetrics(id: string, filter: string): Promise<MetricData[]> {
      try {
        const { data } = await api.get<MetricData[]>(
          `/metrics/${filter}/${id}`,
        );

        return data;
      } catch (error: unknown) {
        returnErrors(error);

        throw new Error(
          error instanceof Error
            ? `Erro inesperado: ${error.message}`
            : "Erro inesperado ao encurtar URL.",
        );
      }
    },
  },
};

function returnErrors(error: unknown) {
  if (isAxiosError(error)) {
    if (error.response) {
      const status = error.response.status;
      const message = error.response.data?.error ?? "Erro desconhecido";
      const formattedMessage =
        typeof message === "string" && message.length > 0
          ? `${message.charAt(0).toUpperCase()}${message.slice(1)}.`
          : "Ocorreu um erro inesperado.";
      throw new ApiError(formattedMessage, status);
    }
    if (error.request) {
      throw new ApiError("Servidor não respondeu. Verifique sua conexão.", 0);
    }
  }
}
