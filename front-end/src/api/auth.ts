import { api } from "@/lib/api";
import { isAxiosError } from "axios";
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
} from "./types/auth";

export async function login(data: LoginRequest): Promise<LoginResponse> {
  try {
    const response = await api.post<LoginResponse>("/auth/login", data);
    return response.data;
  } catch (error: unknown) {
    if (isAxiosError(error)) {
      if (error.response) {
        const status = error.response.status;
        const message =
          (error.response.data as { message?: string })?.message ||
          `Erro ${status}: falha ao realizar login.`;
        throw new Error(message);
      }

      if (error.request) {
        throw new Error("Servidor não respondeu. Verifique sua conexão.");
      }
    }

    throw new Error(
      error instanceof Error
        ? `Erro inesperado: ${error.message}`
        : "Erro inesperado ao realizar login.",
    );
  }
}

export async function register(
  data: RegisterRequest,
): Promise<RegisterResponse> {
  try {
    const response = await api.post<RegisterResponse>("/auth/register", data);
    return response.data;
  } catch (error: unknown) {
    if (isAxiosError(error)) {
      if (error.response) {
        const status = error.response.status;
        // FastAPI geralmente retorna erros no campo 'detail'
        const message =
          (error.response.data as { detail?: string })?.detail ||
          `Erro ${status}: falha ao realizar cadastro.`;
        throw new Error(message);
      }

      if (error.request) {
        throw new Error("Servidor não respondeu. Verifique sua conexão.");
      }
    }

    throw new Error(
      error instanceof Error
        ? `Erro inesperado: ${error.message}`
        : "Erro inesperado ao realizar cadastro.",
    );
  }
}
