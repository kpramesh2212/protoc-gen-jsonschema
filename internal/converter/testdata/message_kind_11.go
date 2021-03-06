package testdata

const MessageKind11 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "properties": {
        "name": {
            "type": "string"
        },
        "ones": {
            "items": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "properties": {
                    "name": {
                        "type": "string"
                    },
                    "timestamp": {
                        "type": "string"
                    },
                    "id": {
                        "type": "integer"
                    },
                    "rating": {
                        "type": "number"
                    },
                    "complete": {
                        "type": "boolean"
                    }
                },
                "additionalProperties": true,
                "type": "object"
            },
            "type": "array"
        },
        "kind2": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "isa": {
                    "type": "boolean"
                },
                "hasa": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "kind3": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "someProp": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "kind4": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "special": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    },
    "additionalProperties": true,
    "type": "object"
}
`
