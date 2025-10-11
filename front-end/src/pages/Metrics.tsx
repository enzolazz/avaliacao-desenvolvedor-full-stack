import MetricsDashboard from "@/components/protected/metrics/MetricsDashboard";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Separator } from "@/components/ui/separator";
import { useData } from "@/hooks/use-data";
import type { ShortURL } from "@/types/url";
import { Link2 } from "lucide-react";
import { useState, useEffect } from "react";
import { Link, useNavigate, useParams } from "react-router";
import { toast } from "sonner";

export default function Metrics() {
  const { data, loading, error, refresh } = useData();
  const params = useParams();
  const navigate = useNavigate();

  const [currentId, setCurrentId] = useState<string | undefined>(params.id);
  const [selectedURL, setCurrentURL] = useState<ShortURL>();

  useEffect(() => {
    if (!data) return;

    console.log(data);

    if (params.id && data.length === 0) {
      navigate("/dashboard/metricas", { replace: true });
      return;
    }

    if (params.id) {
      const exists = data.find((d) => d.id === params.id);

      if (!exists) {
        toast.error("Você não tem permissão para acessar esta métrica.");
        navigate("/403", { replace: true });
        return;
      }

      setCurrentId(params.id);

      window.history.replaceState(null, "", "/dashboard/metricas");
    } else if (data.length > 0) {
      setCurrentId(data[0].id);
    }
  }, [params.id, data, navigate]);

  useEffect(() => {
    if (!data || !currentId) return;
    setCurrentURL(data.find((url) => url.id === currentId));
  }, [currentId, data]);

  const renderContent = () => {
    if (loading) {
      return <div className="text-muted-foreground">Carregando...</div>;
    }

    if (error) {
      return (
        <>
          <p className="text-destructive mb-4">Erro ao carregar suas URLs</p>
          <Button onClick={refresh}>Tentar novamente</Button>
        </>
      );
    }

    if (data.length === 0) {
      return (
        <div className="w-xl flex flex-col gap-4 justify-center mx-auto items-center bg-sidebar rounded-md p-4 border-sidebar-border">
          <div className="flex gap-2 items-center">
            <Link2 className="h-12 w-12 text-muted-foreground/50" />
            <h3 className="font-semibold text-lg">Nenhuma URL encontrada</h3>
          </div>
          <p className="text-muted-foreground text-sm max-w-sm text-center">
            Você ainda não encurtou nenhuma URL. Não é possível ver as métricas
            ainda.
          </p>
        </div>
      );
    }

    return (
      <div className="space-y-6">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center justify-between">
              Selecionar URL:
              <Select value={currentId} onValueChange={setCurrentId}>
                <SelectTrigger className="w-40">
                  <SelectValue placeholder="Selecionar URL" />
                </SelectTrigger>
                <SelectContent>
                  {data
                    .filter((item) => item.status === "active")
                    .map((item) => (
                      <SelectItem key={item.id} value={item.id}>
                        {item.id}
                      </SelectItem>
                    ))}
                </SelectContent>
              </Select>
            </CardTitle>
          </CardHeader>
          <CardContent className="space-y-4">
            <Separator />
            <div className="flex justify-between">
              <span className="font-medium text-foreground">Apelido:</span>
              <span className="text-accent-foreground">
                {selectedURL?.label || "-"}
              </span>
            </div>

            <div className="flex justify-between">
              <span className="font-medium text-foreground">
                Identificador:
              </span>
              <span className="text-accent-foreground">
                {selectedURL?.id || "-"}
              </span>
            </div>

            <div className="flex justify-between">
              <span className="font-medium text-foreground">URL Original:</span>
              <span className="text-accent-foreground max-w-1/2">
                <Link
                  to={selectedURL?.original_url || "/"}
                  className="text-primary wrap-anywhere hover:underline"
                >
                  {selectedURL?.original_url}
                </Link>
              </span>
            </div>
          </CardContent>
        </Card>
        <MetricsDashboard currentId={currentId} />
      </div>
    );
  };

  return (
    <div className="w-full max-w-screen-xl mx-auto space-y-8 p-4">
      {renderContent()}
    </div>
  );
}
