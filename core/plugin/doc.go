/*
Package plugin provides functionality for managing and validating OCPP (Open Charge Point Protocol) messages.

This package is designed to handle the registration, validation, and management of message validators
for OCPP messages. It supports both JSON and SOAP message formats and allows users to hook into the
validation process through pre- and post-validation hooks.

Key features include:

  - **Validator Registration**: Allows users to register custom validators for specific actions in the OCPP protocol.
  - **Message Validation**: Provides a way to validate messages using the registered validators.
  - **Pre- and Post-Validation Hooks**: Lets users add logic to execute before or after validation of messages.
  - **Extensibility**: Through the use of the `MessageValidator` interface, developers can easily extend the package
    to support custom message validation logic.

This package provides a mock validator implementation for testing, allowing users to simulate both
successful and failed validation scenarios.

## Types

### MessageValidator
MessageValidator is an interface that all message validators must implement. It defines the contract for validating
messages in OCPP format (JSON or SOAP).

### mockValidator
mockValidator is a simple implementation of the MessageValidator interface used for testing purposes.
It allows simulating both successful and failed validation scenarios.

## Functions

### RegisterValidator
RegisterValidator registers a custom MessageValidator for a given action. This action will be validated using the
provided validator during the message validation process.

### GetValidator
GetValidator retrieves the MessageValidator associated with a specific action. This function returns the validator
and a boolean indicating if the validator was found.

### ValidateRawJSON
ValidateRawJSON validates a raw JSON payload for a specific action using the registered validator. It will return
the validation result or an error if the validation fails.

### SetPreValidationHook
SetPreValidationHook sets a function to be called before the validation of the message occurs.
This can be used for custom processing before validation.

### SetPostValidationHook
SetPostValidationHook sets a function to be called after the validation of the message is complete.
This can be used for custom processing after validation.
*/
package plugin
