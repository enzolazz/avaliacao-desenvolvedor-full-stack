import BaseMetricsGraph from "./BaseMetricsGraph";

export function LastHourGraph({
  shortLinkId,
}: {
  shortLinkId: string | undefined;
}) {
  return <BaseMetricsGraph shortLinkId={shortLinkId} endpoint="last-hour" />;
}

export function LastDayGraph({
  shortLinkId,
}: {
  shortLinkId: string | undefined;
}) {
  return <BaseMetricsGraph shortLinkId={shortLinkId} endpoint="last-day" />;
}

export function LastMonthGraph({
  shortLinkId,
}: {
  shortLinkId: string | undefined;
}) {
  return <BaseMetricsGraph shortLinkId={shortLinkId} endpoint="last-month" />;
}
