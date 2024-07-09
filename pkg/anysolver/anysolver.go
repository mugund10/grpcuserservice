// a package to solve protobuf any to golang any(interface{}) problem vice versa.
package anysolver

import (
	"encoding/json"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

//definition is in the name
func ConvertInterfaceToAny(v interface{}) (*any.Any, error) {
	anyValue := &any.Any{}

	switch value := v.(type) {
	case bool:
		boolValue := &wrappers.BoolValue{Value: value}
		err := anypb.MarshalFrom(anyValue, boolValue, proto.MarshalOptions{})
		if err != nil {
			return nil, err
		}
	case string:
		stringValue := &wrappers.StringValue{Value: value}
		err := anypb.MarshalFrom(anyValue, stringValue, proto.MarshalOptions{})
		if err != nil {
			return nil, err
		}
	default:
		// For other types, marshal it to bytes and store in wrappers.BytesValue
		bytes, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}
		bytesValue := &wrappers.BytesValue{Value: bytes}
		err = anypb.MarshalFrom(anyValue, bytesValue, proto.MarshalOptions{})
		if err != nil {
			return nil, err
		}
	}

	return anyValue, nil
}


func ConvertAnyToInterface(anyValue *any.Any, soul int) (interface{}, error) {
	var value interface{}

	switch soul {
	case 1:	
				bytesValue := &wrappers.BoolValue{}
				err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
				if err != nil {
					return value, err
				}

				jsonBytes, jsonErr := json.Marshal(bytesValue.Value)
				if jsonErr != nil {
				return value, jsonErr
				}

				uErr := json.Unmarshal(jsonBytes, &value)
				if uErr != nil {
				return value, uErr
				}

				

	case 2:
				bytesValue := &wrappers.StringValue{}
				err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
				if err != nil {
					return value, err
				}
				jsonBytes, jsonErr := json.Marshal(bytesValue.Value)
				if jsonErr != nil {
				return value, jsonErr
				}

				uErr := json.Unmarshal(jsonBytes, &value)
				if uErr != nil {
				return value, uErr
				}

				
	default:
				var value interface{}
				bytesValue := &wrappers.BytesValue{}
				err := anypb.UnmarshalTo(anyValue, bytesValue, proto.UnmarshalOptions{})
				if err != nil {
					return value, err
				}
				uErr := json.Unmarshal(bytesValue.Value, &value)
				if uErr != nil {
					return value, uErr
				}
				return value, nil
				}
	return value, nil
}
