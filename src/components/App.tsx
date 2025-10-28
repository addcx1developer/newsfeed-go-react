import { Suspense } from "react";

import type { ReactElement } from "react";

import RelayEnvironment from "../relay/RelayEnvironment";
import LoadingSpinner from "./LoadingSpinner";
import Newsfeed from "./Newsfeed";
import Sidebar from "./Sidebar";

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
        <div className="app">
          <Newsfeed />
          <Sidebar />
        </div>
      </Suspense>
    </RelayEnvironment>
  );
}
