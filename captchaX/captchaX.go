package captchaX

import (
	"git-pd.megvii-inc.com/bbu-pd/wukong-platform/cmd/sys_app"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

func Generate(width, height int) (id, b64s string, err error) {
	driver := base64Captcha.NewDriverString(height, width, 0, 0, 4, "1234567890", &color.RGBA{254, 254, 254, 254}, sys_app.DefaultEmbeddedFonts, []string{"RitaSmith.ttf"})
	c := base64Captcha.NewCaptcha(driver, store)
	return c.Generate()
}

//verify the captcha
func Verify(id, verifyValue string) bool {
	return store.Verify(id, verifyValue, true)
}
