import { ShortenCard } from "@/components/protected/dashboard/ShortenCard";
import { UserURLs } from "@/components/protected/dashboard/UserURLs";

export default function Dashboard() {
  return (
    <div className="w-full max-w-screen-xl mx-auto space-y-8 p-8">
      <ShortenCard />
      <UserURLs />
    </div>
  );
}
