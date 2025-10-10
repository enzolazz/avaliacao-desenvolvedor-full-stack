import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from "@/components/ui/card";
import { Link2, TableOfContents } from "lucide-react";
import { DataTable } from "./DataTable";
import { columns } from "./Columns";
import { Button } from "@/components/ui/button";
import { useData } from "@/hooks/use-data";

export function UserURLs() {
  const { data, loading, error, refresh } = useData();

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
        <>
          <Link2 className="h-12 w-12 text-muted-foreground/50 mb-4" />
          <h3 className="font-semibold text-lg mb-2">Nenhuma URL ainda</h3>
          <p className="text-muted-foreground text-sm max-w-sm">
            Você ainda não encurtou nenhuma URL. Use o formulário acima para
            criar sua primeira URL encurtada!
          </p>
        </>
      );
    }

    return <DataTable columns={columns} data={data} />;
  };

  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className="flex gap-4 items-center text-xl">
          <TableOfContents /> Suas URLs encurtadas
        </CardTitle>
        {!loading && !error && (
          <CardDescription>
            Gerencie todas as suas URLs encurtadas
          </CardDescription>
        )}
      </CardHeader>
      <CardContent
        className={`${data.length > 0 ? "w-full overflow-x-auto" : "flex flex-col items-center justify-center py-12 text-center"}`}
      >
        {renderContent()}
      </CardContent>
    </Card>
  );
}
