import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { ContactsListFragment$key } from "../../__generated__/ContactsListFragment.graphql";

import Card from "./Card";
import ContactRow from "./ContactRow";

interface ContactsListProps {
  viewer: ContactsListFragment$key;
}

const ContactsListFragment = graphql`
  fragment ContactsListFragment on Viewer {
    contacts {
      id
      ...ContactRowFragment
    }
  }
`;

export default function ContactsList({
  viewer,
}: ContactsListProps): ReactElement {
  const data = useFragment<ContactsListFragment$key>(
    ContactsListFragment,
    viewer,
  );

  return (
    <Card dim={true}>
      <h3>Contacts</h3>
      {data.contacts!.map((contact) => (
        <ContactRow key={contact!.id} contact={contact!} />
      ))}
    </Card>
  );
}
