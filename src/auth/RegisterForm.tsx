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

export function RegisterForm() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Cadastro</CardTitle>
        <CardDescription>
          Crie uma conta para começar a encurtar suas URLs!
        </CardDescription>
      </CardHeader>
      <CardContent className="grid gap-6">
        <div className="grid gap-3">
          <Label htmlFor="name">Nome</Label>
        </div>
        <Input id="name" type="text" placeholder="Fulano Henrique" />
        <div className="grid gap-3">
          <Label htmlFor="surname">Sobrenome</Label>
          <Input id="surname" type="text" placeholder="Costa da Silva" />
        </div>
        <div className="grid gap-3">
          <Label htmlFor="username">Nome de usuário</Label>
          <Input id="username" type="text" placeholder="fulano01" />
        </div>
        <div className="grid gap-3">
          <Label htmlFor="password">Senha</Label>
          <Input id="password" type="password" placeholder="Uma senha forte" />
        </div>
        <div className="grid gap-3">
          <Label htmlFor="password">Confimar senha</Label>
          <Input
            id="password"
            type="password"
            placeholder="A mesma senha forte"
          />
        </div>
      </CardContent>
      <CardFooter>
        <Button className="w-full">Cadastrar</Button>
      </CardFooter>
    </Card>
  );
}
