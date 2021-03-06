package workerschema
//
// This file is automatically generated by schema/generator
//
// **** DO NOT EDIT ****
//
const DiscoveryJSON = `{
  "kind": "discovery#restDescription",
  "discoveryVersion": "v1",
  "id": "dex:v1",
  "name": "workerschema",
  "version": "v1",
  "title": "Dex API",
  "description": "The Dex REST API",
  "documentationLink": "http://github.com/coreos/dex",
  "protocol": "rest",
  "icons": {
    "x16": "",
    "x32": ""
  },
  "labels": [],
  "baseUrl": "$ENDPOINT/api/v1/",
  "basePath": "/api/v1/",
  "rootUrl": "$ENDPOINT/",
  "servicePath": "api/v1/",
  "batchPath": "batch",
  "parameters": {},
  "auth": {},
  "schemas": {
    "Error": {
      "id": "Error",
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "error_description": {
          "type": "string"
        }
      }
    },
    "Client": {
      "id": "Client",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "redirectURIs": {
          "required": true,
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "RefreshClient": {
      "id": "Client",
      "type": "object",
      "description": "A client with associated public metadata.",
      "properties": {
        "clientID": {
          "type": "string"
        },
        "clientName": {
          "type": "string"
        },
        "logoURI": {
          "type": "string"
        },
        "clientURI": {
          "type": "string"
        }
      }
    },
    "RefreshClientList": {
      "id": "RefreshClientList",
      "type": "object",
      "properties": {
        "clients": {
          "type": "array",
          "items": {
            "$ref": "RefreshClient"
          }
        }
      }
    },
    "ClientWithSecret": {
      "id": "Client",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "secret": {
          "type": "string"
        },
        "redirectURIs": {
          "required": true,
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "ClientPage": {
      "id": "ClientPage",
      "type": "object",
      "properties": {
        "clients": {
          "type": "array",
          "items": {
            "$ref": "Client"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    },
    "User": {
      "id": "User",
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "displayName": {
          "type": "string"
        },
        "emailVerified": {
          "type": "boolean"
        },
        "admin": {
          "type": "boolean"
        },
        "disabled": {
          "type": "boolean"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "UserResponse": {
      "id": "UserResponse",
      "type": "object",
      "properties": {
        "user": {
          "$ref": "User"
        }
      }
    },
    "UsersResponse": {
      "id": "UsersResponse",
      "type": "object",
      "properties": {
        "users": {
          "type": "array",
          "items": {
            "$ref": "User"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    },
    "UserCreateRequest": {
      "id": "UserCreateRequest",
      "type": "object",
      "properties": {
        "user": {
          "$ref": "User"
        },
        "redirectURL": {
          "type": "string",
          "format": "url"
        }
      }
    },
    "UserCreateResponse": {
      "id": "UserCreateResponse",
      "type": "object",
      "properties": {
        "user": {
          "type": "object",
          "$ref": "User"
        },
        "resetPasswordLink": {
          "type": "string"
        },
        "emailSent": {
          "type": "boolean"
        }
      }
    },
    "UserDisableRequest": {
      "id": "UserDisableRequest",
      "type": "object",
      "properties": {
        "disable": {
          "type": "boolean",
          "description": "If true, disable this user, if false, enable them. No error is signaled if the user state doesn't change."
        }
      }
    },
    "UserDisableResponse": {
      "id": "UserDisableResponse",
      "type": "object",
      "properties": {
        "ok": {
          "type": "boolean"
        }
      }
    },
    "ResendEmailInvitationRequest": {
      "id": "UserDisableRequest",
      "type": "object",
      "properties": {
        "redirectURL": {
          "type": "string",
          "format": "url"
        }
      }
    },
    "ResendEmailInvitationResponse": {
      "id": "UserDisableResponse",
      "type": "object",
      "properties": {
        "resetPasswordLink": {
          "type": "string"
        },
        "emailSent": {
          "type": "boolean"
        }
      }
    }
  },
  "resources": {
    "Clients": {
      "methods": {
        "List": {
          "id": "dex.Client.List",
          "description": "Retrieve a page of Client objects.",
          "httpMethod": "GET",
          "path": "clients",
          "parameters": {
            "nextPageToken": {
              "type": "string",
              "location": "query"
            }
          },
          "response": {
            "$ref": "ClientPage"
          }
        },
        "Create": {
          "id": "dex.Client.Create",
          "description": "Register a new Client.",
          "httpMethod": "POST",
          "path": "clients",
          "request": {
            "$ref": "Client"
          },
          "response": {
            "$ref": "ClientWithSecret"
          }
        }
      }
    },
    "Users": {
      "methods": {
        "List": {
          "id": "dex.User.List",
          "description": "Retrieve a page of User objects.",
          "httpMethod": "GET",
          "path": "users",
          "parameters": {
            "nextPageToken": {
              "type": "string",
              "location": "query"
            },
            "maxResults": {
              "type": "integer",
              "location": "query"
            }
          },
          "response": {
            "$ref": "UsersResponse"
          }
        },
        "Get": {
          "id": "dex.User.Get",
          "description": "Get a single User object by id.",
          "httpMethod": "GET",
          "path": "users/{id}",
          "parameters": {
            "id": {
              "type": "string",
              "required": true,
              "location": "path"
            }
          },
          "parameterOrder": [
            "id"
          ],
          "response": {
            "$ref": "UserResponse"
          }
        },
        "Create": {
          "id": "dex.User.Create",
          "description": "Create a new User.",
          "httpMethod": "POST",
          "path": "users",
          "request": {
            "$ref": "UserCreateRequest"
          },
          "response": {
            "$ref": "UserCreateResponse"
          }
        },
        "Disable": {
          "id": "dex.User.Disable",
          "description": "Enable or disable a user.",
          "httpMethod": "POST",
          "path": "users/{id}/disable",
          "parameters": {
            "id": {
              "type": "string",
              "required": true,
              "location": "path"
            }
          },
          "parameterOrder": [
            "id"
          ],
          "request": {
            "$ref": "UserDisableRequest"
          },
          "response": {
            "$ref": "UserDisableResponse"
          }
        },
        "ResendEmailInvitation": {
          "id": "dex.User.ResendEmailInvitation",
          "description": "Resend invitation email to an existing user with unverified email.",
          "httpMethod": "POST",
          "path": "users/{id}/resend-invitation",
          "parameters": {
            "id": {
              "type": "string",
              "required": true,
              "location": "path"
            }
          },
          "parameterOrder": [
            "id"
          ],
          "request": {
            "$ref": "ResendEmailInvitationRequest"
          },
          "response": {
            "$ref": "ResendEmailInvitationResponse"
          }
        }
      }
    },
    "RefreshClient": {
      "methods": {
        "List": {
          "id": "dex.RefreshClient.List",
          "description": "List all clients that hold refresh tokens for the specified user.",
          "httpMethod": "GET",
          "path": "account/{userid}/refresh",
          "parameters": {
            "userid": {
              "type": "string",
              "required": true,
              "location": "path"
            }
          },
          "parameterOrder": [
            "userid"
          ],
          "response": {
            "$ref": "RefreshClientList"
          }
        },
        "Revoke": {
          "id": "dex.RefreshClient.Revoke",
          "description": "Revoke all refresh tokens issues to the client for the specified user.",
          "httpMethod": "DELETE",
          "path": "account/{userid}/refresh/{clientid}",
          "parameterOrder": [
            "userid",
            "clientid"
          ],
          "parameters": {
            "clientid": {
              "type": "string",
              "required": true,
              "location": "path"
            },
            "userid": {
              "type": "string",
              "required": true,
              "location": "path"
            }
          }
        }
      }
    }
  }
}
`