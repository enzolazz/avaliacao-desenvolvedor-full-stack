import { DataContext } from "@/types/data";
import { useContext } from "react";

export const useData = () => {
  const ctx = useContext(DataContext);
  if (!ctx)
    throw new Error("useShortURLs must be used inside a ShortURLProvider");
  return ctx;
};
