import { useRefetchableFragment } from "react-relay";
import { graphql } from "relay-runtime";

import { useState, type ReactElement } from "react";

import type { ContactsListFragment$key } from "../../__generated__/ContactsListFragment.graphql";
import type { ContactsListRefetchQuery } from "../../__generated__/ContactsListRefetchQuery.graphql";

import Card from "./Card";
import ContactRow from "./ContactRow";
import SearchInput from "./SearchInput";

interface ContactsListProps {
  viewer: ContactsListFragment$key;
}

const ContactsListFragment = graphql`
  fragment ContactsListFragment on Viewer
  @refetchable(queryName: "ContactsListRefetchQuery")
  @argumentDefinitions(search: { type: "String" }) {
    contacts(search: $search) {
      id
      ...ContactRowFragment
    }
  }
`;

export default function ContactsList({
  viewer,
}: ContactsListProps): ReactElement {
  const [data, refetch] = useRefetchableFragment<
    ContactsListRefetchQuery,
    ContactsListFragment$key
  >(ContactsListFragment, viewer);
  const [searchString, setSearchString] = useState("");

  const onSearchStringChanged = (value: string) => {
    setSearchString(value);
    refetch({ search: value });
  };

  return (
    <Card dim={true}>
      <h3>Contacts</h3>
      <SearchInput value={searchString} onChange={onSearchStringChanged} />
      {data.contacts!.map((contact) => (
        <ContactRow key={contact!.id} contact={contact!} />
      ))}
    </Card>
  );
}
