"use client";

import type { ShortURL } from "@/types/url";
import type { ColumnDef } from "@tanstack/react-table";
import { CircleCheck, CircleX } from "lucide-react";
import { Link } from "react-router";

function truncate(text: string, maxLength: number) {
  return text.length > maxLength ? text.slice(0, maxLength) + "…" : text;
}

export const columns: ColumnDef<ShortURL>[] = [
  {
    accessorKey: "label",
    header: "Nome",

    cell: ({ row }) => (
      <div className="capitalize">
        {truncate(row.getValue<string>("label"), 20)}
      </div>
    ),
  },
  {
    accessorKey: "short_url",
    header: "URL Encurtada",
    cell: ({ row }) => {
      const short_url = row.getValue<string>("short_url");

      return (
        <Link
          to={short_url}
          onClick={(e) => e.stopPropagation()}
          className="text-primary"
        >
          {short_url}
        </Link>
      );
    },
  },
  {
    accessorKey: "original_url",
    header: "URL Original",
    cell: ({ row }) => {
      const original_url = row.getValue<string>("original_url");

      return (
        <Link
          to={original_url}
          onClick={(e) => e.stopPropagation()}
          className="text-primary"
        >
          {truncate(original_url, 30)}
        </Link>
      );
    },
  },
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => (
      <div>
        {row.getValue("status") === "active" ? (
          <CircleCheck color="green" />
        ) : (
          <CircleX color="red" />
        )}
      </div>
    ),
  },
];
