import { DataContext } from "@/types/data";
import { useContext } from "react";

export const useData = () => {
  const ctx = useContext(DataContext);
  if (!ctx)
    throw new Error("useData must be used inside a DataContext provider");
  return ctx;
};
