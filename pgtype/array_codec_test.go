package pgtype_test

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype/testutil"
	"github.com/stretchr/testify/assert"
)

func TestArrayCodec(t *testing.T) {
	conn := testutil.MustConnectPgx(t)
	defer testutil.MustCloseContext(t, conn)

	for i, tt := range []struct {
		expected interface{}
	}{
		{[]int16(nil)},
		{[]int16{}},
		{[]int16{1, 2, 3}},
	} {
		var actual []int16
		err := conn.QueryRow(
			context.Background(),
			"select $1::smallint[]",
			tt.expected,
		).Scan(&actual)
		assert.NoErrorf(t, err, "%d", i)
		assert.Equalf(t, tt.expected, actual, "%d", i)
	}

	newInt16 := func(n int16) *int16 { return &n }

	for i, tt := range []struct {
		expected interface{}
	}{
		{[]*int16{newInt16(1), nil, newInt16(3), nil, newInt16(5)}},
	} {
		var actual []*int16
		err := conn.QueryRow(
			context.Background(),
			"select $1::smallint[]",
			tt.expected,
		).Scan(&actual)
		assert.NoErrorf(t, err, "%d", i)
		assert.Equalf(t, tt.expected, actual, "%d", i)
	}
}

func TestArrayCodecAnySlice(t *testing.T) {
	conn := testutil.MustConnectPgx(t)
	defer testutil.MustCloseContext(t, conn)

	type _int16Slice []int16

	for i, tt := range []struct {
		expected interface{}
	}{
		{_int16Slice(nil)},
		{_int16Slice{}},
		{_int16Slice{1, 2, 3}},
	} {
		var actual _int16Slice
		err := conn.QueryRow(
			context.Background(),
			"select $1::smallint[]",
			tt.expected,
		).Scan(&actual)
		assert.NoErrorf(t, err, "%d", i)
		assert.Equalf(t, tt.expected, actual, "%d", i)
	}
}
