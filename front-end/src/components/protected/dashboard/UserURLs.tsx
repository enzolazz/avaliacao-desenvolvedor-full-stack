import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import { TableOfContents } from "lucide-react";
import { DataTable } from "./DataTable";
import type { ShortURL } from "@/types/url";
import { columns } from "./Columns";
import { useState, useEffect } from "react";

async function getData(): Promise<ShortURL[]> {
  return [
    {
      id: "m5gr84i9",
      label: "Google",
      status: "active",
      short_url: "https://meusite.com/abcd",
      original_url: "https://meusite.com/abcd",
    },
    {
      id: "m5gr84i9",
      label: "Google",
      status: "active",
      short_url: "https://meusite.com/abcd",
      original_url: "https://meusite.com/abcd",
    },
    {
      id: "m5gr84i9",
      label: "Google",
      status: "active",
      short_url: "https://meusite.com/abcd",
      original_url: "https://meusite.com/abcd",
    },
    {
      id: "m5gr84i9",
      label: "Google",
      status: "active",
      short_url: "https://meusite.com/abcd",
      original_url:
        "https://meusite.com/abcd999999999999999999999999999999999999999999999999999999999999999999999999",
    },
    {
      id: "m5gr84i9",
      label: "abacaxi",
      status: "active",
      short_url: "https://meusite.com/abcd",
      original_url: "https://meusite.com/abcd",
    },
    {
      id: "m5gr84i9",
      label: "Google",
      status: "active",
      short_url: "https://meusite.com/abcd",
      original_url: "https://meusite.com/abcd",
    },
    {
      id: "m5gr84i9",
      label: "Google",
      status: "active",
      short_url: "https://meusite.com/abcd",
      original_url: "https://meusite.com/abcd",
    },
  ];
}

export function UserURLs() {
  const [data, setData] = useState<ShortURL[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    async function fetchData() {
      const result: ShortURL[] = await getData();
      setData(result);
      setLoading(false);
    }
    fetchData();
  }, []);

  if (loading) return <div>Carregando...</div>;

  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className="flex gap-4 items-center text-xl">
          <TableOfContents /> Suas URLs encurtadas
        </CardTitle>
        <CardDescription>
          Gerencie todas as suas URLs encurtadas
        </CardDescription>
      </CardHeader>

      <CardContent className="w-full overflow-x-auto">
        <DataTable columns={columns} data={data} />
      </CardContent>
    </Card>
  );
}
