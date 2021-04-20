package wrapper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func Test_convertToZapAny(t *testing.T) {
	tests := []struct {
		name string
		args []interface{}
		want []zapcore.Field
	}{
		{"nil args", nil, []zapcore.Field{}},
		{"empty args", []interface{}{}, []zapcore.Field{}},
		{"even args number w invalid pairs", []interface{}{1, 2, "keyTwo", 3.14, 1.1, "one-one"}, []zapcore.Field{
			{Key: "arg0", Type: zapcore.Int64Type, Integer: 1},
			{Key: "arg1", Type: zapcore.Int64Type, Integer: 2},
			{Key: "keyTwo", Type: zapcore.Float64Type, Integer: 4614253070214989087},
			{Key: "arg4", Type: zapcore.Float64Type, Integer: 4607632778762754458},
			{Key: "arg5", Type: zapcore.StringType, String: "one-one"},
		}},
		{"non-even args w invalid pairs", []interface{}{"keyOne", 1, "keyTwo"}, []zapcore.Field{
			{Key: "arg0", Type: zapcore.StringType, String: "keyOne"},
			{Key: "arg1", Type: zapcore.Int64Type, Integer: 1},
			{Key: "arg2", Type: zapcore.StringType, String: "keyTwo"},
		}},
		{"single int arg", []interface{}{1}, []zapcore.Field{
			{Key: "arg0", Type: zapcore.Int64Type, Integer: 1},
		}},
		{"single string arg", []interface{}{"something"}, []zapcore.Field{
			{Key: "arg0", Type: zapcore.StringType, String: "something"},
		}},
		{"invalid pair", []interface{}{1, 3.14}, []zapcore.Field{
			{Key: "arg0", Type: zapcore.Int64Type, Integer: 1},
			{Key: "arg1", Type: zapcore.Float64Type, Integer: 4614253070214989087},
		}},
		{"even args w valid pairs", []interface{}{"keyOne", 1, "keyTwo", true, "keyThree", "three"}, []zapcore.Field{
			{Key: "keyOne", Type: zapcore.Int64Type, Integer: 1},
			{Key: "keyTwo", Type: zapcore.BoolType, Integer: 1},
			{Key: "keyThree", Type: zapcore.StringType, String: "three"},
		}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := convertToZapAny(tt.args...)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}
