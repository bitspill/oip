package oip5

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/bitspill/flod/flojson"
	"github.com/oipwg/oip/datastore"
)

func TestIntakeRecordTemplate(t *testing.T) {
	t.SkipNow()

	b, err := base64.StdEncoding.DecodeString("CvMBCh5nb29nbGUvcHJvdG9idWYvZHVyYXRpb24ucHJvdG8SD2dvb2dsZS5wcm90b2J1ZiI6CghEdXJhdGlvbhIYCgdzZWNvbmRzGAEgASgDUgdzZWNvbmRzEhQKBW5hbm9zGAIgASgFUgVuYW5vc0J8ChNjb20uZ29vZ2xlLnByb3RvYnVmQg1EdXJhdGlvblByb3RvUAFaKmdpdGh1Yi5jb20vZ29sYW5nL3Byb3RvYnVmL3B0eXBlcy9kdXJhdGlvbvgBAaICA0dQQqoCHkdvb2dsZS5Qcm90b2J1Zi5XZWxsS25vd25UeXBlc2IGcHJvdG8zCvIKCgdwLnByb3RvEhVvaXA1LnJlY29yZC50ZW1wbGF0ZXMaHmdvb2dsZS9wcm90b2J1Zi9kdXJhdGlvbi5wcm90byLYAQoBUBIQCgNwaWQYASABKAlSA3BpZBISCgRuYW1lGAIgASgJUgRuYW1lEiAKC2Rlc2NyaXB0aW9uGAMgASgJUgtkZXNjcmlwdGlvbhIQCgNsYWIYBCADKAlSA2xhYhIgCgtpbnN0aXR1dGlvbhgFIAMoCVILaW5zdGl0dXRpb24SIAoLZGV2ZWxvcGVkQnkYBiADKAlSC2RldmVsb3BlZEJ5EjUKCGR1cmF0aW9uGAcgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uUghkdXJhdGlvbkILWgl0ZW1wbGF0ZXNKvwgKCAoBDBIDAAASCggKAQISAwIAHgoJCgIDABIDBgAoCgkKAggLEgMEACAKCgoCBAASBAgAJQEKCgoDBAABEgMICAkKNQoEBAACABIDCwQTGiggSW50ZXJuYWwgUHJvdG9jb2wgSUQNCiBFeGFtcGxlOiBOUy0wMDENCgwKBQQAAgABEgMLCw4KDAoFBAACAAUSAwsECgoMCgUEAAIAAxIDCxESCjgKBAQAAgESAw8EFBorIFByb3RvY29sJ3MgbmFtZQ0KIEV4YW1wbGU6IG5lZ2F0aXZlIHN0YWluDQoMCgUEAAIBARIDDwsPCgwKBQQAAgEFEgMPBAoKDAoFBAACAQMSAw8SEwrWAQoEBAACAhIDFQQbGsgBIEJyaWVmIGRlc2NyaXB0aW9uIG9mIHRoZSBtZXRob2QNCiBFeGFtcGxlOg0KIDIgbWljcm8gbGl0ZXJzIG9mIHNhbXBsZSwgd2FpdCBmb3IgNjAgc2Vjb25kcywgYmxvdCB3aXRoIHBhcGVyIDMgdGltZXMsDQogMiBtaWNybyBsaXRlcnMgb2YgdXJhbnlsIGFjZXRhdGUsIHdhaXQgZm9yIDYwIHNlY29uZHMsIGJsb3Qgd2l0aCBwYXBlciAzIHRpbWVzLg0KDAoFBAACAgESAxULFgoMCgUEAAICBRIDFQQKCgwKBQQAAgIDEgMVGRoKXAoEBAACAxIDGQQcGk8gTGlzdCBvZiBsYWJzIGFzc29jaWF0ZWQgd2l0aCB0aGUgc2FtcGxlIGNvbGxlY3Rpb24NCiBFeGFtcGxlOiBbIERleHRlciBMYWJzIF0NCgwKBQQAAgMBEgMZFBcKDAoFBAACAwUSAxkNEwoMCgUEAAIDBBIDGQQMCgwKBQQAAgMDEgMZGhsKeQoEBAACBBIDHQQkGmwgTGlzdCBvZiBuYW1lIG9mIHRoZSBpbnN0aXR1dGlvbiBmcm9tIHRoZSBsYWJzIGludm9sdmVkIGluIHNhbXBsZSBjb2xsZWN0aW9uDQogRXhhbXBsZTogWyBDYXJ0b29uIE5ldHdvcmsgXQ0KDAoFBAACBAESAx0UHwoMCgUEAAIEBRIDHQ0TCgwKBQQAAgQEEgMdBAwKDAoFBAACBAMSAx0iIwpVCgQEAAIFEgMhBCQaSCBMaXN0IG9mIHBlb3BsZSB3aG8gZGV2ZWxvcGVkIHRoZSBwcm90b2NvbA0KIEV4YW1wbGU6IFsgQ2hhcmxpZSwgRG91ZyBdDQoMCgUEAAIFARIDIRQfCgwKBQQAAgUFEgMhDRMKDAoFBAACBQQSAyEEDAoMCgUEAAIFAxIDISIjCjEKBAQAAgYSAyQEKhokIEV4YW1wbGUgb2YgdXNpbmcgYSBzdGFuZGFyZCBpbXBvcnQNCgwKBQQAAgYBEgMkHSUKDAoFBAACBgUSAyQEHAoMCgUEAAIGAxIDJCgpYgZwcm90bzM=")
	if err != nil {
		t.Fatal(err)
	}
	bc := &RecordTemplateProto{
		Description:        "a description",
		FriendlyName:       "Research Protocol BC",
		DescriptorSetProto: b,
		Required:           []uint64{},
		Recommended:        []uint64{0xcafebabe, 0xdeadbeef},
	}

	bctx := &datastore.TransactionData{
		Transaction: &flojson.TxRawResult{
			Txid: "00000000deadbeef",
		},
	}
	cb := &RecordTemplateProto{
		Description:        "a description",
		FriendlyName:       "Research Protocol CB",
		DescriptorSetProto: b,
		Required:           []uint64{},
		Recommended:        []uint64{0xdeadbeef},
	}

	cbtx := &datastore.TransactionData{
		Transaction: &flojson.TxRawResult{
			Txid: "00000000cafebabe",
		},
	}

	err = datastore.Setup(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	bir, err := intakeRecordTemplate(cb, cbtx)
	if err != nil {
		t.Fatal("failed :(")
	}
	datastore.AutoBulk.Add(bir)

	bir, err = intakeRecordTemplate(bc, bctx)
	if err != nil {
		t.Fatal("failed :(")
	}
	datastore.AutoBulk.Add(bir)

	_, err = datastore.AutoBulk.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	_ = bir
}

func TestLoadTemplatesFromES(t *testing.T) {
	t.SkipNow()
	err := datastore.Setup(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	err = LoadTemplatesFromES(context.Background())
	if err != nil {
		t.Fatal(err)
	}
}

func TestDescriptorFromProtobufJs(t *testing.T) {
	// dsc := []byte{10, 113, 10, 27, 111, 105, 112, 53, 95, 114, 101, 99, 111, 114, 100, 95,
	// 	116, 101, 109, 112, 108, 97, 116, 101, 115, 46, 112, 114, 111, 116, 111, 18,
	// 	21, 111, 105, 112, 53, 46, 114, 101, 99, 111, 114, 100, 46, 116, 101, 109,
	// 	112, 108, 97, 116, 101, 115, 34, 51, 10, 1, 80, 18, 14, 10, 6, 102, 114,
	// 	117, 105, 116, 115, 24, 1, 32, 3, 40, 9, 18, 17, 10, 9, 102, 105, 114,
	// 	115, 116, 78, 97, 109, 101, 24, 2, 32, 1, 40, 9, 18, 11, 10, 3, 97,
	// 	103, 101, 24, 3, 32, 1, 40, 5, 98, 6, 112, 114, 111, 116, 111, 51,
	// }

	dsc := []byte{10, 85, 10, 27, 111, 105, 112, 53, 95, 114, 101, 99, 111, 114, 100, 95, 116, 101, 109, 112, 108, 97, 116, 101, 115, 46, 112, 114, 111, 116, 111, 18, 21, 111, 105, 112, 53, 46, 114, 101, 99, 111, 114, 100, 46, 116, 101, 109, 112, 108, 97, 116, 101, 115, 34, 23, 10, 1, 80, 18, 18, 10, 10, 102, 114, 117, 105, 116, 115, 32, 114, 114, 114, 24, 1, 32, 3, 40, 9, 98, 6, 112, 114, 111, 116, 111, 51}

	bc := &RecordTemplateProto{
		Description:        "Test generated from protobuf.js",
		FriendlyName:       "Protobuf.js test",
		DescriptorSetProto: dsc,
	}

	bctx := &datastore.TransactionData{
		Transaction: &flojson.TxRawResult{
			Txid: "000000000badbabe",
		},
	}

	err := datastore.Setup(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	bir, err := intakeRecordTemplate(bc, bctx)
	if err != nil {
		t.Fatal("failed :(")
	}
	datastore.AutoBulk.Add(bir)

	_, err = datastore.AutoBulk.Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	_ = bir
}
