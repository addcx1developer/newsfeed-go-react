import { useFragment } from "react-relay";
import { graphql } from "relay-runtime";

import type { ReactElement } from "react";

import type { ImageFragment$key } from "../../__generated__/ImageFragment.graphql";

interface ImageProps {
  image: ImageFragment$key;
  width?: number;
  height?: number;
  className?: string;
}

const ImageFragment = graphql`
  fragment ImageFragment on Image
  @argumentDefinitions(
    width: { type: "Int", defaultValue: null }
    height: { type: "Int", defaultValue: null }
  ) {
    url(width: $width, height: $height)
    altText
  }
`;

export default function Image({
  image,
  width,
  height,
  className,
}: ImageProps): ReactElement | null {
  const data = useFragment<ImageFragment$key>(ImageFragment, image);

  if (!data) {
    return null;
  }

  return (
    <img
      key={data.url}
      src={data.url}
      width={width}
      height={height}
      className={className}
      alt={data.altText ?? undefined}
    />
  );
}
