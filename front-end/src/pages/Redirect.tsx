import { apiClient } from "@/api/client";
import { useEffect, useRef } from "react";
import { useNavigate, useParams } from "react-router";

export default function Redirect() {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const fetchedRef = useRef(false);

  useEffect(() => {
    if (fetchedRef.current) return;
    fetchedRef.current = true;

    if (!id) {
      navigate("/404", { replace: true });
      return;
    }

    const fetchURL = async () => {
      try {
        const response = await apiClient.url.redirect(id);

        if (response?.url) {
          window.location.href = response.url;
        } else {
          navigate("/404", { replace: true });
        }
      } catch (err) {
        navigate("/404", { replace: true });
      }
    };

    fetchURL();
  }, [id, navigate]);

  return (
    <div className="h-screen w-full flex flex-col justify-center items-center gap-12">
      <h1 className="text-7xl text-muted-foreground">Redirecionando...</h1>
      <p className="text-xl text-muted-foreground">
        Aguarde um momento, estamos te redirecionando...
      </p>
    </div>
  );
}
