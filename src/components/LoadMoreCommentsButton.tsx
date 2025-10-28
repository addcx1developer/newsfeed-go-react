import type { ReactElement } from "react";

interface LoadMoreCommentsButtonProps {
  onClick: () => void;
  disabled?: boolean;
}

export default function LoadMoreCommentsButton({
  onClick,
  disabled,
}: LoadMoreCommentsButtonProps): ReactElement {
  return (
    <button
      onClick={onClick}
      disabled={disabled}
      className="comments__loadMore"
    >
      Load more comments
    </button>
  );
}
