import { graphql } from "relay-runtime";
import { useFragment, useMutation, ConnectionHandler } from "react-relay";

import { useState } from "react";

import type { StoryCommentsComposerFragment$key } from "../../__generated__/StoryCommentsComposerFragment.graphql";

interface StoryCommentsComposerProps {
  story: StoryCommentsComposerFragment$key;
}

const StoryCommentsComposerFragment = graphql`
  fragment StoryCommentsComposerFragment on Story {
    id
  }
`;

const StoryCommentsComposerPostMutation = graphql`
  mutation StoryCommentsComposerPostMutation(
    $id: ID!
    $text: String!
    $connections: [ID!]!
  ) {
    postStoryComment(id: $id, text: $text) {
      commentEdge @prependEdge(connections: $connections) {
        node {
          id
          text
        }
      }
    }
  }
`;

export default function StoryCommentsComposer({
  story,
}: StoryCommentsComposerProps) {
  const data = useFragment(StoryCommentsComposerFragment, story);
  const [text, setText] = useState("");
  const [commitMutation, isMutationInFlight] = useMutation(
    StoryCommentsComposerPostMutation,
  );

  const onPost = () => {
    setText("");
    const connectionID = ConnectionHandler.getConnectionID(
      data.id,
      "StoryCommentsSectionFragment_comments",
    );
    commitMutation({
      variables: {
        id: data.id,
        text,
        connections: [connectionID],
      },
    });
  };

  return (
    <div className="commentsComposer">
      <TextComposer
        text={text}
        onChange={setText}
        onReturn={onPost}
        disabled={isMutationInFlight}
      />
      <PostButton onClick={onPost} disabled={isMutationInFlight} />
    </div>
  );
}

interface TextComposerProps {
  text: string;
  onChange: (newValue: string) => void;
  onReturn: () => void;
  disabled?: boolean;
}

function TextComposer({
  text,
  onChange,
  onReturn,
  disabled,
}: TextComposerProps) {
  return (
    <input
      value={text}
      onChange={(e) => onChange(e.target.value)}
      onKeyDown={(e) => {
        if (e.key === "Enter" && !e.shiftKey) {
          onReturn();
        }
      }}
      disabled={disabled}
    />
  );
}

interface PostButtonProps {
  onClick: () => void;
  disabled?: boolean;
}

function PostButton({ onClick, disabled }: PostButtonProps) {
  return (
    <button onClick={onClick} disabled={disabled}>
      Post
    </button>
  );
}
