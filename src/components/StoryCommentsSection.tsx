import { usePaginationFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { StoryCommentsSectionFragment$key } from "../../__generated__/StoryCommentsSectionFragment.graphql";
import type { StoryCommentsSectionPaginationQuery } from "../../__generated__/StoryCommentsSectionPaginationQuery.graphql";

import Comment from "./Comment";
import LoadMoreCommentsButton from "./LoadMoreCommentsButton";
import SmallSpinner from "./SmallSpinner";

interface StoryCommentsSectionProps {
  story: StoryCommentsSectionFragment$key;
}

const StoryCommentsSectionFragment = graphql`
  fragment StoryCommentsSectionFragment on Story
  @refetchable(queryName: "StoryCommentsSectionPaginationQuery")
  @argumentDefinitions(
    cursor: { type: "String" }
    count: { type: "Int", defaultValue: 3 }
  ) {
    comments(after: $cursor, first: $count)
      @connection(key: "StoryCommentsSectionFragment_comments") {
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
  const { data, loadNext, isLoadingNext } = usePaginationFragment<
    StoryCommentsSectionPaginationQuery,
    StoryCommentsSectionFragment$key
  >(StoryCommentsSectionFragment, story);

  const onLoadMore = () => loadNext(3);

  return (
    <div>
      {data!.comments!.edges!.map((edge) => (
        <Comment key={edge!.node!.id} comment={edge!.node!} />
      ))}
      {data!.comments!.pageInfo!.hasNextPage && (
        <LoadMoreCommentsButton onClick={onLoadMore} disabled={isLoadingNext} />
      )}
      {isLoadingNext && <SmallSpinner />}
    </div>
  );
}
