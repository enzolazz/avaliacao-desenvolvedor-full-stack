import { Skeleton } from "@/components/ui/skeleton";

export function DataTableSkeleton() {
  return (
    <div className="w-full h-full flex flex-col gap-6">
      <Skeleton className="w-full h-8" />
      <Skeleton className="w-full h-72" />
    </div>
  );
}
