package array

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format(time.DateTime))
}
