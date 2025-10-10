import { Button } from "@/components/ui/button";
import { Link } from "react-router";

export default function Forbidden() {
  return (
    <div className="h-screen w-full flex flex-col justify-center items-center gap-12">
      <h1 className="text-7xl text-muted-foreground">403</h1>
      <p className="text-xl text-muted-foreground">
        Acesso negado! Você não tem permissões para acessar esta página.
      </p>

      <Button asChild>
        <Link to="/" className="text-xl">
          Retornar ao início
        </Link>
      </Button>
    </div>
  );
}
