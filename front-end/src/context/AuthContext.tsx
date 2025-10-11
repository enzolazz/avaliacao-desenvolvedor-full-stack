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

  const isAuthenticated = !!user;

  const login = async (data: LoginRequest) => {
    const response: LoginResponse = await apiClient.auth.login(data);
    setUser(response.user);
    localStorage.setItem("user", JSON.stringify(response.user));
  };

  const register = async (data: RegisterRequest) => {
    await apiClient.auth.register(data);
  };

  const logout = async () => {
    await apiClient.auth.logout();
    setUser(null);
    localStorage.removeItem("user");
  };

  return (
    <AuthContext.Provider
      value={{ user, login, register, logout, isAuthenticated }}
    >
      {children}
    </AuthContext.Provider>
  );
}
