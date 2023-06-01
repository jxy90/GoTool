package captchaX

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaptchaX(t *testing.T) {
	id, bs64s, err := Generate(100, 100)
	assert := assert.New(t)
	if err != nil {
		assert.Fail("err", err)
	}
	fmt.Printf("id: %v bs64s: %v", id, bs64s)
}
