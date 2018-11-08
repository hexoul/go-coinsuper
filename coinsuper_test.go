package coinsuper

import (
	"testing"
)

func init() {
	GetInstanceWithKey("YOUR_ACCESS_KEY", "YOUR_SECRET_KEY")
}

func TestUserAssetInfo(t *testing.T) {
	if _, err := GetInstance().UserAssetInfo(nil); err != nil {
		t.Fatal(err)
	}
}
