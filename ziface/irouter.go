package ziface

type IRouter interface {
	// PreHandle 在处理conn业务之前钩子方法Hook
	PreHandle(request IRequest)

	// Handle 在处理conn业务住方法
	Handle(request IRequest)

	// PostHandle 在处理conn业务之后钩子方法Hook
	PostHandle(request IRequest)
}
