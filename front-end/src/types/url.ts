import z from "zod";

export type ShortURL = {
  id: string;
  label: string;
  original_url: string;
  short_url: string;
  status: "active" | "disabled";
};

export const urlFormSchema = z.object({
  label: z
    .string("Apelido inválido")
    .max(15, "Apelido deve ter no máximo 15 caracteres"),
  url: z.string("URL inválida").min(1, "URL é obrigatório"),
});

export type URLFormValues = z.infer<typeof urlFormSchema>;
