import { graphql } from "relay-runtime";
import { useLazyLoadQuery } from "react-relay";

import type { ReactElement } from "react";
import type { NewsfeedQuery as NewsfeedQueryType } from "../../__generated__/NewsfeedQuery.graphql";

import Story from "./Story";

const NewsfeedQuery = graphql`
  query NewsfeedQuery {
    topStories {
      id
      ...StoryFragment
    }
  }
`;

export default function Newsfeed(): ReactElement | null {
  const data = useLazyLoadQuery<NewsfeedQueryType>(NewsfeedQuery, {});
  const stories = data.topStories;

  return (
    <div className="newsfeed">
      {stories?.map((story) => (
        <Story key={story?.id} story={story} />
      ))}
    </div>
  );
}
