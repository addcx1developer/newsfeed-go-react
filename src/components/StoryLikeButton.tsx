import { useFragment, useMutation } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { StoryLikeButtonFragment$key } from "../../__generated__/StoryLikeButtonFragment.graphql";
import type { StoryLikeButtonLikeMutation } from "../../__generated__/StoryLikeButtonLikeMutation.graphql";

interface StoryLikeButtonProps {
  story: StoryLikeButtonFragment$key;
}

const StoryLikeButtonFragment = graphql`
  fragment StoryLikeButtonFragment on Story {
    id
    likeCount
    doesViewerLike
  }
`;

const StoryLikeButtonLikeMutation = graphql`
  mutation StoryLikeButtonLikeMutation($id: ID!, $doesLike: Boolean!) {
    likeStory(id: $id, doesLike: $doesLike) {
      story {
        ...StoryLikeButtonFragment
      }
    }
  }
`;

export default function StoryLikeButton({
  story,
}: StoryLikeButtonProps): ReactElement {
  const data = useFragment<StoryLikeButtonFragment$key>(
    StoryLikeButtonFragment,
    story,
  );
  const [commitMutation, isMutationInFlight] =
    useMutation<StoryLikeButtonLikeMutation>(StoryLikeButtonLikeMutation);

  const onLikeButtonClicked = () => {
    commitMutation({
      variables: {
        id: data.id,
        doesLike: !data.doesViewerLike,
      },
    });
  };

  return (
    <div className="likeButton">
      <LikeCount count={data.likeCount!} />
      <LikeButton
        doesViewerLike={data.doesViewerLike!}
        onClick={onLikeButtonClicked}
        disabled={isMutationInFlight}
      />
    </div>
  );
}

interface LikeCountProps {
  count: number;
}

function LikeCount({ count }: LikeCountProps): ReactElement {
  return <div className="likeButton__count">{count} likes</div>;
}

interface LikeButtonProps {
  doesViewerLike: boolean;
  onClick: () => void;
  disabled?: boolean;
}

function LikeButton({
  doesViewerLike,
  onClick,
  disabled,
}: LikeButtonProps): ReactElement {
  return (
    <button
      className="likeButton__button"
      onClick={onClick}
      disabled={disabled}
    >
      <span
        className={
          doesViewerLike
            ? "likeButton__thumb__viewerLikes"
            : "likeButton__thumb__viewerDoesNotLike"
        }
      >
        👍
      </span>{" "}
      <span
        className={
          doesViewerLike
            ? "likeButton__label__viewerLikes"
            : "likeButton__label__viewerDoesNotLike"
        }
      >
        Like
      </span>
    </button>
  );
}
