import type { LoginRequest, RegisterRequest } from "@/api/types/auth";
import { createContext } from "react";
import z from "zod";
import type { User } from "./user";

export const loginFormSchema = z.object({
  username: z
    .string("Nome de usuário inválido")
    .min(3, "Nome de usuário deve ter pelo menos 3 caracteres")
    .max(15, "Nome de usuário deve ter no máximo 15 caracteres")
    .regex(
      /^[a-z0-9_.-]{3,15}$/,
      "Nome de usuário só pode conter letras, números, underscore (_), ponto (.) e hífen (-)",
    ),
  password: z
    .string("Senha inválida")
    .min(5, "Senha deve ter pelo menos 5 caracteres"),
});

export const registerFormSchema = z
  .object({
    name: z.string("Nome inválido").trim().min(1, "Nome é obrigatório"),
    surname: z
      .string("Sobrenome inválido")
      .trim()
      .min(1, "Sobrenome é obrigatório"),
    username: z
      .string()
      .min(3, "Nome de usuário deve ter pelo menos 3 caracteres")
      .max(15, "Nome de usuário deve ter no máximo 15 caracteres")
      .regex(
        /^[a-z0-9_.-]{3,15}$/,
        "Nome de usuário só pode conter letras, números, underscore (_), ponto (.) e hífen (-)",
      ),
    password: z.string().min(5, "Senha deve ter pelo menos 5 caracteres"),
    confirmPassword: z.string("").min(5, "Confirmação de senha é obrigatória"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ["confirmPassword"],
    error: "As senhas não coincidem",
  });

export type LoginFormValues = z.infer<typeof loginFormSchema>;
export type RegisterFormValues = z.infer<typeof registerFormSchema>;

interface AuthContextType {
  user: User | null;
  token: string | null;
  login: (data: LoginRequest) => Promise<void>;
  register: (data: RegisterRequest) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);
