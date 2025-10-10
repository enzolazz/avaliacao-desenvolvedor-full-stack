import { ShortenCard } from "@/components/protected/dashboard/ShortenCard";
import { UserURLs } from "@/components/protected/dashboard/UserURLs";
import { useAuth } from "@/hooks/use-auth";

export default function Dashboard() {
  const { user } = useAuth();

  return (
    <div className="w-full max-w-screen-xl mx-auto space-y-8 p-4">
      <h1 className="text-3xl text-center">
        Bem-vindo(a), <strong className="text-primary">{user?.name}</strong>!
      </h1>
      <ShortenCard />
      <UserURLs />
    </div>
  );
}
