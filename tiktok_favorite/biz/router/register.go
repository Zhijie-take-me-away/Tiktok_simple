// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	favorite "Tiktok_simple/tiktok_favorite/biz/router/favorite"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	favorite.Register(r)
}