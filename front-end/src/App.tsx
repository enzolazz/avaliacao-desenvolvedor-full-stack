import { ThemeProvider } from "@/components/ThemeProvider";
import Auth from "@/pages/Auth";
import NotFound from "@/pages/NotFound";
import { BrowserRouter, Navigate, Route, Routes } from "react-router";

function App() {
  return (
    <ThemeProvider defaultTheme="light" storageKey="vite-ui-theme">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Navigate to="/dashboard" replace />} />
          <Route path="/auth" element={<Auth />} />
          {/* <Route */}
          {/*   path="/dashboard" */}
          {/*   element={ */}
          {/*     <ProtectedRoute> */}
          {/*       <DashboardLayout /> */}
          {/*     </ProtectedRoute> */}
          {/*   } */}
          {/* > */}
          {/*   <Route index element={<Dashboard />} /> */}
          {/*   <Route path="profile" element={<Profile />} /> */}
          {/* </Route> */}
          <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
