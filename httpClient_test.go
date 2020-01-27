package zeusclient

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestZeusHttpClient_CheckPerm(t *testing.T) {
	Convey("验证用户权限", t, func() {
		c := NewZeusHttpClient(
			"http://localhost:8082/v1",
			"crawlnovel",
			"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODAxMzkwNzgsImlkIjoyLCJuYW1lIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4MDA1MjY3OCwidWlkIjoyLCJ1bmFtZSI6ImFkbWluIn0.UHUx3faF-Knz1Bc0W1EtuWNVsK_OJQzGowGu-2H9v1ADkqeefBSv663aLMRkjpfFyFzS1jZN6emX8zrOeasHVssJnf3dfU9N8P-vE-GM0oiKqa84pkJX5IVY6V-HI-pTl8Zhbdg7eWMcY4_31j_71b5m-HJ8eLT2FMJ90fQkw2w")
		can, err := c.CheckPerm("/taskmanage/list")
		So(err, ShouldEqual, nil)
		So(can, ShouldEqual, false)
	})
}

func TestZeusHttpClient_GetUserPerms(t *testing.T) {
	Convey("用户所有权限", t, func() {
		c := NewZeusHttpClient(
			"http://localhost:8082/v1",
			"crawlnovel",
			"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODAxMzkwNzgsImlkIjoyLCJuYW1lIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4MDA1MjY3OCwidWlkIjoyLCJ1bmFtZSI6ImFkbWluIn0.UHUx3faF-Knz1Bc0W1EtuWNVsK_OJQzGowGu-2H9v1ADkqeefBSv663aLMRkjpfFyFzS1jZN6emX8zrOeasHVssJnf3dfU9N8P-vE-GM0oiKqa84pkJX5IVY6V-HI-pTl8Zhbdg7eWMcY4_31j_71b5m-HJ8eLT2FMJ90fQkw2w")
		perms, err := c.GetUserPerms()
		So(err, ShouldEqual, nil)
		So(len(perms), ShouldEqual, 1)
	})
}

func TestZeusHttpClient_GetUserMenus(t *testing.T) {
	Convey("用户所有菜单", t, func() {
		c := NewZeusHttpClient(
			"http://localhost:8082/v1",
			"crawlnovel",
			"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODAxMzkwNzgsImlkIjoyLCJuYW1lIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTU4MDA1MjY3OCwidWlkIjoyLCJ1bmFtZSI6ImFkbWluIn0.UHUx3faF-Knz1Bc0W1EtuWNVsK_OJQzGowGu-2H9v1ADkqeefBSv663aLMRkjpfFyFzS1jZN6emX8zrOeasHVssJnf3dfU9N8P-vE-GM0oiKqa84pkJX5IVY6V-HI-pTl8Zhbdg7eWMcY4_31j_71b5m-HJ8eLT2FMJ90fQkw2w")
		perms, err := c.GetUserMenus()
		So(err, ShouldEqual, nil)
		So(len(perms), ShouldEqual, 5)
	})
}