import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { ContactRowFragment$key } from "../../__generated__/ContactRowFragment.graphql";

import Image from "./Image";

interface ContactRowProps {
  contact: ContactRowFragment$key;
}

const ContactRowFragment = graphql`
  fragment ContactRowFragment on Actor {
    name
    profilePicture {
      ...ImageFragment
    }
  }
`;

export default function ContactRow({ contact }: ContactRowProps): ReactElement {
  const data = useFragment<ContactRowFragment$key>(ContactRowFragment, contact);

  return (
    <div className="contactRow">
      <Image
        image={data!.profilePicture!}
        width={60}
        height={60}
        className="contactRow__image"
      />
      <div className="contactRow__name">{data.name}</div>
    </div>
  );
}
