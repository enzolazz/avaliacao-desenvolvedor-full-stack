"use client";

import * as React from "react";
import { Button } from "@/components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { toast } from "sonner";

interface CopyButtonProps extends React.ComponentProps<"button"> {
  text: string;
  toastMessage: string;
  children?: React.ReactNode;
}

export function CopyButton({
  text,
  toastMessage,
  children,
  ...props
}: CopyButtonProps) {
  const handleCopy = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    try {
      await navigator.clipboard.writeText(text);

      toast(toastMessage);
    } catch (err) {
      console.error("Falha ao copiar:", err);
    }
  };

  return (
    <TooltipProvider delayDuration={100}>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button {...props} onClick={handleCopy} variant="ghost">
            {children}
          </Button>
        </TooltipTrigger>
        <TooltipContent>
          <p>Copiar</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
