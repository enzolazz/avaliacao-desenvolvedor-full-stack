import { useEffect, useRef, useState } from "react";

export function useOnceAsync<T>(asyncFn: () => Promise<T>) {
  const hasRun = useRef(false);
  const [data, setData] = useState<T | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<unknown>(null);

  useEffect(() => {
    if (!hasRun.current) {
      hasRun.current = true;

      const run = async () => {
        setLoading(true);
        setError(null);
        try {
          const result = await asyncFn();
          setData(result);
        } catch (err) {
          setError(err);
        } finally {
          setLoading(false);
        }
      };

      run();
    }
  }, [asyncFn]);

  return { data, loading, error };
}
