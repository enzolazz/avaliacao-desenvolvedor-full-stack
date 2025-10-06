import { AuthTabs } from "@/components/auth/AuthTabs";
import { ThemeToggle } from "@/components/ThemeToggle";

export default function Auth() {
  return (
    <div className="relative min-h-screen flex flex-col justify-center">
      <div className="absolute inset-x-0 p-6 top-1 w-fit left-1/2 -translate-x-1/2">
        <ThemeToggle iconSize="6" />
      </div>
      <main className="flex-1 flex justify-center items-center p-6 ">
        <AuthTabs />
      </main>
    </div>
  );
}
