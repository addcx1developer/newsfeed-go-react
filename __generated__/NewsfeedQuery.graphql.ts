/**
 * @generated SignedSource<<71e09a85c3f8fb3fe08ad36902e622a5>>
 * @relayHash cbf5ad8384a715c59b0d025c0482a203
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

// @relayRequestID 981920e7b62e02fc2915ca71ed17e1c9389d8c57e83795029ce20741547efdd9

import { ConcreteRequest } from 'relay-runtime';
import { FragmentRefs } from "relay-runtime";
export type NewsfeedQuery$variables = Record<PropertyKey, never>;
export type NewsfeedQuery$data = {
  readonly topStory: {
    readonly " $fragmentSpreads": FragmentRefs<"StoryFragment">;
  } | null | undefined;
};
export type NewsfeedQuery = {
  response: NewsfeedQuery$data;
  variables: NewsfeedQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "kind": "ScalarField",
    "name": "url",
    "storageKey": null
  },
  {
    "alias": null,
    "args": null,
    "kind": "ScalarField",
    "name": "altText",
    "storageKey": null
  }
],
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "NewsfeedQuery",
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "Story",
        "kind": "LinkedField",
        "name": "topStory",
        "plural": false,
        "selections": [
          {
            "args": null,
            "kind": "FragmentSpread",
            "name": "StoryFragment"
          }
        ],
        "storageKey": null
      }
    ],
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "NewsfeedQuery",
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "Story",
        "kind": "LinkedField",
        "name": "topStory",
        "plural": false,
        "selections": [
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "title",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "summary",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "createdAt",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "concreteType": null,
            "kind": "LinkedField",
            "name": "poster",
            "plural": false,
            "selections": [
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "__typename",
                "storageKey": null
              },
              {
                "kind": "TypeDiscriminator",
                "abstractKey": "__isActor"
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "name",
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "concreteType": "Image",
                "kind": "LinkedField",
                "name": "profilePicture",
                "plural": false,
                "selections": (v0/*: any*/),
                "storageKey": null
              },
              {
                "kind": "InlineFragment",
                "selections": [
                  (v1/*: any*/)
                ],
                "type": "Node",
                "abstractKey": "__isNode"
              }
            ],
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "concreteType": "Image",
            "kind": "LinkedField",
            "name": "thumbnail",
            "plural": false,
            "selections": (v0/*: any*/),
            "storageKey": null
          },
          (v1/*: any*/)
        ],
        "storageKey": null
      }
    ]
  },
  "params": {
    "id": "981920e7b62e02fc2915ca71ed17e1c9389d8c57e83795029ce20741547efdd9",
    "metadata": {},
    "name": "NewsfeedQuery",
    "operationKind": "query",
    "text": null
  }
};
})();

(node as any).hash = "65c357bf80ef742b634fe4d5946d1b3a";

export default node;
