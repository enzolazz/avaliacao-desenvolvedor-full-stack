import { api } from "@/lib/api";
import { isAxiosError } from "axios";
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
} from "./types/auth";

export class ApiError extends Error {
  status: number;
  constructor(message: string, status: number) {
    super(message);
    this.name = "ApiError";
    this.status = status;
  }
}

export async function login(data: LoginRequest): Promise<LoginResponse> {
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
}

export async function register(
  data: RegisterRequest,
): Promise<RegisterResponse> {
  try {
    const response = await api.post<RegisterResponse>("/auth/register", data);
    return response.data;
  } catch (error: unknown) {
    returnErrors(error);

    throw new ApiError(
      error instanceof Error
        ? `Erro inesperado: ${error.message}`
        : "Erro inesperado ao realizar cadastro.",
      0,
    );
  }
}

function returnErrors(error: unknown) {
  if (isAxiosError(error)) {
    if (error.response) {
      const status = error.response.status;
      const message = error.response.data.error;
      const formattedMessage =
        typeof message === "string" && message.length > 0
          ? `${message.charAt(0).toUpperCase()}${message.slice(1)}.`
          : "Ocorreu um erro inesperado.";
      throw new ApiError(
        formattedMessage,
        status,
      );
    }
    if (error.request) {
      throw new ApiError("Servidor não respondeu. Verifique sua conexão.", 0);
    }
  }
}
