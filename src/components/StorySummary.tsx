import type { ReactElement } from "react";

interface StorySummaryProps {
  summary: string | null | undefined;
}

export default function StorySummary({
  summary,
}: StorySummaryProps): ReactElement {
  return (
    <div className="story__summary">
      <p>{summary}</p>
    </div>
  );
}
