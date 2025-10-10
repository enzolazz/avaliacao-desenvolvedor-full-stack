import { useState } from "react";
import { Card, CardHeader, CardTitle, CardContent } from "@/components/ui/card";
import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectItem,
} from "@/components/ui/select";
import { LastHourGraph, LastDayGraph, LastMonthGraph } from "./MetricsGraphs";
import type { MetricsQuery } from "@/types/data";

const MetricsDashboard = ({ currentId }: { currentId: string | undefined }) => {
  const [selected, setSelected] = useState<MetricsQuery>("last-hour");

  const renderGraph = () => {
    switch (selected) {
      case "last-hour":
        return <LastHourGraph shortLinkId={currentId} />;
      case "last-day":
        return <LastDayGraph shortLinkId={currentId} />;
      case "last-month":
        return <LastMonthGraph shortLinkId={currentId} />;
      default:
        return null;
    }
  };

  return (
    <Card className="w-full">
      <CardHeader className="flex flex-col sm:flex-row sm:justify-between sm:items-center gap-4">
        <CardTitle>Métricas da URL</CardTitle>
        <Select
          value={selected}
          onValueChange={(value) => setSelected(value as MetricsQuery)}
        >
          <SelectTrigger className="w-40">
            <SelectValue placeholder="Selecionar intervalo" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="last-hour">Última Hora</SelectItem>
            <SelectItem value="last-day">Último Dia</SelectItem>
            <SelectItem value="last-month">Último Mês</SelectItem>
          </SelectContent>
        </Select>
      </CardHeader>
      <CardContent>{renderGraph()}</CardContent>
    </Card>
  );
};

export default MetricsDashboard;
