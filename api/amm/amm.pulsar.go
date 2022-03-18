// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package ammv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Pool                     protoreflect.MessageDescriptor
	fd_Pool_pair                protoreflect.FieldDescriptor
	fd_Pool_quote_asset_reserve protoreflect.FieldDescriptor
)

func init() {
	file_amm_amm_proto_init()
	md_Pool = File_amm_amm_proto.Messages().ByName("Pool")
	fd_Pool_pair = md_Pool.Fields().ByName("pair")
	fd_Pool_quote_asset_reserve = md_Pool.Fields().ByName("quote_asset_reserve")
}

var _ protoreflect.Message = (*fastReflection_Pool)(nil)

type fastReflection_Pool Pool

func (x *Pool) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Pool)(x)
}

func (x *Pool) slowProtoReflect() protoreflect.Message {
	mi := &file_amm_amm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Pool_messageType fastReflection_Pool_messageType
var _ protoreflect.MessageType = fastReflection_Pool_messageType{}

type fastReflection_Pool_messageType struct{}

func (x fastReflection_Pool_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Pool)(nil)
}
func (x fastReflection_Pool_messageType) New() protoreflect.Message {
	return new(fastReflection_Pool)
}
func (x fastReflection_Pool_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Pool
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Pool) Descriptor() protoreflect.MessageDescriptor {
	return md_Pool
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Pool) Type() protoreflect.MessageType {
	return _fastReflection_Pool_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Pool) New() protoreflect.Message {
	return new(fastReflection_Pool)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Pool) Interface() protoreflect.ProtoMessage {
	return (*Pool)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Pool) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Pair != "" {
		value := protoreflect.ValueOfString(x.Pair)
		if !f(fd_Pool_pair, value) {
			return
		}
	}
	if x.QuoteAssetReserve != "" {
		value := protoreflect.ValueOfString(x.QuoteAssetReserve)
		if !f(fd_Pool_quote_asset_reserve, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Pool) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "matrix.amm.v1.Pool.pair":
		return x.Pair != ""
	case "matrix.amm.v1.Pool.quote_asset_reserve":
		return x.QuoteAssetReserve != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: matrix.amm.v1.Pool"))
		}
		panic(fmt.Errorf("message matrix.amm.v1.Pool does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "matrix.amm.v1.Pool.pair":
		x.Pair = ""
	case "matrix.amm.v1.Pool.quote_asset_reserve":
		x.QuoteAssetReserve = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: matrix.amm.v1.Pool"))
		}
		panic(fmt.Errorf("message matrix.amm.v1.Pool does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Pool) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "matrix.amm.v1.Pool.pair":
		value := x.Pair
		return protoreflect.ValueOfString(value)
	case "matrix.amm.v1.Pool.quote_asset_reserve":
		value := x.QuoteAssetReserve
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: matrix.amm.v1.Pool"))
		}
		panic(fmt.Errorf("message matrix.amm.v1.Pool does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "matrix.amm.v1.Pool.pair":
		x.Pair = value.Interface().(string)
	case "matrix.amm.v1.Pool.quote_asset_reserve":
		x.QuoteAssetReserve = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: matrix.amm.v1.Pool"))
		}
		panic(fmt.Errorf("message matrix.amm.v1.Pool does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "matrix.amm.v1.Pool.pair":
		panic(fmt.Errorf("field pair of message matrix.amm.v1.Pool is not mutable"))
	case "matrix.amm.v1.Pool.quote_asset_reserve":
		panic(fmt.Errorf("field quote_asset_reserve of message matrix.amm.v1.Pool is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: matrix.amm.v1.Pool"))
		}
		panic(fmt.Errorf("message matrix.amm.v1.Pool does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Pool) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "matrix.amm.v1.Pool.pair":
		return protoreflect.ValueOfString("")
	case "matrix.amm.v1.Pool.quote_asset_reserve":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: matrix.amm.v1.Pool"))
		}
		panic(fmt.Errorf("message matrix.amm.v1.Pool does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Pool) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in matrix.amm.v1.Pool", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Pool) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pool) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Pool) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Pool) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Pool)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Pair)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.QuoteAssetReserve)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Pool)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.QuoteAssetReserve) > 0 {
			i -= len(x.QuoteAssetReserve)
			copy(dAtA[i:], x.QuoteAssetReserve)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.QuoteAssetReserve)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Pair) > 0 {
			i -= len(x.Pair)
			copy(dAtA[i:], x.Pair)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Pair)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Pool)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pool: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Pair", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Pair = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field QuoteAssetReserve", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.QuoteAssetReserve = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: amm/amm.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Direction int32

const (
	Direction_UNDEFINED       Direction = 0
	Direction_ADD_TO_AMM      Direction = 1
	Direction_REMOVE_FROM_AMM Direction = 2
)

// Enum value maps for Direction.
var (
	Direction_name = map[int32]string{
		0: "UNDEFINED",
		1: "ADD_TO_AMM",
		2: "REMOVE_FROM_AMM",
	}
	Direction_value = map[string]int32{
		"UNDEFINED":       0,
		"ADD_TO_AMM":      1,
		"REMOVE_FROM_AMM": 2,
	}
)

func (x Direction) Enum() *Direction {
	p := new(Direction)
	*p = x
	return p
}

func (x Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_amm_amm_proto_enumTypes[0].Descriptor()
}

func (Direction) Type() protoreflect.EnumType {
	return &file_amm_amm_proto_enumTypes[0]
}

func (x Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Direction.Descriptor instead.
func (Direction) EnumDescriptor() ([]byte, []int) {
	return file_amm_amm_proto_rawDescGZIP(), []int{0}
}

type Pool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair              string `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
	QuoteAssetReserve string `protobuf:"bytes,2,opt,name=quote_asset_reserve,json=quoteAssetReserve,proto3" json:"quote_asset_reserve,omitempty"`
}

func (x *Pool) Reset() {
	*x = Pool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_amm_amm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_amm_amm_proto_rawDescGZIP(), []int{0}
}

func (x *Pool) GetPair() string {
	if x != nil {
		return x.Pair
	}
	return ""
}

func (x *Pool) GetQuoteAssetReserve() string {
	if x != nil {
		return x.QuoteAssetReserve
	}
	return ""
}

var File_amm_amm_proto protoreflect.FileDescriptor

var file_amm_amm_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6d, 0x6d, 0x2f, 0x61, 0x6d, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x61, 0x6d, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1d,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a,
	0x04, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x3a, 0x10, 0xf2, 0x9e, 0xd3, 0x8e, 0x03,
	0x0a, 0x0a, 0x06, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x2a, 0x3f, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45,
	0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x44, 0x44, 0x5f, 0x54,
	0x4f, 0x5f, 0x41, 0x4d, 0x4d, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x4d, 0x4f, 0x56,
	0x45, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x41, 0x4d, 0x4d, 0x10, 0x02, 0x42, 0x9e, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x61, 0x6d, 0x6d, 0x2e,
	0x76, 0x31, 0x42, 0x08, 0x41, 0x6d, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x74, 0x72, 0x69,
	0x78, 0x44, 0x61, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6d, 0x6d, 0x3b, 0x61, 0x6d, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x41, 0x58, 0xaa,
	0x02, 0x0d, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x41, 0x6d, 0x6d, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0d, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5c, 0x41, 0x6d, 0x6d, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x19, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5c, 0x41, 0x6d, 0x6d, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x3a, 0x3a, 0x41, 0x6d, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_amm_amm_proto_rawDescOnce sync.Once
	file_amm_amm_proto_rawDescData = file_amm_amm_proto_rawDesc
)

func file_amm_amm_proto_rawDescGZIP() []byte {
	file_amm_amm_proto_rawDescOnce.Do(func() {
		file_amm_amm_proto_rawDescData = protoimpl.X.CompressGZIP(file_amm_amm_proto_rawDescData)
	})
	return file_amm_amm_proto_rawDescData
}

var file_amm_amm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_amm_amm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_amm_amm_proto_goTypes = []interface{}{
	(Direction)(0), // 0: matrix.amm.v1.Direction
	(*Pool)(nil),   // 1: matrix.amm.v1.Pool
}
var file_amm_amm_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_amm_amm_proto_init() }
func file_amm_amm_proto_init() {
	if File_amm_amm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_amm_amm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_amm_amm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_amm_amm_proto_goTypes,
		DependencyIndexes: file_amm_amm_proto_depIdxs,
		EnumInfos:         file_amm_amm_proto_enumTypes,
		MessageInfos:      file_amm_amm_proto_msgTypes,
	}.Build()
	File_amm_amm_proto = out.File
	file_amm_amm_proto_rawDesc = nil
	file_amm_amm_proto_goTypes = nil
	file_amm_amm_proto_depIdxs = nil
}
