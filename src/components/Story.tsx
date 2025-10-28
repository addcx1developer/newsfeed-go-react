import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { StoryFragment$key } from "../../__generated__/StoryFragment.graphql";

import Card from "./Card";
import PosterByline from "./PosterByline";
import Heading from "./Heading";
import Image from "./Image";
import StorySummary from "./StorySummary";
import Timestamp from "./Timestamp";
import StoryCommentsSection from "./StoryCommentsSection";

interface StoryProps {
  story: StoryFragment$key;
}

const StoryFragment = graphql`
  fragment StoryFragment on Story {
    title
    summary
    createdAt
    poster {
      ...PosterBylineFragment
    }
    thumbnail {
      ...ImageFragment @arguments(width: 400)
    }
    ...StoryCommentsSectionFragment
  }
`;

export default function Story({ story }: StoryProps): ReactElement {
  const data = useFragment<StoryFragment$key>(StoryFragment, story);

  return (
    <Card>
      <PosterByline poster={data!.poster} />
      <Heading>{data!.title}</Heading>
      <Timestamp time={data!.createdAt} />
      <Image image={data!.thumbnail!} width={400} height={400} />
      <StorySummary summary={data!.summary!} />
      <StoryCommentsSection story={data!} />
    </Card>
  );
}
