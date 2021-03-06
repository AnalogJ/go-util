package utils_test

import (
	"github.com/analogj/go-util/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMapKeys(t *testing.T) {
	t.Parallel()

	//test
	actual := utils.MapKeys(map[string]interface{}{"example": "17"})

	//assert
	require.Equal(t, []string{"example"}, actual, "should correctly retrieve keys from a map")
}

func TestMapKeys_Multiple(t *testing.T) {
	t.Parallel()

	//test
	actual := utils.MapKeys(map[string]interface{}{"example": "17", "example2": "18"})

	//assert
	require.Equal(t, []string{"example", "example2"}, actual, "should correctly retrieve keys from a map")
}

func TestMapDeepCopy(t *testing.T) {
	t.Parallel()

	//test
	orig := map[string]interface{}{"example": map[string]interface{}{"nested": "17", "other": "18"}}
	actual, err := utils.MapDeepCopy(orig)
	orig["example"].(map[string]interface{})["nested"] = "18"

	//assert
	require.NoError(t, err, "Should finish successfully")
	require.Equal(t, "17", actual["example"].(map[string]interface{})["nested"], "should create deep cloned data structure")
	require.Equal(t, "18", actual["example"].(map[string]interface{})["other"], "should create deep cloned data structure")
}

func TestMapMergeGeneric(t *testing.T) {
	t.Parallel()

	//test
	a := map[string]interface{}{"foo": "bar", "a": "b"}
	b := map[string]interface{}{"foo": "baz", "c": "d"}
	merged := utils.MapMergeGeneric(a, b)

	//assert
	require.Equal(t, "baz", merged["foo"])
	require.Equal(t, "b", merged["a"])
	require.Equal(t, "d", merged["c"])
}

func TestMapMerge(t *testing.T) {
	t.Parallel()

	//test
	a := map[string]string{"foo": "bar", "a": "b"}
	b := map[string]string{"foo": "baz", "c": "d"}
	merged := utils.MapMerge(a, b)

	//assert
	require.Equal(t, "baz", merged["foo"])
	require.Equal(t, "b", merged["a"])
	require.Equal(t, "d", merged["c"])
}
