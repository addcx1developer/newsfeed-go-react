import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { ViewerProfileFragment$key } from "../../__generated__/ViewerProfileFragment.graphql";

import Card from "./Card";

interface ViewerProfileProps {
  viewer: ViewerProfileFragment$key;
}

const ViewerProfileFragment = graphql`
  fragment ViewerProfileFragment on Viewer {
    actor {
      name
      profilePicture {
        url
      }
    }
  }
`;

export default function ViewerProfile({
  viewer,
}: ViewerProfileProps): ReactElement {
  const data = useFragment<ViewerProfileFragment$key>(
    ViewerProfileFragment,
    viewer,
  );

  return (
    <Card dim={true}>
      <div className="viewerProfile">
        <img src={data.actor!.profilePicture!.url} height="60" width="60" />
        <div className="viewerProfile__name">{data.actor!.name}</div>
        <div className="viewerProfile__menu">⋯</div>
      </div>
    </Card>
  );
}
