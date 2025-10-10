import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { urlFormSchema, type URLFormValues } from "@/types/url";
import { zodResolver } from "@hookform/resolvers/zod";
import { Link2 } from "lucide-react";
import { type UseFormReturn, useForm } from "react-hook-form";
import { ShortenForm } from "./ShortenForm";
import { apiClient, ApiError } from "@/api/client";
import { toast } from "sonner";
import { useData } from "@/hooks/use-data";

export function ShortenCard() {
  const { refresh } = useData();

  const form: UseFormReturn<URLFormValues> = useForm({
    resolver: zodResolver(urlFormSchema),
    defaultValues: { label: "", url: "" },
  });

  const onSubmit = async (values: URLFormValues) => {
    try {
      await apiClient.url.shorten(values);
      toast.success("URL encurtada!");

      form.reset();
      refresh();
    } catch (err: unknown) {
      if (err instanceof ApiError) {
        form.setError("url", {
          type: "manual",
          message: err.message,
        });
      } else {
        form.setError("root", {
          type: "manual",
          message: "Erro inesperado ao encurtar URL",
        });
      }
    }
  };

  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className="flex gap-4 items-center text-xl">
          <Link2 /> Encurtar nova URL
        </CardTitle>
        <CardDescription>
          Cole a URL que deseja encurtar e clique em "Encurtar"
        </CardDescription>
      </CardHeader>

      <CardContent className="w-full overflow-x-auto">
        <ShortenForm onSubmit={onSubmit} form={form} />
      </CardContent>
    </Card>
  );
}
