// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nibiru/inflation/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the inflation module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// period is the amount of past periods, based on the epochs per period param
	Period uint64 `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
	// skipped_epochs is the number of epochs that have passed while inflation is
	// disabled
	SkippedEpochs uint64 `protobuf:"varint,3,opt,name=skipped_epochs,json=skippedEpochs,proto3" json:"skipped_epochs,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d00e2bb98c08f74, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *GenesisState) GetSkippedEpochs() uint64 {
	if m != nil {
		return m.SkippedEpochs
	}
	return 0
}

// Params holds parameters for the inflation module.
type Params struct {
	// inflation_enabled is the parameter that enables inflation and halts
	// increasing the skipped_epochs
	InflationEnabled bool `protobuf:"varint,1,opt,name=inflation_enabled,json=inflationEnabled,proto3" json:"inflation_enabled,omitempty"`
	// exponential_calculation takes in the variables to calculate exponential
	// inflation
	ExponentialCalculation ExponentialCalculation `protobuf:"bytes,2,opt,name=exponential_calculation,json=exponentialCalculation,proto3" json:"exponential_calculation"`
	// inflation_distribution of the minted denom
	InflationDistribution InflationDistribution `protobuf:"bytes,3,opt,name=inflation_distribution,json=inflationDistribution,proto3" json:"inflation_distribution"`
	// epochs_per_period is the number of epochs that must pass before a new
	// period is created
	EpochsPerPeriod uint64 `protobuf:"varint,4,opt,name=epochs_per_period,json=epochsPerPeriod,proto3" json:"epochs_per_period,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d00e2bb98c08f74, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetInflationEnabled() bool {
	if m != nil {
		return m.InflationEnabled
	}
	return false
}

func (m *Params) GetExponentialCalculation() ExponentialCalculation {
	if m != nil {
		return m.ExponentialCalculation
	}
	return ExponentialCalculation{}
}

func (m *Params) GetInflationDistribution() InflationDistribution {
	if m != nil {
		return m.InflationDistribution
	}
	return InflationDistribution{}
}

func (m *Params) GetEpochsPerPeriod() uint64 {
	if m != nil {
		return m.EpochsPerPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nibiru.inflation.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "nibiru.inflation.v1.Params")
}

func init() { proto.RegisterFile("nibiru/inflation/v1/genesis.proto", fileDescriptor_2d00e2bb98c08f74) }

var fileDescriptor_2d00e2bb98c08f74 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4f, 0xfa, 0x30,
	0x1c, 0xde, 0x80, 0x2c, 0xff, 0x94, 0xbf, 0x2f, 0x54, 0x45, 0x82, 0xc9, 0x44, 0x8c, 0x09, 0x81,
	0x64, 0x0b, 0x78, 0xf2, 0x0a, 0x12, 0xc3, 0xc5, 0x10, 0xbc, 0x79, 0x59, 0xba, 0xad, 0x8e, 0xea,
	0x68, 0x9b, 0xad, 0x23, 0xf8, 0x0d, 0x3c, 0xfa, 0x69, 0xfc, 0x0c, 0x1c, 0x39, 0x7a, 0x32, 0x06,
	0xbe, 0x88, 0xb1, 0x9d, 0x8c, 0xc3, 0x6e, 0xed, 0xf3, 0xd2, 0xe7, 0xe9, 0x2f, 0x3f, 0x70, 0x41,
	0x89, 0x4b, 0xa2, 0xc4, 0x26, 0xf4, 0x29, 0x44, 0x82, 0x30, 0x6a, 0xcf, 0xbb, 0x76, 0x80, 0x29,
	0x8e, 0x49, 0x6c, 0xf1, 0x88, 0x09, 0x06, 0x8f, 0x94, 0xc4, 0xda, 0x4a, 0xac, 0x79, 0xb7, 0x7e,
	0x1c, 0xb0, 0x80, 0x49, 0xde, 0xfe, 0x3d, 0x29, 0x69, 0xfd, 0x32, 0xef, 0xb5, 0xcc, 0x27, 0x45,
	0xcd, 0x37, 0x1d, 0xfc, 0xbf, 0x53, 0x09, 0x0f, 0x02, 0x09, 0x0c, 0x6f, 0x80, 0xc1, 0x51, 0x84,
	0x66, 0x71, 0x4d, 0x6f, 0xe8, 0xad, 0x72, 0xef, 0xcc, 0xca, 0x49, 0xb4, 0xc6, 0x52, 0xd2, 0x2f,
	0x2d, 0xbf, 0xce, 0xb5, 0x49, 0x6a, 0x80, 0x55, 0x60, 0x70, 0x1c, 0x11, 0xe6, 0xd7, 0x0a, 0x0d,
	0xbd, 0x55, 0x9a, 0xa4, 0x37, 0x78, 0x05, 0xf6, 0xe3, 0x17, 0xc2, 0x39, 0xf6, 0x1d, 0xcc, 0x99,
	0x37, 0x8d, 0x6b, 0x45, 0xc9, 0xef, 0xa5, 0xe8, 0x50, 0x82, 0xcd, 0x8f, 0x02, 0x30, 0xd4, 0xbb,
	0xb0, 0x03, 0x2a, 0xdb, 0x38, 0x07, 0x53, 0xe4, 0x86, 0xd8, 0x97, 0x7d, 0xfe, 0x4d, 0x0e, 0xb7,
	0xc4, 0x50, 0xe1, 0xf0, 0x19, 0x9c, 0xe2, 0x05, 0x67, 0x14, 0x53, 0x41, 0x50, 0xe8, 0x78, 0x28,
	0xf4, 0x12, 0xa5, 0x90, 0x3d, 0xca, 0xbd, 0x4e, 0xee, 0x17, 0x86, 0x99, 0x67, 0x90, 0x59, 0xd2,
	0x2f, 0x55, 0x71, 0x2e, 0x0b, 0x03, 0x50, 0xcd, 0x8a, 0xf9, 0x24, 0x16, 0x11, 0x71, 0x13, 0x19,
	0x55, 0x94, 0x51, 0xed, 0xdc, 0xa8, 0xd1, 0xdf, 0xe5, 0x76, 0xc7, 0x91, 0x26, 0x9d, 0x90, 0x3c,
	0x12, 0xb6, 0x41, 0x45, 0xcd, 0xca, 0xe1, 0x38, 0x72, 0xd2, 0xb1, 0x96, 0xe4, 0xd8, 0x0e, 0x14,
	0x31, 0xc6, 0xd1, 0x58, 0xc2, 0xfd, 0xd1, 0x72, 0x6d, 0xea, 0xab, 0xb5, 0xa9, 0x7f, 0xaf, 0x4d,
	0xfd, 0x7d, 0x63, 0x6a, 0xab, 0x8d, 0xa9, 0x7d, 0x6e, 0x4c, 0xed, 0xd1, 0x0e, 0x88, 0x98, 0x26,
	0xae, 0xe5, 0xb1, 0x99, 0x7d, 0x2f, 0x8b, 0x0d, 0xa6, 0x88, 0x50, 0x3b, 0xdd, 0x8c, 0xc5, 0xce,
	0x6e, 0x88, 0x57, 0x8e, 0x63, 0xd7, 0x90, 0x5b, 0x71, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0xdb, 0xd9, 0x20, 0x8a, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SkippedEpochs != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SkippedEpochs))
		i--
		dAtA[i] = 0x18
	}
	if m.Period != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EpochsPerPeriod != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.EpochsPerPeriod))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.InflationDistribution.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.ExponentialCalculation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.InflationEnabled {
		i--
		if m.InflationEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Period != 0 {
		n += 1 + sovGenesis(uint64(m.Period))
	}
	if m.SkippedEpochs != 0 {
		n += 1 + sovGenesis(uint64(m.SkippedEpochs))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InflationEnabled {
		n += 2
	}
	l = m.ExponentialCalculation.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.InflationDistribution.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochsPerPeriod != 0 {
		n += 1 + sovGenesis(uint64(m.EpochsPerPeriod))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SkippedEpochs", wireType)
			}
			m.SkippedEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SkippedEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InflationEnabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExponentialCalculation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExponentialCalculation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationDistribution", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationDistribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochsPerPeriod", wireType)
			}
			m.EpochsPerPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochsPerPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
