import { useEffect, useState } from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import { Badge } from "@/components/ui/badge";
import dayjs from "dayjs";
import { toast } from "sonner";
import { apiClient } from "@/api/client";
import type { MetricData } from "@/api/types/metric";
import type { MetricsQuery } from "@/types/data";

type BaseMetricsGraphProps = {
  shortLinkId?: string;
  endpoint: MetricsQuery;
};

const BaseMetricsGraph = ({ shortLinkId, endpoint }: BaseMetricsGraphProps) => {
  const [data, setData] = useState<MetricData[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  useEffect(() => {
    if (!shortLinkId) return;

    const fetchData = async () => {
      setLoading(true);
      setError(false);
      try {
        const data = await apiClient.metrics.getMetrics(shortLinkId, endpoint);
        const formattedData = data.map((d) => ({
          ...d,
          displayTime:
            endpoint === "last-hour"
              ? dayjs(d.time).format("HH:mm")
              : endpoint === "last-day"
                ? dayjs(d.time).format("HH:mm")
                : dayjs(d.time).format("DD/MM"),
        }));
        setData(formattedData);
      } catch (err) {
        setError(true);
        toast.error("Não foi possível carregar métricas");
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [shortLinkId, endpoint]);

  if (loading) return <Badge variant="secondary">Carregando...</Badge>;
  if (error)
    return <Badge variant="destructive">Erro ao carregar métricas</Badge>;
  if (data.length === 0)
    return <Badge variant="outline">Nenhuma métrica</Badge>;

  return (
    <ResponsiveContainer width="100%" height={400}>
      <LineChart
        data={data}
        margin={{ top: 20, right: 40, bottom: 20, left: 0 }}
      >
        <CartesianGrid stroke="#e5e7eb" strokeDasharray="5 5" />

        <XAxis
          dataKey="displayTime"
          tick={{
            fill: "var(--color-foreground)",
            fontSize: 12,
            fontWeight: 500,
          }}
          tickLine={false}
          padding={{ left: 10, right: 10 }}
        />

        <YAxis
          allowDecimals={false}
          tick={{ fill: "var(--color-foreground)", fontSize: 12 }}
          tickLine={false}
          axisLine={{ stroke: "#d1d5db" }}
        />

        <Tooltip
          contentStyle={{
            backgroundColor: "var(--color-background)",
            borderRadius: 8,
            border: "none",
            padding: 10,
          }}
          labelStyle={{ color: "var(--color-foreground)", fontWeight: 600 }}
          formatter={(value: any) => [value, "Cliques"]}
        />

        <Line
          type="monotone"
          dataKey="count"
          stroke="var(--color-primary)"
          strokeWidth={3}
          dot={{
            r: 4,
            fill: "var(--color-primary)",
            stroke: "#fff",
            strokeWidth: 2,
          }}
          activeDot={{
            r: 6,
            fill: "var(--color-accent)",
            stroke: "#fff",
            strokeWidth: 2,
          }}
        />
      </LineChart>
    </ResponsiveContainer>
  );
};

export default BaseMetricsGraph;
