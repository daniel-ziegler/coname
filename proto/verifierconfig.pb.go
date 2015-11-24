// Code generated by protoc-gen-gogo.
// source: verifierconfig.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/andres-erbsen/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "gogoproto"

import bytes "bytes"

import strings "strings"
import github_com_andres_erbsen_protobuf_proto "github.com/andres-erbsen/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VerifierConfig struct {
	ID                   uint64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SigningKeyID         string              `protobuf:"bytes,2,opt,name=signing_key_id,proto3" json:"signing_key_id,omitempty"`
	Realm                string              `protobuf:"bytes,3,opt,name=realm,proto3" json:"realm,omitempty"`
	TLS                  *TLSConfig          `protobuf:"bytes,4,opt,name=tls" json:"tls,omitempty"`
	KeyserverAddr        string              `protobuf:"bytes,5,opt,name=keyserver_addr,proto3" json:"keyserver_addr,omitempty"`
	InitialKeyserverAuth AuthorizationPolicy `protobuf:"bytes,6,opt,name=initial_keyserver_auth" json:"initial_keyserver_auth"`
	TreeNonce            []byte              `protobuf:"bytes,7,opt,name=tree_nonce,proto3" json:"tree_nonce,omitempty"`
}

func (m *VerifierConfig) Reset()      { *m = VerifierConfig{} }
func (*VerifierConfig) ProtoMessage() {}

func (m *VerifierConfig) GetTLS() *TLSConfig {
	if m != nil {
		return m.TLS
	}
	return nil
}

func (m *VerifierConfig) GetInitialKeyserverAuth() AuthorizationPolicy {
	if m != nil {
		return m.InitialKeyserverAuth
	}
	return AuthorizationPolicy{}
}

func (this *VerifierConfig) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*VerifierConfig)
	if !ok {
		return fmt.Errorf("that is not of type *VerifierConfig")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *VerifierConfig but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *VerifierConfigbut is not nil && this == nil")
	}
	if this.ID != that1.ID {
		return fmt.Errorf("ID this(%v) Not Equal that(%v)", this.ID, that1.ID)
	}
	if this.SigningKeyID != that1.SigningKeyID {
		return fmt.Errorf("SigningKeyID this(%v) Not Equal that(%v)", this.SigningKeyID, that1.SigningKeyID)
	}
	if this.Realm != that1.Realm {
		return fmt.Errorf("Realm this(%v) Not Equal that(%v)", this.Realm, that1.Realm)
	}
	if !this.TLS.Equal(that1.TLS) {
		return fmt.Errorf("TLS this(%v) Not Equal that(%v)", this.TLS, that1.TLS)
	}
	if this.KeyserverAddr != that1.KeyserverAddr {
		return fmt.Errorf("KeyserverAddr this(%v) Not Equal that(%v)", this.KeyserverAddr, that1.KeyserverAddr)
	}
	if !this.InitialKeyserverAuth.Equal(&that1.InitialKeyserverAuth) {
		return fmt.Errorf("InitialKeyserverAuth this(%v) Not Equal that(%v)", this.InitialKeyserverAuth, that1.InitialKeyserverAuth)
	}
	if !bytes.Equal(this.TreeNonce, that1.TreeNonce) {
		return fmt.Errorf("TreeNonce this(%v) Not Equal that(%v)", this.TreeNonce, that1.TreeNonce)
	}
	return nil
}
func (this *VerifierConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*VerifierConfig)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.SigningKeyID != that1.SigningKeyID {
		return false
	}
	if this.Realm != that1.Realm {
		return false
	}
	if !this.TLS.Equal(that1.TLS) {
		return false
	}
	if this.KeyserverAddr != that1.KeyserverAddr {
		return false
	}
	if !this.InitialKeyserverAuth.Equal(&that1.InitialKeyserverAuth) {
		return false
	}
	if !bytes.Equal(this.TreeNonce, that1.TreeNonce) {
		return false
	}
	return true
}
func (this *VerifierConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&proto.VerifierConfig{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "SigningKeyID: "+fmt.Sprintf("%#v", this.SigningKeyID)+",\n")
	s = append(s, "Realm: "+fmt.Sprintf("%#v", this.Realm)+",\n")
	if this.TLS != nil {
		s = append(s, "TLS: "+fmt.Sprintf("%#v", this.TLS)+",\n")
	}
	s = append(s, "KeyserverAddr: "+fmt.Sprintf("%#v", this.KeyserverAddr)+",\n")
	s = append(s, "InitialKeyserverAuth: "+strings.Replace(this.InitialKeyserverAuth.GoString(), `&`, ``, 1)+",\n")
	s = append(s, "TreeNonce: "+fmt.Sprintf("%#v", this.TreeNonce)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringVerifierconfig(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringVerifierconfig(e map[int32]github_com_andres_erbsen_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *VerifierConfig) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *VerifierConfig) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		data[i] = 0x8
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(m.ID))
	}
	if len(m.SigningKeyID) > 0 {
		data[i] = 0x12
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(len(m.SigningKeyID)))
		i += copy(data[i:], m.SigningKeyID)
	}
	if len(m.Realm) > 0 {
		data[i] = 0x1a
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(len(m.Realm)))
		i += copy(data[i:], m.Realm)
	}
	if m.TLS != nil {
		data[i] = 0x22
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(m.TLS.Size()))
		n1, err := m.TLS.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.KeyserverAddr) > 0 {
		data[i] = 0x2a
		i++
		i = encodeVarintVerifierconfig(data, i, uint64(len(m.KeyserverAddr)))
		i += copy(data[i:], m.KeyserverAddr)
	}
	data[i] = 0x32
	i++
	i = encodeVarintVerifierconfig(data, i, uint64(m.InitialKeyserverAuth.Size()))
	n2, err := m.InitialKeyserverAuth.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.TreeNonce != nil {
		if len(m.TreeNonce) > 0 {
			data[i] = 0x3a
			i++
			i = encodeVarintVerifierconfig(data, i, uint64(len(m.TreeNonce)))
			i += copy(data[i:], m.TreeNonce)
		}
	}
	return i, nil
}

func encodeFixed64Verifierconfig(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Verifierconfig(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintVerifierconfig(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedVerifierConfig(r randyVerifierconfig, easy bool) *VerifierConfig {
	this := &VerifierConfig{}
	this.ID = uint64(uint64(r.Uint32()))
	this.SigningKeyID = randStringVerifierconfig(r)
	this.Realm = randStringVerifierconfig(r)
	if r.Intn(10) != 0 {
		this.TLS = NewPopulatedTLSConfig(r, easy)
	}
	this.KeyserverAddr = randStringVerifierconfig(r)
	v1 := NewPopulatedAuthorizationPolicy(r, easy)
	this.InitialKeyserverAuth = *v1
	v2 := r.Intn(100)
	this.TreeNonce = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.TreeNonce[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyVerifierconfig interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneVerifierconfig(r randyVerifierconfig) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringVerifierconfig(r randyVerifierconfig) string {
	v3 := r.Intn(100)
	tmps := make([]rune, v3)
	for i := 0; i < v3; i++ {
		tmps[i] = randUTF8RuneVerifierconfig(r)
	}
	return string(tmps)
}
func randUnrecognizedVerifierconfig(r randyVerifierconfig, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldVerifierconfig(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldVerifierconfig(data []byte, r randyVerifierconfig, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		v4 := r.Int63()
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		data = encodeVarintPopulateVerifierconfig(data, uint64(v4))
	case 1:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateVerifierconfig(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateVerifierconfig(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateVerifierconfig(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *VerifierConfig) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovVerifierconfig(uint64(m.ID))
	}
	l = len(m.SigningKeyID)
	if l > 0 {
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	l = len(m.Realm)
	if l > 0 {
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	if m.TLS != nil {
		l = m.TLS.Size()
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	l = len(m.KeyserverAddr)
	if l > 0 {
		n += 1 + l + sovVerifierconfig(uint64(l))
	}
	l = m.InitialKeyserverAuth.Size()
	n += 1 + l + sovVerifierconfig(uint64(l))
	if m.TreeNonce != nil {
		l = len(m.TreeNonce)
		if l > 0 {
			n += 1 + l + sovVerifierconfig(uint64(l))
		}
	}
	return n
}

func sovVerifierconfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozVerifierconfig(x uint64) (n int) {
	return sovVerifierconfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *VerifierConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VerifierConfig{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`SigningKeyID:` + fmt.Sprintf("%v", this.SigningKeyID) + `,`,
		`Realm:` + fmt.Sprintf("%v", this.Realm) + `,`,
		`TLS:` + strings.Replace(fmt.Sprintf("%v", this.TLS), "TLSConfig", "TLSConfig", 1) + `,`,
		`KeyserverAddr:` + fmt.Sprintf("%v", this.KeyserverAddr) + `,`,
		`InitialKeyserverAuth:` + strings.Replace(strings.Replace(this.InitialKeyserverAuth.String(), "AuthorizationPolicy", "AuthorizationPolicy", 1), `&`, ``, 1) + `,`,
		`TreeNonce:` + fmt.Sprintf("%v", this.TreeNonce) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringVerifierconfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *VerifierConfig) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifierconfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VerifierConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifierConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigningKeyID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Realm", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Realm = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TLS", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TLS == nil {
				m.TLS = &TLSConfig{}
			}
			if err := m.TLS.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyserverAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyserverAddr = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialKeyserverAuth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialKeyserverAuth.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TreeNonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthVerifierconfig
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TreeNonce = append([]byte{}, data[iNdEx:postIndex]...)
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerifierconfig(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifierconfig
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
func skipVerifierconfig(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerifierconfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVerifierconfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthVerifierconfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowVerifierconfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipVerifierconfig(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthVerifierconfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerifierconfig   = fmt.Errorf("proto: integer overflow")
)
