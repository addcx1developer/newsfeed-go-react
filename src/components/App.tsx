import { Suspense } from "react";
import type { ReactElement } from "react";
import RelayEnvironment from "../relay/RelayEnvironment";
import LoadingSpinner from "./LoadingSpinner";

export default function App(): ReactElement {
  return (
    <RelayEnvironment>
      <Suspense
        fallback={
          <div className="app-loading-spinner">
            <LoadingSpinner />
          </div>
        }
      >
        <div className="app"></div>
      </Suspense>
    </RelayEnvironment>
  );
}
