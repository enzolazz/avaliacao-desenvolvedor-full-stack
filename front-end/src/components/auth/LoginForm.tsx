import { Button } from "@/components/ui/button";
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
  CardFooter,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@radix-ui/react-label";

export function LoginForm() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Realizar Login</CardTitle>
        <CardDescription>
          Digite suas credenciais para acessar sua conta!
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-6">
        <div className="grid gap-3">
          <Label htmlFor="username">Nome de usuário</Label>
          <Input id="username" type="text" placeholder="fulano01" />
        </div>
        <div className="grid gap-3">
          <Label htmlFor="password">Senha</Label>
          <Input id="password" type="password" placeholder="Uma senha forte" />
        </div>
      </CardContent>
      <CardFooter>
        <Button className="w-full">Login</Button>
      </CardFooter>
    </Card>
  );
}
