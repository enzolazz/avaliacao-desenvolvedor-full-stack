import { useState } from "react";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";
import { toast } from "sonner";
import { useNavigate } from "react-router";
import { ApiError } from "@/api/client";
import { LoginForm } from "./LoginForm";
import { RegisterForm } from "./RegisterForm";
import { registerFormSchema, loginFormSchema } from "@/types/auth";
import type { RegisterFormValues, LoginFormValues } from "@/types/auth";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm, type UseFormReturn } from "react-hook-form";
import { useAuth } from "@/hooks/use-auth";

export function AuthTabs() {
  const navigate = useNavigate();
  const { login, register } = useAuth();

  const [activeTab, setActiveTab] = useState<"login" | "register">("login");

  const registerForm: UseFormReturn<RegisterFormValues> = useForm({
    resolver: zodResolver(registerFormSchema),
    defaultValues: {
      name: "",
      surname: "",
      username: "",
      password: "",
      confirmPassword: "",
    },
  });

  const loginForm: UseFormReturn<LoginFormValues> = useForm({
    resolver: zodResolver(loginFormSchema),
    defaultValues: { username: "", password: "" },
  });

  const handleRegister = async (values: RegisterFormValues) => {
    try {
      await register(values);
      toast.success("Registro realizado com sucesso!");

      registerForm.reset();
      navigate("/dashboard");
    } catch (err: unknown) {
      if (err instanceof ApiError) {
        if (err.status === 409) {
          registerForm.setError("username", {
            type: "manual",
            message: err.message,
          });
        } else {
          registerForm.setError("root", {
            type: "manual",
            message: err.message,
          });
        }
      } else {
        registerForm.setError("root", {
          type: "manual",
          message: "Erro inesperado",
        });
      }
    }
  };

  const handleLogin = async (values: LoginFormValues) => {
    try {
      await login(values);
      toast.success("Login realizado!");

      loginForm.reset();
      navigate("/dashboard");
    } catch (err: unknown) {
      if (err instanceof ApiError) {
        loginForm.setError("username", {
          type: "manual",
          message: err.message,
        });

        loginForm.setError("password", {
          type: "manual",
          message: "",
        });
      } else {
        loginForm.setError("root", {
          type: "manual",
          message: "Erro inesperado ao realizar login",
        });
      }
    }
  };
  return (
    <div className="flex w-full max-w-xl flex-col gap-6 text-xl">
      <Tabs
        value={activeTab}
        onValueChange={(v) => setActiveTab(v as "login" | "register")}
      >
        <TabsList>
          <TabsTrigger value="login">Login</TabsTrigger>
          <TabsTrigger value="register">Registrar</TabsTrigger>
        </TabsList>

        <TabsContent value="login">
          <LoginForm onSubmit={handleLogin} form={loginForm} />
        </TabsContent>

        <TabsContent value="register">
          <RegisterForm onSubmit={handleRegister} form={registerForm} />
        </TabsContent>
      </Tabs>
    </div>
  );
}
