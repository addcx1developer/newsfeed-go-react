/**
 * @generated SignedSource<<727188fdfb5c0c7d78cf606d707594c5>>
 * @relayHash 6e1c32cb099c36284aebd84f32273af5
 * @lightSyntaxTransform
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

// @relayRequestID 437964ab01e9ca657d9b701c25a7dd43ff6858272a0cca58b1a7a6ed966c2911

import { ConcreteRequest } from 'relay-runtime';
export type StoryCommentsComposerPostMutation$variables = {
  connections: ReadonlyArray<string>;
  id: string;
  text: string;
};
export type StoryCommentsComposerPostMutation$data = {
  readonly postStoryComment: {
    readonly commentEdge: {
      readonly node: {
        readonly id: string;
        readonly text: string | null | undefined;
      } | null | undefined;
    } | null | undefined;
  } | null | undefined;
};
export type StoryCommentsComposerPostMutation = {
  response: StoryCommentsComposerPostMutation$data;
  variables: StoryCommentsComposerPostMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "connections"
},
v1 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "id"
},
v2 = {
  "defaultValue": null,
  "kind": "LocalArgument",
  "name": "text"
},
v3 = [
  {
    "kind": "Variable",
    "name": "id",
    "variableName": "id"
  },
  {
    "kind": "Variable",
    "name": "text",
    "variableName": "text"
  }
],
v4 = {
  "alias": null,
  "args": null,
  "concreteType": "commentsEdge",
  "kind": "LinkedField",
  "name": "commentEdge",
  "plural": false,
  "selections": [
    {
      "alias": null,
      "args": null,
      "concreteType": "Comment",
      "kind": "LinkedField",
      "name": "node",
      "plural": false,
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
          "name": "text",
          "storageKey": null
        }
      ],
      "storageKey": null
    }
  ],
  "storageKey": null
};
return {
  "fragment": {
    "argumentDefinitions": [
      (v0/*:: as any*/),
      (v1/*:: as any*/),
      (v2/*:: as any*/)
    ],
    "kind": "Fragment",
    "metadata": null,
    "name": "StoryCommentsComposerPostMutation",
    "selections": [
      {
        "alias": null,
        "args": (v3/*:: as any*/),
        "concreteType": "StoryCommentMutationResponse",
        "kind": "LinkedField",
        "name": "postStoryComment",
        "plural": false,
        "selections": [
          (v4/*:: as any*/)
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
      (v2/*:: as any*/),
      (v0/*:: as any*/)
    ],
    "kind": "Operation",
    "name": "StoryCommentsComposerPostMutation",
    "selections": [
      {
        "alias": null,
        "args": (v3/*:: as any*/),
        "concreteType": "StoryCommentMutationResponse",
        "kind": "LinkedField",
        "name": "postStoryComment",
        "plural": false,
        "selections": [
          (v4/*:: as any*/),
          {
            "alias": null,
            "args": null,
            "filters": null,
            "handle": "prependEdge",
            "key": "",
            "kind": "LinkedHandle",
            "name": "commentEdge",
            "handleArgs": [
              {
                "kind": "Variable",
                "name": "connections",
                "variableName": "connections"
              }
            ]
          }
        ],
        "storageKey": null
      }
    ]
  },
  "params": {
    "id": "437964ab01e9ca657d9b701c25a7dd43ff6858272a0cca58b1a7a6ed966c2911",
    "metadata": {},
    "name": "StoryCommentsComposerPostMutation",
    "operationKind": "mutation",
    "text": null
  }
};
})();

(node as any).hash = "04dc56e9fef111ddfea1cf21a6be3442";

export default node;
