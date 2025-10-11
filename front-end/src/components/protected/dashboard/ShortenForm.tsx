"use client";

import { Button } from "@/components/ui/button";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

import { Input } from "@/components/ui/input";
import type { URLFormValues } from "@/types/url";
import type { UseFormReturn } from "react-hook-form";

interface URLFormProps {
  onSubmit: (values: URLFormValues) => void | Promise<void>;
  form: UseFormReturn<URLFormValues>;
}

export function ShortenForm({ onSubmit, form }: URLFormProps) {
  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4 w-full">
        <FormField
          control={form.control}
          name="label"
          render={({ field }) => (
            <FormItem>
              <FormLabel htmlFor="label">Apelido</FormLabel>
              <FormControl>
                <Input
                  id="label"
                  placeholder="URL 1"
                  {...field}
                  className="w-full max-w-full break-words"
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="url"
          render={({ field }) => (
            <FormItem>
              <FormLabel htmlFor="url">URL</FormLabel>
              <FormControl>
                <Input
                  id="url"
                  placeholder="https://exemplo.com/url-muito-longa"
                  {...field}
                  className="w-full max-w-full break-words"
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
          Encurtar
        </Button>
      </form>
    </Form>
  );
}
