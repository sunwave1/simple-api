package fiber

import (
	"github.com/gofiber/fiber"
	"sync"
)

var (
	fbOnce sync.Once
	fb     *fiber.App
)

func GetFiber() *fiber.App {
	fbOnce.Do(func() {
		fb = fiber.New()
	})
	return fb
}
