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
import { login } from "@/api/auth";

const formSchema = z.object({
  username: z
    .string("Nome de usuário inválido")
    .min(3, "Nome de usuário deve ter pelo menos 3 caracteres")
    .max(15, "Nome de usuário deve ter no máximo 15 caracteres")
    .regex(
      /^[a-z0-9_.-]{3,20}$/,
      "Nome de usuário só pode conter letras, números, underscore (_), ponto (.) e hífen (-)",
    ),
  password: z
    .string("Senha inválida")
    .min(5, "Senha deve ter pelo menos 5 caracteres"),
});

type FormValues = z.infer<typeof formSchema>;

export function LoginForm() {
  const navigate = useNavigate();

  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const handleLogin = async (values: FormValues) => {
    try {
      await login(values);
      navigate("/dashboard");
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
          <form onSubmit={form.handleSubmit(handleLogin)} className="space-y-8">
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
            <Button
              type="submit"
              className="w-full"
              disabled={form.formState.isSubmitting}
            >
              Login
            </Button>
          </form>
        </Form>
      </CardContent>
    </Card>
  );
}
