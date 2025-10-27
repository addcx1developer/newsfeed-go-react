import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import { useRef, type ReactElement } from "react";

import type { PosterBylineFragment$key } from "../../__generated__/PosterBylineFragment.graphql";

import Image from "./Image";
import Hovercard from "./Hovercard";
import PosterDetailsHovercardContents from "./PosterDetailsHovercardContents";

interface PosterBylineProps {
  poster?: PosterBylineFragment$key;
}

const PosterBylineFragment = graphql`
  fragment PosterBylineFragment on Actor {
    id
    name
    profilePicture {
      ...ImageFragment @arguments(width: 60, height: 60)
    }
  }
`;

export default function PosterByline({
  poster,
}: PosterBylineProps): ReactElement | null {
  const data = useFragment<PosterBylineFragment$key>(
    PosterBylineFragment,
    poster,
  );
  const hoverRef = useRef<HTMLDivElement | null>(null);

  if (!data) {
    return null;
  }

  return (
    <div ref={hoverRef} className="byline">
      <Image
        image={data.profilePicture}
        width={60}
        height={60}
        className="byline__image"
      />
      <div className="byline__name">{data.name}</div>
      <Hovercard targetRef={hoverRef}>
        <PosterDetailsHovercardContents posterID={data.id} />
      </Hovercard>
    </div>
  );
}
