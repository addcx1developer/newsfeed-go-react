import { useMemo } from "react";
import type { ReactElement, ReactNode } from "react";
import { RelayEnvironmentProvider } from "react-relay";

import { createEnvironment } from "./environment";

interface RelayEnvironmentProps {
  children: ReactNode;
}

export default function RelayEnvironment({
  children,
}: RelayEnvironmentProps): ReactElement {
  const environment = useMemo(() => {
    return createEnvironment();
  }, []);

  return (
    <RelayEnvironmentProvider environment={environment}>
      {children}
    </RelayEnvironmentProvider>
  );
}
