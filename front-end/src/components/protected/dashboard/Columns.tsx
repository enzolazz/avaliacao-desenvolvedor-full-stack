"use client";

import { CopyButton } from "@/components/ui/copy-button";
import type { ShortURL } from "@/types/url";
import {
  TooltipProvider,
  Tooltip,
  TooltipTrigger,
  TooltipContent,
} from "@/components/ui/tooltip";
import type { ColumnDef } from "@tanstack/react-table";
import { Link } from "react-router";
import { Badge } from "@/components/ui/badge";

function truncate(text: string, maxLength: number) {
  return text.length > maxLength ? text.slice(0, maxLength) + "…" : text;
}

export const columns: ColumnDef<ShortURL>[] = [
  {
    accessorKey: "label",
    header: "Apelido",

    cell: ({ row }) => (
      <div className="capitalize">
        {row.getValue<string>("label")
          ? truncate(row.getValue<string>("label"), 20)
          : "-"}
      </div>
    ),
  },
  {
    accessorKey: "id",
    header: "Identificador",
    cell: ({ row }) => {
      const short_url = row.getValue<string>("id");
      const rootURL = window.location.origin;

      return (
        <CopyButton
          text={`${rootURL}/${short_url}`}
          toastMessage="Link copiado com sucesso!"
        >
          {short_url}
        </CopyButton>
      );
    },
  },
  {
    accessorKey: "original_url",
    header: "URL Original",
    cell: ({ row }) => {
      const original_url = row.getValue<string>("original_url");

      return (
        <TooltipProvider delayDuration={100}>
          <Tooltip>
            <TooltipTrigger asChild>
              <Link
                to={original_url}
                onClick={(e) => e.stopPropagation()}
                className="text-primary"
              >
                {truncate(original_url, 30)}
              </Link>
            </TooltipTrigger>
            <TooltipContent>
              <p>{original_url}</p>
            </TooltipContent>
          </Tooltip>
        </TooltipProvider>
      );
    },
  },
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => (
      <div className="w-16 flex justify-center">
        {row.getValue("status") === "active" ? (
          <Badge
            variant="default"
            className="bg-green-500 h-5 w-5 rounded-full"
          ></Badge>
        ) : (
          <Badge variant="destructive" className="h-5 w-5 rounded-full"></Badge>
        )}
      </div>
    ),
  },
];
