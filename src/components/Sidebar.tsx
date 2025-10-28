import { Suspense } from "react";
import { useLazyLoadQuery } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { SidebarQuery as SidebarQueryType } from "../../__generated__/SidebarQuery.graphql";

import LoadingSpinner from "./LoadingSpinner";
import Card from "./Card";
import Image from "./Image";
import ViewerProfile from "./ViewerProfile";

const SidebarQuery = graphql`
  query SidebarQuery {
    viewer {
      ...ViewerProfileFragment
      contacts {
        id
        name
        profilePicture {
          ...ImageFragment
        }
      }
    }
  }
`;

export default function Sidebar(): ReactElement {
  return (
    <div className="sidebar">
      <Suspense fallback={<LoadingSpinner />}>
        <SidebarContents />
      </Suspense>
    </div>
  );
}

function SidebarContents(): ReactElement {
  const data = useLazyLoadQuery<SidebarQueryType>(SidebarQuery, {});

  return (
    <>
      <ViewerProfile viewer={data.viewer!} />
      <Card dim={true}>
        <h3>Contacts</h3>
        {data.viewer!.contacts!.map((contact) => (
          <div key={contact!.id} className="contactRow">
            <Image
              image={contact!.profilePicture}
              width={60}
              height={60}
              className="contactRow__image"
            />
            <div className="contactRow__name">{contact!.name}</div>
          </div>
        ))}
      </Card>
    </>
  );
}
