import { Separator } from "@/components/ui/separator";
import { Skeleton } from "@/components/ui/skeleton";

export function MetricsSkeleton() {
  return (
    <div className="flex flex-col gap-6">
      <div className="space-y-6">
        <Skeleton className="w-full bg-sidebar flex flex-col p-8 justify-between">
          <div className="w-full flex justify-between items-center gap-4">
            <Skeleton className="w-32 h-6 bg-muted-foreground/20" />
            <Skeleton className="w-44 h-8 bg-muted-foreground/20" />
          </div>
          <Separator className="my-8" />
          <div className="flex flex-col gap-6">
            <div className="w-full flex justify-between items-center gap-4">
              <Skeleton className="w-32 h-4 bg-muted-foreground/20" />
              <Skeleton className="w-44 h-4 bg-muted-foreground/20" />
            </div>
            <div className="w-full flex justify-between items-center gap-4">
              <Skeleton className="w-40 h-4 bg-muted-foreground/20" />
              <Skeleton className="w-44 h-4 bg-muted-foreground/20" />
            </div>
            <div className="w-full flex justify-between items-center gap-4">
              <Skeleton className="w-44 h-4 bg-muted-foreground/20" />
              <Skeleton className="w-56 h-4 bg-muted-foreground/20" />
            </div>
          </div>
        </Skeleton>
      </div>
      <div className="space-y-6">
        <Skeleton className="w-full bg-sidebar flex flex-col p-8 justify-between">
          <div className="w-full flex justify-between items-center gap-4">
            <Skeleton className="w-32 h-6 bg-muted-foreground/20" />
            <Skeleton className="w-44 h-8 bg-muted-foreground/20" />
          </div>
          <Skeleton className="h-80 w-full mt-6" />
        </Skeleton>
      </div>
    </div>
  );
}
