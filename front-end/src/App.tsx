import { ThemeProvider } from "@/components/ThemeProvider";
import { DashboardLayout } from "@/components/protected/DashboardLayout";
import { Toaster } from "@/components/ui/sonner";

import { ProtectedRoute } from "@/components/ProtectedRoute";
import { AuthProvider } from "@/context/AuthContext";
import Auth from "@/pages/Auth";
import Dashboard from "@/pages/Dashboard";
import NotFound from "@/pages/NotFound";
import Profile from "@/pages/Profile";
import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import Redirect from "@/pages/Redirect";
import Metrics from "@/pages/Metrics";
import Forbidden from "@/pages/Forbidden";

function App() {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <Toaster richColors />
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Navigate to="/dashboard" replace />} />
            <Route path="/auth" element={<Auth />} />
            <Route
              path="/dashboard"
              element={
                <ProtectedRoute>
                  <DashboardLayout />
                </ProtectedRoute>
              }
            >
              <Route index element={<Dashboard />} />
              <Route path="profile" element={<Profile />} />
              <Route path="metrics/:id?" element={<Metrics />} />
            </Route>
            <Route path="404" element={<NotFound />} />
            <Route path="403" element={<Forbidden />} />
            <Route path="/:id" element={<Redirect />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </ThemeProvider>
  );
}

export default App;
