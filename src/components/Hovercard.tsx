import { Suspense, useEffect, useState } from "react";
import { createPortal } from "react-dom";

import type { ReactElement, ReactNode, RefObject } from "react";

import LoadingSpinner from "./LoadingSpinner";

interface HovercardProps {
  children: ReactNode;
  targetRef: RefObject<HTMLDivElement | null>;
  onBeginHover?: () => void;
}

export default function Hovercard({
  children,
  targetRef,
  onBeginHover,
}: HovercardProps): ReactElement | null {
  const [hoverState, setHoverState] = useState<{ x: number; y: number } | null>(
    null,
  );

  useEffect(() => {
    const target = targetRef.current;

    if (!target) return;

    const enterCallback = (event: MouseEvent) => {
      onBeginHover?.();
      setHoverState({
        x: event.clientX,
        y: event.clientY,
      });
    };

    const leaveCallback = () => {
      setHoverState(null);
    };

    target.addEventListener("mouseenter", enterCallback);
    target.addEventListener("mouseleave", leaveCallback);

    return () => {
      target.removeEventListener("mouseenter", enterCallback);
      target.removeEventListener("mouseleave", leaveCallback);
    };
  });

  if (!hoverState) {
    return null;
  }

  return createPortal(
    <div
      className="hovercard"
      style={{
        top: hoverState.y + "px",
        left: hoverState.x + "px",
      }}
    >
      <Suspense fallback={<LoadingSpinner />}>{children}</Suspense>
    </div>,
    document.body,
  );
}
