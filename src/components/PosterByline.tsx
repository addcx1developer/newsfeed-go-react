import type { ReactElement } from "react";

import Image from "./Image";

export interface PosterBylineProps {
  poster?: {
    name: string | null | undefined;
    profilePicture: { url: string } | null | undefined;
  };
}

export default function PosterByline({
  poster,
}: PosterBylineProps): ReactElement | null {
  if (!poster) {
    return null;
  }

  return (
    <div className="byline">
      <Image
        image={poster.profilePicture}
        width={60}
        height={60}
        className="byline__image"
      />
      <div className="byline__name">{poster.name}</div>
    </div>
  );
}
