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
  const wsRef = useRef<WebSocket | null>(null);

  const fetchData = async () => {
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
    setLoading(true);
    setError(false);
    if (fetchedRef.current) return;
    fetchedRef.current = true;

    const backend = import.meta.env.VITE_BACKEND_SERVER.replace(
      /^http?:\/\//,
      "",
    );

    const connect = () => {
      const wsProtocol = window.location.protocol === "https:" ? "wss" : "ws";
      const ws = new WebSocket(`${wsProtocol}://${backend}/updates/ws`);
      wsRef.current = ws;
      ws.onopen = () => console.log("WebSocket connected");
      ws.onmessage = () => fetchData();
      ws.onclose = () => {
        console.log("WebSocket disconnected. Reconnecting...");
        wsRef.current = null;
        setTimeout(() => {
          if (!wsRef.current) connect();
        }, 3000);
      };
    };

    connect();
    fetchData();
  }, []);

  return (
    <DataContext.Provider value={{ data, loading, error, refresh: fetchData }}>
      {children}
    </DataContext.Provider>
  );
}
