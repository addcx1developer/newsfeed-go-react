import { useFragment, useLazyLoadQuery } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import Image from "./Image";

import type {
  PosterDetailsHovercardContentsQuery$variables,
  PosterDetailsHovercardContentsQuery as PosterDetailsHovercardContentsQueryType,
} from "../../__generated__/PosterDetailsHovercardContentsQuery.graphql";
import type { PosterDetailsHovercardContentsBodyFragment$key } from "../../__generated__/PosterDetailsHovercardContentsBodyFragment.graphql";
import Timestamp from "./Timestamp";

interface PosterDetailsHovercardContentsProps {
  posterID: PosterDetailsHovercardContentsQuery$variables["posterID"];
}

const PosterDetailsHovercardContentsQuery = graphql`
  query PosterDetailsHovercardContentsQuery($posterID: ID!) {
    node(id: $posterID) {
      ... on Actor {
        ...PosterDetailsHovercardContentsBodyFragment
      }
    }
  }
`;

export default function PosterDetailsHovercardContents({
  posterID,
}: PosterDetailsHovercardContentsProps): ReactElement {
  const data = useLazyLoadQuery<PosterDetailsHovercardContentsQueryType>(
    PosterDetailsHovercardContentsQuery,
    { posterID },
  );

  return (
    <div className="posterHovercard">
      <PosterDetailsHovercardContentsBody poster={data.node} />
    </div>
  );
}

interface PosterDetailsHovercardContentsBodyProps {
  poster?: PosterDetailsHovercardContentsBodyFragment$key | null;
}

const PosterDetailsHovercardContentsBodyFragment = graphql`
  fragment PosterDetailsHovercardContentsBodyFragment on Actor {
    id
    name
    profilePicture {
      ...ImageFragment
    }
    joined
  }
`;

function PosterDetailsHovercardContentsBody({
  poster,
}: PosterDetailsHovercardContentsBodyProps): ReactElement {
  const data = useFragment<PosterDetailsHovercardContentsBodyFragment$key>(
    PosterDetailsHovercardContentsBodyFragment,
    poster,
  );

  return (
    <>
      <Image
        image={data?.profilePicture}
        width={128}
        height={128}
        className="posterHovercard__image"
      />
      <div className="posterHovercard__name">{data?.name}</div>
      <ul className="posterHovercard__details">
        <li>
          Joined <Timestamp time={data?.joined} />
        </li>
      </ul>
      <div className="posterHovercard__buttons">
        <button>Friend</button>
        <button>Message</button>
      </div>
    </>
  );
}
