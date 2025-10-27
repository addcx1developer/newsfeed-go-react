import type { ReactElement } from "react";

interface StorySummaryProps {
  summary?: string | null;
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
