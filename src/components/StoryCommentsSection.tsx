import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { StoryCommentsSectionFragment$key } from "../../__generated__/StoryCommentsSectionFragment.graphql";

import Comment from "./Comment";

interface StoryCommentsSectionProps {
  story: StoryCommentsSectionFragment$key;
}

const StoryCommentsSectionFragment = graphql`
  fragment StoryCommentsSectionFragment on Story {
    comments(first: 3) {
      edges {
        node {
          id
          ...CommentFragment
        }
      }
      pageInfo {
        hasNextPage
      }
    }
  }
`;

export default function StoryCommentsSection({
  story,
}: StoryCommentsSectionProps): ReactElement {
  const data = useFragment<StoryCommentsSectionFragment$key>(
    StoryCommentsSectionFragment,
    story,
  );

  return (
    <div>
      {data!.comments!.edges!.map((edge) => (
        <Comment key={edge!.node!.id} comment={edge!.node!} />
      ))}
    </div>
  );
}
