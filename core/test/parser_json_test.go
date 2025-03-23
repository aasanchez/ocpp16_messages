package core_test

import (
	"encoding/json"
	"testing"

	"github.com/aasanchez/ocpp16_messages/core"
)

func TestParseJSONMessage_ValidCALL(t *testing.T) {
	raw := []byte(`[2, "12345", "Authorize", {"idTag": "ABC123"}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALL {
		t.Errorf("expected TypeID CALL, got %v", msg.TypeID)
	}
	if msg.UniqueID != "12345" {
		t.Errorf("expected UniqueID '12345', got '%s'", msg.UniqueID)
	}
	if msg.Action != "Authorize" {
		t.Errorf("expected Action 'Authorize', got '%s'", msg.Action)
	}
	if !jsonEqual(msg.Payload, `{"idTag": "ABC123"}`) {
		t.Errorf("unexpected Payload: %s", string(msg.Payload))
	}
}

func TestParseJSONMessage_ValidCALLRESULT(t *testing.T) {
	raw := []byte(`[3, "67890", {"status": "Accepted"}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALLRESULT {
		t.Errorf("expected TypeID CALLRESULT, got %v", msg.TypeID)
	}
	if msg.UniqueID != "67890" {
		t.Errorf("expected UniqueID '67890', got '%s'", msg.UniqueID)
	}
	if !jsonEqual(msg.Payload, `{"status": "Accepted"}`) {
		t.Errorf("unexpected Payload: %s", string(msg.Payload))
	}
}

func TestParseJSONMessage_ValidCALLERROR(t *testing.T) {
	raw := []byte(`[4, "99999", "InternalError", "Something went wrong", {"reason": "crash"}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if msg.TypeID != core.CALLERROR {
		t.Errorf("expected TypeID CALLERROR, got %v", msg.TypeID)
	}
	if msg.UniqueID != "99999" {
		t.Errorf("expected UniqueID '99999', got '%s'", msg.UniqueID)
	}
	if msg.ErrorCode != "InternalError" {
		t.Errorf("expected ErrorCode 'InternalError', got '%s'", msg.ErrorCode)
	}
	if msg.ErrorDescription != "Something went wrong" {
		t.Errorf("expected ErrorDescription mismatch")
	}
	if !jsonEqual(msg.ErrorDetails, `{"reason": "crash"}`) {
		t.Errorf("unexpected ErrorDetails: %s", string(msg.ErrorDetails))
	}
}

func TestParseJSONMessage_InvalidJSON(t *testing.T) {
	raw := []byte(`not valid`)
	_, err := core.ParseJSONMessage(raw)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestParseJSONMessage_TooFewElements(t *testing.T) {
	raw := []byte(`[2]`)
	_, err := core.ParseJSONMessage(raw)
	if err == nil {
		t.Error("expected error for too few elements")
	}
}

func TestParseJSONMessage_InvalidTypeID(t *testing.T) {
	raw := []byte(`["abc", "123", "Action", {}]`)
	_, err := core.ParseJSONMessage(raw)
	if err == nil {
		t.Error("expected error for invalid message type ID")
	}
}

func TestParseJSONMessage_InvalidUniqueID(t *testing.T) {
	raw := []byte(`[2, {}, "Action", {}]`)
	_, err := core.ParseJSONMessage(raw)
	if err == nil {
		t.Error("expected error for invalid unique ID")
	}
}

func TestParseJSONMessage_CALL_InvalidAction(t *testing.T) {
	raw := []byte(`[2, "123", {}, {}]`)
	_, err := core.ParseJSONMessage(raw)
	if err == nil {
		t.Error("expected error for invalid action")
	}
}

func TestParseJSONMessage_CALL_MissingFields(t *testing.T) {
	raw := []byte(`[2, "123"]`)
	_, err := core.ParseJSONMessage(raw)
	if err == nil {
		t.Error("expected error for missing CALL fields")
	}
}

func TestParseJSONMessage_CALL_WrongNumberOfElements(t *testing.T) {
	raw := []byte(`[2, "123", "Authorize"]`) // Only 3 elements
	_, err := core.ParseJSONMessage(raw)
	if err == nil || err.Error() != "CALL message must have 4 elements" {
		t.Errorf("expected error for CALL message with wrong number of elements, got: %v", err)
	}
}

func TestParseJSONMessage_CALLRESULT_WrongNumberOfElements(t *testing.T) {
	raw := []byte(`[3, "123"]`) // Only 2 elements
	_, err := core.ParseJSONMessage(raw)
	if err == nil || err.Error() != "CALLRESULT message must have 3 elements" {
		t.Errorf("expected error for CALLRESULT message with wrong number of elements, got: %v", err)
	}
}

func TestParseJSONMessage_CALLERROR_TooShort(t *testing.T) {
	raw := []byte(`[4, "id", "code", "desc"]`) // Only 4 elements
	_, err := core.ParseJSONMessage(raw)
	if err == nil {
		t.Error("expected error for too short CALLERROR message")
	}
}

func TestParseJSONMessage_CALLERROR_InvalidCode(t *testing.T) {
	raw := []byte(`[4, "id", 123, "desc", {}]`) // errorCode is not string
	_, err := core.ParseJSONMessage(raw)
	if err == nil || err.Error() != "invalid errorCode" {
		t.Errorf("expected error for invalid errorCode, got: %v", err)
	}
}

func TestParseJSONMessage_CALLERROR_InvalidDescription(t *testing.T) {
	raw := []byte(`[4, "id", "code", 123, {}]`) // errorDescription is not string
	_, err := core.ParseJSONMessage(raw)
	if err == nil || err.Error() != "invalid errorDescription" {
		t.Errorf("expected error for invalid errorDescription, got: %v", err)
	}
}

func TestParseJSONMessage_UnsupportedMessageType(t *testing.T) {
	raw := []byte(`[99, "uid", "x", {}]`)
	msg, err := core.ParseJSONMessage(raw)
	if err == nil || err.Error() != "unsupported message type ID: 99" {
		t.Errorf("expected unsupported message type error, got: %v", err)
	}
	if msg != nil {
		t.Errorf("expected msg to be nil, got: %+v", msg)
	}
}

// jsonEqual compares two JSON-encoded values semantically.
func jsonEqual(a json.RawMessage, b string) bool {
	var objA, objB any
	_ = json.Unmarshal(a, &objA)
	_ = json.Unmarshal([]byte(b), &objB)
	return jsonMarshal(objA) == jsonMarshal(objB)
}

func jsonMarshal(v any) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}
