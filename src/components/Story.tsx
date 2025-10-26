import type { ReactElement } from "react";

import type { PosterBylineProps } from "./PosterByline";

import Card from "./Card";
import PosterByline from "./PosterByline";
import Heading from "./Heading";
import Image from "./Image";
import StorySummary from "./StorySummary";

interface StoryProps {
  story: {
    title: string;
    summary: string | null | undefined;
    thumbnail: {
      url: string;
    } | null | undefined;
    poster: PosterBylineProps["poster"] | null | undefined;
  } | null | undefined;
}

export default function Story({ story }: StoryProps): ReactElement {
  return (
    <Card>
      <PosterByline poster={story?.poster} />
      <Heading>{story?.title}</Heading>
      <Image image={story?.thumbnail} width={400} height={400} />
      <StorySummary summary={story?.summary} />
    </Card>
  );
}
