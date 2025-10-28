import { Suspense } from "react";
import { useLazyLoadQuery } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { SidebarQuery as SidebarQueryType } from "../../__generated__/SidebarQuery.graphql";

import LoadingSpinner from "./LoadingSpinner";
import ViewerProfile from "./ViewerProfile";
import ContactsList from "./ContactsList";

const SidebarQuery = graphql`
  query SidebarQuery {
    viewer {
      ...ViewerProfileFragment
      ...ContactsListFragment
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
      <ContactsList viewer={data.viewer!} />
    </>
  );
}
