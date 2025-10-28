import { useFragment, useQueryLoader } from "react-relay";
import { graphql } from "relay-runtime";

import { useRef, type ReactElement } from "react";

import type { PosterBylineFragment$key } from "../../__generated__/PosterBylineFragment.graphql";
import type { PosterDetailsHovercardContentsQuery as HovercardQueryType } from "../../__generated__/PosterDetailsHovercardContentsQuery.graphql";

import Image from "./Image";
import Hovercard from "./Hovercard";
import PosterDetailsHovercardContents, {
  PosterDetailsHovercardContentsQuery,
} from "./PosterDetailsHovercardContents";

interface PosterBylineProps {
  poster: PosterBylineFragment$key;
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

  const [hovercardQueryRef, loadHovercardQuery] =
    useQueryLoader<HovercardQueryType>(PosterDetailsHovercardContentsQuery);

  if (!data) {
    return null;
  }

  function onBeginHover() {
    loadHovercardQuery({ posterID: data!.id });
  }

  return (
    <div ref={hoverRef} className="byline">
      <Image
        image={data!.profilePicture!}
        width={60}
        height={60}
        className="byline__image"
      />
      <div className="byline__name">{data.name}</div>
      <Hovercard onBeginHover={onBeginHover} targetRef={hoverRef}>
        <PosterDetailsHovercardContents queryRef={hovercardQueryRef!} />
      </Hovercard>
    </div>
  );
}
