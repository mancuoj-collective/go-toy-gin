package gg

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

// 1. 先新建一个 Engine
// 2. 然后往 Engine 的 router 里添加路由及其对应的 handler
// 3. 启动服务，然后所有的请求都会由 engine.ServeHTTP 处理
// 4. 每个请求都新建一个 Context 对象
// 5. 找到请求 URL 对应的 handler
// 6. 调用 handler 处理请求
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
