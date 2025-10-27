import type { ReactElement } from "react";

interface ImageProps {
  image?: {
    url: string;
  } | null;
  width?: number;
  height?: number;
  className?: string;
}

export default function Image({
  image,
  width,
  height,
  className,
}: ImageProps): ReactElement | null {
  if (!image) {
    return null;
  }

  return (
    <img
      key={image.url}
      src={image.url}
      width={width}
      height={height}
      className={className}
    />
  );
}
