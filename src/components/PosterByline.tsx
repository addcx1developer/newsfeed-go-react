import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import Image from "./Image";

import type { PosterBylineFragment$key } from "../../__generated__/PosterBylineFragment.graphql";

interface PosterBylineProps {
  poster?: PosterBylineFragment$key;
}

const PosterBylineFragment = graphql`
  fragment PosterBylineFragment on Actor {
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

  if (!data) {
    return null;
  }

  return (
    <div className="byline">
      <Image
        image={data.profilePicture}
        width={60}
        height={60}
        className="byline__image"
      />
      <div className="byline__name">{data.name}</div>
    </div>
  );
}
