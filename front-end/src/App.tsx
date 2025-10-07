import { ThemeProvider } from "@/components/ThemeProvider";
import { DashboardLayout } from "@/components/protected/DashboardLayout";
import { Toaster } from "@/components/ui/sonner";

import Auth from "@/pages/Auth";
import NotFound from "@/pages/NotFound";
import Dashboard from "@/pages/Dashboard";
import Profile from "@/pages/Profile";
import { BrowserRouter, Navigate, Route, Routes } from "react-router";

function App() {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <Toaster richColors />
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Navigate to="/dashboard" replace />} />
          <Route path="/auth" element={<Auth />} />
          <Route path="/dashboard" element={<DashboardLayout />}>
            {/* <Route */}
            {/*   path="/dashboard" */}
            {/*   element={ */}
            {/*     <ProtectedRoute> */}
            {/*       <DashboardLayout /> */}
            {/*     </ProtectedRoute> */}
            {/*   } */}
            <Route index element={<Dashboard />} />
            <Route path="profile" element={<Profile />} />
          </Route>
          <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
