// Code generated by thriftrw v1.4.0
// @generated

package exceptions

import "go.uber.org/thriftrw/thriftreflect"

var ThriftModule = &thriftreflect.ThriftModule{Name: "exceptions", Package: "go.uber.org/thriftrw/gen/testdata/exceptions", FilePath: "exceptions.thrift", SHA1: "7515e1a7bcd9ef547ad37a0eaaeaf89df21897fa", Raw: rawIDL}

const rawIDL = "exception EmptyException {}\n\nexception DoesNotExistException {\n    1: required string key\n    2: optional string Error (go.name=\"Error2\")\n}\n"
