import type { ReactElement } from "react";

import type { OrganizationKind } from "../../__generated__/PosterDetailsHovercardContentsBodyFragment.graphql";

interface OrganizationKindProps {
  kind: OrganizationKind;
}

export default function OrganizationKind({
  kind,
}: OrganizationKindProps): ReactElement {
  return <div className="byline__detail">{label(kind)}</div>;
}

function label(kind: OrganizationKind): string {
  switch (kind) {
    case "NONPROFIT":
      return "Non-Profit Organization";
  }
  return "";
}
