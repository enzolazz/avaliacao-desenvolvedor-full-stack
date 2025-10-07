"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import { Input } from "@/components/ui/input";
import { useNavigate } from "react-router";
import { toast } from "sonner";
import { register } from "@/api/auth";

const formSchema = z
  .object({
    name: z.string("Nome inválido").trim().min(1, "Nome é obrigatório"),
    surname: z.string("Nome inválido").trim().min(1, "Sobrenome é obrigatório"),
    username: z
      .string()
      .min(3, "Nome de usuário deve ter pelo menos 3 caracteres")
      .max(20, "Nome de usuário deve ter no máximo 20 caracteres")
      .regex(
        /^[a-z0-9_.-]{3,20}$/,
        "Nome de usuário só pode conter letras, números, underscore (_), ponto (.) e hífen (-)",
      ),
    password: z.string().min(5, "Senha deve ter pelo menos 8 caracteres"),
    confirmPassword: z.string("").min(5, "Confirmação de senha é obrigatória"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ["confirmPassword"],
    error: "As senhas não coincidem",
  });

type FormValues = z.infer<typeof formSchema>;

export function RegisterForm() {
  const navigate = useNavigate();

  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const handleRegister = async (values: FormValues) => {
    try {
      await register(values);
      navigate("/auth");
    } catch (error: unknown) {
      const message =
        error instanceof Error ? error.message : "Erro inesperado.";
      toast.error(message);
    }
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle>Realizar Login</CardTitle>
        <CardDescription>
          Digite suas credenciais para acessar sua conta!
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-6">
        <Form {...form}>
          <form
            onSubmit={form.handleSubmit(handleRegister)}
            className="space-y-8"
          >
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Nome</FormLabel>
                  <FormControl>
                    <Input placeholder="Fulano Henrique" {...field} />
                  </FormControl>

                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="surname"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Sobrenome</FormLabel>
                  <FormControl>
                    <Input placeholder="Costa da Silva" {...field} />
                  </FormControl>

                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Nome de usuário</FormLabel>
                  <FormControl>
                    <Input placeholder="fulano01" {...field} />
                  </FormControl>

                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Senha</FormLabel>
                  <FormControl>
                    <Input
                      placeholder="Uma senha forte"
                      {...field}
                      type="password"
                    />
                  </FormControl>

                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="confirmPassword"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Senha</FormLabel>
                  <FormControl>
                    <Input
                      placeholder="A mesma senha forte"
                      {...field}
                      type="password"
                    />
                  </FormControl>

                  <FormMessage />
                </FormItem>
              )}
            />

            <Button
              type="submit"
              className="w-full"
              disabled={form.formState.isSubmitting}
            >
              Cadastro
            </Button>
          </form>
        </Form>
      </CardContent>
    </Card>
  );
}
