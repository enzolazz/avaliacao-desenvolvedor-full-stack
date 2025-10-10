import { apiClient, ApiError } from "@/api/client";
import { DataContext } from "@/types/data";
import type { ShortURL } from "@/types/url";
import { useEffect, useRef, useState, type ReactNode } from "react";
import { toast } from "sonner";

export function DataProvider({ children }: { children: ReactNode }) {
  const [data, setData] = useState<ShortURL[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);
  const fetchedRef = useRef(false);

  const fetchData = async () => {
    setLoading(true);
    setError(false);
    try {
      const result = await apiClient.url.getAllLinks();
      setData(result || []);
    } catch (err: unknown) {
      setError(true);
      toast.error(
        err instanceof ApiError
          ? err.message
          : "Erro inesperado ao carregar URLs",
      );
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (fetchedRef.current) return;
    fetchedRef.current = true;

    fetchData();
  }, []);

  return (
    <DataContext.Provider value={{ data, loading, error, refresh: fetchData }}>
      {children}
    </DataContext.Provider>
  );
}
