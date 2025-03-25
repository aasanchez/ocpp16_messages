/*
Package types provides core types used in the OCPP (Open Charge Point Protocol) message validation process.

This package contains essential data types, enums, and utility functions that are used throughout the OCPP
message processing, especially for message validation and handling various identifiers and tokens.

Key features include:

- **CiString**: A custom string type with specific validation rules, used for OCPP protocol message fields.
- **IdToken**: A type used for representing tokens in OCPP, typically used for authorization and identification.
- **Enums**: Provides common enumerated types used within OCPP messages (e.g., status enums, message types).
- **Utility Functions**: Provides functions for handling and validating these core types.

## Types

### CiString
CiString is a custom string type that is used in various OCPP messages to handle specific string formats
required by the protocol. It typically enforces length restrictions or specific character validation.

### IdToken
IdToken represents an identifier token used in OCPP messages, commonly used for identifying users,
charging stations, or other entities within the OCPP protocol. It is commonly used in the Authorize message type.

### Enums
This package contains a set of enumerated types used throughout the OCPP protocol, such as:
- **StatusEnum**: Represents various status values for charging stations or charging processes.
- **MessageTypeEnum**: Represents message types for OCPP communications.

## Functions

### NewCiString
NewCiString is a constructor function that ensures a string conforms to the CiString format.
It checks for length and format constraints and returns a validated CiString object.

### NewIdToken
NewIdToken creates and validates a new IdToken. It ensures that the token conforms to OCPP standards.

### ParseEnum
ParseEnum takes a string and attempts to parse it into the corresponding enum value, ensuring that it matches
one of the predefined enum values (e.g., for status or message type).
*/
package types
