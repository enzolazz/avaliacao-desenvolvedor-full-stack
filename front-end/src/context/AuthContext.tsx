import { useState, type ReactNode } from "react";
import { apiClient } from "@/api/client";
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
} from "@/api/types/auth";
import { AuthContext } from "@/types/auth";
import type { User } from "@/types/user";

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(() => {
    const stored = localStorage.getItem("user");
    return stored ? JSON.parse(stored) : null;
  });
  const [token, setToken] = useState<string | null>(() =>
    localStorage.getItem("token"),
  );

  const isAuthenticated = !!token;

  const login = async (data: LoginRequest) => {
    const response: LoginResponse = await apiClient.auth.login(data);
    setUser(response.user);
    setToken(response.token);
    localStorage.setItem("user", JSON.stringify(response.user));
    localStorage.setItem("token", response.token);
  };

  const register = async (data: RegisterRequest) => {
    await apiClient.auth.register(data);
  };

  const logout = () => {
    setUser(null);
    setToken(null);
    localStorage.removeItem("user");
    localStorage.removeItem("token");
  };

  return (
    <AuthContext.Provider
      value={{ user, token, login, register, logout, isAuthenticated }}
    >
      {children}
    </AuthContext.Provider>
  );
}
