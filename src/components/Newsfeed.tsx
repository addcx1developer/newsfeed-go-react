import { graphql } from "relay-runtime";
import { useLazyLoadQuery } from "react-relay";

import type { ReactElement } from "react";
import type { NewsfeedQuery as NewsfeedQueryType } from "../../__generated__/NewsfeedQuery.graphql";

import Story from "./Story";

const NewsfeedQuery = graphql`
  query NewsfeedQuery {
    topStory {
      ...StoryFragment
    }
  }
`;

export default function Newsfeed(): ReactElement | null {
  const data = useLazyLoadQuery<NewsfeedQueryType>(NewsfeedQuery, {});
  const story = data.topStory;

  return (
    <div className="newsfeed">
      <Story story={story} />
    </div>
  );
}
