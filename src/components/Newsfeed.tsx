import type { ReactElement } from "react";

export default function Newsfeed(): ReactElement {
  const story = {
    title: "Placeholder Story",
    summary: "Placeholder data, to be replaced with data fetched via GraphQL",
    poster: {
      name: "Placeholder Person",
      profilePicture: {
        url: "/assets/cat_avatar.png",
      },
    },
    thumbnail: {
      url: "/assets/placeholder.jpeg",
    },
  };

  return (
    <div className="newsfeed">
      <div className="card">
        <div className="byline">
          <img
            key={story.poster.profilePicture.url}
            src={story.poster.profilePicture.url}
            width={60}
            height={60}
            className="byline__image"
          />
          <div className="byline__name">{story.poster.name}</div>
        </div>
        <h2 className="heading">{story.title}</h2>
        <img
          key={story.thumbnail.url}
          src={story.thumbnail.url}
          width={400}
          height={400}
        />
        <div className="story__summary">
          <p>{story.summary}</p>
        </div>
      </div>
    </div>
  );
}
