import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { CommentFragment$key } from "../../__generated__/CommentFragment.graphql";

interface CommentProps {
  comment: CommentFragment$key;
}

const CommentFragment = graphql`
  fragment CommentFragment on Comment {
    text
  }
`;

export default function Comment({ comment }: CommentProps): ReactElement {
  const data = useFragment<CommentFragment$key>(CommentFragment, comment);

  return <div className="comment">{data.text}</div>;
}
