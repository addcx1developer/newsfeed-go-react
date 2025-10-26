import type { ReactElement, ReactNode } from "react";

interface HeadingProps {
  children: ReactNode;
}

export default function Heading({ children }: HeadingProps): ReactElement {
  return <h2 className="heading">{children}</h2>;
}
