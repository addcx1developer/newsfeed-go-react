import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { ContactsListFragment$key } from "../../__generated__/ContactsListFragment.graphql";

import Card from "./Card";
import Image from "./Image";

interface ContactsListProps {
  viewer: ContactsListFragment$key;
}

const ContactsListFragment = graphql`
  fragment ContactsListFragment on Viewer {
    contacts {
      id
      name
      profilePicture {
        ...ImageFragment
      }
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
  );
}
