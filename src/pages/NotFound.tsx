import { Button } from "@/components/ui/button";
import { Link } from "react-router";

export default function NotFound() {
  return (
    <div className="h-screen w-full flex flex-col justify-center items-center gap-12">
      <h1 className="text-7xl text-muted-foreground">404</h1>
      <p className="text-xl text-muted-foreground">
        Ops! Página não encontrada!
      </p>

      <Button asChild>
        <Link to="/" className="text-xl">
          Retornar ao início
        </Link>
      </Button>
    </div>
  );
}
