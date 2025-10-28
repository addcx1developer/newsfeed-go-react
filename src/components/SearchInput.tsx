import type { ReactElement } from "react";

import SmallSpinner from "./SmallSpinner";

interface SearchInputProps {
  value: string;
  onChange: (newValue: string) => void;
  isPending?: boolean;
}

export default function SearchInput({
  value,
  onChange,
  isPending,
}: SearchInputProps): ReactElement {
  return (
    <div className="searchInput">
      <input
        value={value}
        onChange={(e) => onChange(e.target.value)}
        placeholder="Search..."
      />
      {isPending && <SmallSpinner />}
    </div>
  );
}
