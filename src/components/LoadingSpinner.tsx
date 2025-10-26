import type { ReactElement } from "react";

export default function LoadingSpinner(): ReactElement {
  return (
    <div className="loading-spinner-wrapper">
      <div className="loading-spinner">
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
        <div></div>
      </div>
    </div>
  );
}
