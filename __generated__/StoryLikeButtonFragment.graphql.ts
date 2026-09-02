/**
 * @generated SignedSource<<d1af2f564338fab6e526a90f97191353>>
 * @lightSyntaxTransform
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ReaderFragment } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type StoryLikeButtonFragment$data = {
  readonly $updatableFragmentSpreads: FragmentRefs<"StoryLikeButton_updatable">;
  readonly doesViewerLike: boolean | null | undefined;
  readonly id: string;
  readonly likeCount: number | null | undefined;
  readonly " $fragmentType": "StoryLikeButtonFragment";
};
export type StoryLikeButtonFragment$key = {
  readonly " $data"?: StoryLikeButtonFragment$data;
  readonly " $fragmentSpreads": FragmentRefs<"StoryLikeButtonFragment">;
};

const node: ReaderFragment = {
  "argumentDefinitions": [],
  "kind": "Fragment",
  "metadata": null,
  "name": "StoryLikeButtonFragment",
  "selections": [
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "id",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "likeCount",
      "storageKey": null
    },
    {
      "alias": null,
      "args": null,
      "kind": "ScalarField",
      "name": "doesViewerLike",
      "storageKey": null
    },
    {
      "args": null,
      "kind": "FragmentSpread",
      "name": "StoryLikeButton_updatable"
    }
  ],
  "type": "Story",
  "abstractKey": null
};

(node as any).hash = "06c200b881818c231520f1fa911785f2";

export default node;
