import { createContext } from "react";
import type { ShortURL } from "./url";

type DataContextType = {
  data: ShortURL[];
  loading: boolean;
  error: boolean;
  refresh: () => void;
};

export const DataContext = createContext<DataContextType | undefined>(
  undefined,
);

export type MetricsQuery = "last-hour" | "last-day" | "last-month";
