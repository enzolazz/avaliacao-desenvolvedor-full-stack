import type { User } from "@/types/user";

export interface RegisterRequest {
  name: string;
  surname: string;
  username: string;
  password: string;
}

export interface RegisterResponse {
  id: string;
  name: string;
  surname: string;
  username: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  user: User;
}
