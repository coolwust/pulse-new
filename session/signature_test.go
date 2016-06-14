package session

import (
	"testing"
)

var signTests = []struct {
	unsigned string
	key      []byte
	signed   string
}{
	{"hello", []byte("foo"), "hello.NAJWxCdrLTSD8CvzyLdIDhLH6pcCsiAKldMySgs4"},
	{"world", []byte("bar"), "world.6kpaLfNqzist38XcIeU9ejULKGCJK7u23K9Qwbyb4sk"},
}

func TestSign(t *testing.T) {
	for i, tt := range signTests {
		if signed := sign(tt.unsigned, tt.key); signed != tt.signed {
			t.Errorf("%d: sign(%#v, %#v) = %#v, want %#v", i, tt.unsigned, tt.key, signed, tt.signed)
		}
	}
}

var unsignTests = []struct {
	signed   string
	keys     [][]byte
	unsigned string
	err      error
}{
	{"hello.NAJWxCdrLTSD8CvzyLdIDhLH6pcCsiAKldMySgs4", [][]byte{[]byte("foo"), []byte("bar")}, "hello", nil},
	{"world.6kpaLfNqzist38XcIeU9ejULKGCJK7u23K9Qwbyb4sk", [][]byte{[]byte("foo"), []byte("bar")}, "world", nil},
	{"hello.NAJWxCdrLTSD8CvzyLdIDhLH6pcCsiAKldMySgs5", [][]byte{[]byte("foo"), []byte("bar")}, "", ErrInvalidSignature},
	{"world.6kpaLfNqzist38XcIeU9ejULKGCJK7u23K9Qwbyb4sk", [][]byte{[]byte("baz"), []byte("qux")}, "", ErrInvalidSignature},
}

func TestUnsign(t *testing.T) {
	for i, tt := range unsignTests {
		if unsigned, err := unsign(tt.signed, tt.keys); unsigned != tt.unsigned || err != tt.err {
			t.Errorf("%d: unsign(%#v, %#v) = %#v, %#v, want %#v, %#v", i, tt.signed, tt.keys, unsigned, err, tt.unsigned, tt.err)
		}
	}
}
