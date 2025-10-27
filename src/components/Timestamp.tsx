import type { ReactElement } from "react";

interface TimestampProps {
  time?: string;
}

export default function Timestamp({
  time,
}: TimestampProps): ReactElement | null {
  if (time == null) {
    return null;
  }

  return <span className="timestamp">{new Date(time).toDateString()}</span>;
}
