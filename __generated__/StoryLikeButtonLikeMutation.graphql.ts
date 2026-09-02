/**
 * @generated SignedSource<<de4c54e2382b9a8ef1a28d542a306853>>
 * @relayHash c7cb88ef72137815b2e6992d1d20db47
 * @lightSyntaxTransform
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

// @relayRequestID 8005b72dd6e2315029055e6925130296581502fa35c7f2ee2825b3322e614231

import { ConcreteRequest } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type StoryLikeButtonLikeMutation$variables = {
  doesLike: boolean;
  id: string;
};
export type StoryLikeButtonLikeMutation$data = {
  readonly likeStory: {
    readonly story: {
      readonly __typename: "Story";
      readonly " $fragmentSpreads": FragmentRefs<"StoryLikeButtonFragment">;
    } | null | undefined;
  } | null | undefined;
};
export type StoryLikeButtonLikeMutation = {
  response: StoryLikeButtonLikeMutation$data;
  variables: StoryLikeButtonLikeMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "doesLike"
},
v1 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "id"
},
v2 = [
  {
    "kind": "Variable",
    "name": "doesLike",
    "variableName": "doesLike"
  },
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "id"
  }
],
v3 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "__typename",
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": [
      (v0/*:: as any*/),
      (v1/*:: as any*/)
    ],
    "kind": "Fragment",
    "metadata": null,
    "name": "StoryLikeButtonLikeMutation",
    "selections": [
      {
        "alias": null,
        "args": (v2/*:: as any*/),
        "concreteType": "StoryMutationResponse",
        "kind": "LinkedField",
        "name": "likeStory",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "Story",
            "kind": "LinkedField",
            "name": "story",
            "plural": false,
            "selections": [
              (v3/*:: as any*/),
              {
                "args": null,
                "kind": "FragmentSpread",
                "name": "StoryLikeButtonFragment"
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [
      (v1/*:: as any*/),
      (v0/*:: as any*/)
    ],
    "kind": "Operation",
    "name": "StoryLikeButtonLikeMutation",
    "selections": [
      {
        "alias": null,
        "args": (v2/*:: as any*/),
        "concreteType": "StoryMutationResponse",
        "kind": "LinkedField",
        "name": "likeStory",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "Story",
            "kind": "LinkedField",
            "name": "story",
            "plural": false,
            "selections": [
              (v3/*:: as any*/),
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
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ]
  },
  "params": {
    "id": "8005b72dd6e2315029055e6925130296581502fa35c7f2ee2825b3322e614231",
    "metadata": {},
    "name": "StoryLikeButtonLikeMutation",
    "operationKind": "mutation",
    "text": null
  }
};
})();

(node as any).hash = "d4742c7699165bfb2e69c31bd5646e3a";

export default node;
