import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarProvider,
} from "@/components/ui/sidebar";
import { Header } from "./Header";
import { Link, Outlet, useLocation } from "react-router";
import { Link2, User } from "lucide-react";
import { ThemeToggle } from "../ThemeToggle";

const items = [
  {
    title: "URLs",
    url: "/dashboard",
    icon: Link2,
  },
  {
    title: "Perfil",
    url: "/dashboard/profile",
    icon: User,
  },
];

function AppSidebar() {
  return (
    <Sidebar>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" className="flex flex-row-reverse">
              <ThemeToggle iconSize="2" />
              <div className="flex-1">
                <Link to="/dashboard" className="flex gap-2">
                  <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary">
                    <Link2 className="h-5 w-5 text-primary-foreground" />
                  </div>
                  <div className="flex flex-col gap-0.5 leading-none">
                    <span className="font-semibold">Encurtador</span>
                    <span className="text-xs text-muted-foreground">
                      Dashboard
                    </span>
                  </div>
                </Link>
              </div>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Navegação</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <Link to={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}

export function DashboardLayout() {
  const location = useLocation();
  const title =
    items.find((item) => item.url === location.pathname)?.title || "Dashboard";
  return (
    <SidebarProvider>
      <AppSidebar />
      <div className="flex flex-col w-full h-screen overflow-hidden">
        <header className="flex-shrink-0">
          <Header title={title} />
        </header>
        <main className="flex-1 overflow-y-auto">
          <Outlet />
        </main>
      </div>
    </SidebarProvider>
  );
}
