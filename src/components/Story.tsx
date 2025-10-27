import type { ReactElement } from "react";

import type { PosterBylineProps } from "./PosterByline";

import Card from "./Card";
import PosterByline from "./PosterByline";
import Heading from "./Heading";
import Image from "./Image";
import StorySummary from "./StorySummary";
import Timestamp from "./Timestamp";

interface StoryProps {
  story?: {
    title: string;
    summary?: string | null;
    createdAt?: string;
    thumbnail?: {
      url: string;
    } | null;
    poster: PosterBylineProps["poster"];
  } | null;
}

export default function Story({ story }: StoryProps): ReactElement {
  return (
    <Card>
      <PosterByline poster={story?.poster} />
      <Heading>{story?.title}</Heading>
      <Timestamp time={story?.createdAt} />
      <Image image={story?.thumbnail} width={400} height={400} />
      <StorySummary summary={story?.summary} />
    </Card>
  );
}
