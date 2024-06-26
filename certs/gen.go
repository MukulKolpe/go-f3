// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package certs

import (
	"fmt"
	"io"
	"math"
	"sort"

	gpbft "github.com/filecoin-project/go-f3/gpbft"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
	big "math/big"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufPowerTableDelta = []byte{131}

func (t *PowerTableDelta) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufPowerTableDelta); err != nil {
		return err
	}

	// t.ParticipantID (gpbft.ActorID) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.ParticipantID)); err != nil {
		return err
	}

	// t.PowerDelta (big.Int) (struct)
	{
		if err := cw.CborWriteHeader(cbg.MajTag, 2); err != nil {
			return err
		}
		var b []byte
		if t.PowerDelta != nil {
			b = t.PowerDelta.Bytes()
		}

		if err := cw.CborWriteHeader(cbg.MajByteString, uint64(len(b))); err != nil {
			return err
		}
		if _, err := cw.Write(b); err != nil {
			return err
		}
	}

	// t.SigningKey (gpbft.PubKey) (slice)
	if len(t.SigningKey) > 2097152 {
		return xerrors.Errorf("Byte array in field t.SigningKey was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.SigningKey))); err != nil {
		return err
	}

	if _, err := cw.Write(t.SigningKey); err != nil {
		return err
	}

	return nil
}

func (t *PowerTableDelta) UnmarshalCBOR(r io.Reader) (err error) {
	*t = PowerTableDelta{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ParticipantID (gpbft.ActorID) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ParticipantID = gpbft.ActorID(extra)

	}
	// t.PowerDelta (big.Int) (struct)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if maj != cbg.MajTag || extra != 2 {
		return fmt.Errorf("big ints should be cbor bignums")
	}

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if maj != cbg.MajByteString {
		return fmt.Errorf("big ints should be tagged cbor byte strings")
	}

	if extra > 256 {
		return fmt.Errorf("t.PowerDelta: cbor bignum was too large")
	}

	if extra > 0 {
		buf := make([]byte, extra)
		if _, err := io.ReadFull(cr, buf); err != nil {
			return err
		}
		t.PowerDelta = big.NewInt(0).SetBytes(buf)
	} else {
		t.PowerDelta = big.NewInt(0)
	}
	// t.SigningKey (gpbft.PubKey) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > 2097152 {
		return fmt.Errorf("t.SigningKey: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.SigningKey = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.SigningKey); err != nil {
		return err
	}

	return nil
}

var lengthBufFinalityCertificate = []byte{134}

func (t *FinalityCertificate) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufFinalityCertificate); err != nil {
		return err
	}

	// t.GPBFTInstance (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.GPBFTInstance)); err != nil {
		return err
	}

	// t.ECChain (gpbft.ECChain) (slice)
	if len(t.ECChain) > 8192 {
		return xerrors.Errorf("Slice value in field t.ECChain was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.ECChain))); err != nil {
		return err
	}
	for _, v := range t.ECChain {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}

	// t.SupplementalData (gpbft.SupplementalData) (struct)
	if err := t.SupplementalData.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Signers (bitfield.BitField) (struct)
	if err := t.Signers.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Signature ([]uint8) (slice)
	if len(t.Signature) > 2097152 {
		return xerrors.Errorf("Byte array in field t.Signature was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Signature))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Signature); err != nil {
		return err
	}

	// t.PowerTableDelta ([]certs.PowerTableDelta) (slice)
	if len(t.PowerTableDelta) > 8192 {
		return xerrors.Errorf("Slice value in field t.PowerTableDelta was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.PowerTableDelta))); err != nil {
		return err
	}
	for _, v := range t.PowerTableDelta {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}

	}
	return nil
}

func (t *FinalityCertificate) UnmarshalCBOR(r io.Reader) (err error) {
	*t = FinalityCertificate{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 6 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.GPBFTInstance (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.GPBFTInstance = uint64(extra)

	}
	// t.ECChain (gpbft.ECChain) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > 8192 {
		return fmt.Errorf("t.ECChain: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.ECChain = make([]gpbft.TipSet, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				if err := t.ECChain[i].UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.ECChain[i]: %w", err)
				}

			}

		}
	}
	// t.SupplementalData (gpbft.SupplementalData) (struct)

	{

		if err := t.SupplementalData.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.SupplementalData: %w", err)
		}

	}
	// t.Signers (bitfield.BitField) (struct)

	{

		if err := t.Signers.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Signers: %w", err)
		}

	}
	// t.Signature ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > 2097152 {
		return fmt.Errorf("t.Signature: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Signature = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Signature); err != nil {
		return err
	}

	// t.PowerTableDelta ([]certs.PowerTableDelta) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > 8192 {
		return fmt.Errorf("t.PowerTableDelta: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.PowerTableDelta = make([]PowerTableDelta, extra)
	}

	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error
			_ = maj
			_ = extra
			_ = err

			{

				if err := t.PowerTableDelta[i].UnmarshalCBOR(cr); err != nil {
					return xerrors.Errorf("unmarshaling t.PowerTableDelta[i]: %w", err)
				}

			}

		}
	}
	return nil
}
