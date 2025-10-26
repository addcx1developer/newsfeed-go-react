import type { ReactElement, ReactNode } from "react";

interface CardProps {
  children: ReactNode;
  dim?: boolean;
}

export default function Card({ children, dim }: CardProps): ReactElement {
  return <div className={`card ${dim ? "dim" : ""}`}>{children}</div>;
}
