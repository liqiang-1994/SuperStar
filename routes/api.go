package routes

import (
	"SuperStar/handlers"
	"SuperStar/internal/middlemare"
	"github.com/gofiber/fiber/v2"
)

func addAPIGroup(router *fiber.App,
	account *handlers.AccountHandler,
	login *handlers.LoginHandler,
	poem *handlers.PoemHandler,
	tag *handlers.TagHandler,
	circle *handlers.CircleHandler,
	storage *handlers.StorageHandler,
	jwtMiddle *middlemare.JwtMiddleware,
) {
	api := router.Group("/api")
	accountApi := api.Group("/account")
	{
		accountApi.Get("/:id", account.GetUserByID)
	}
	api.Post("/login", login.Login)
	api.Post("/sendSms", login.SendSms)
	poemApi := api.Group("/poem")
	{
		poemApi.Post("/list", jwtMiddle.JWTProtected, poem.QueryPoemList)
	}
	tagApi := api.Group("/tag")
	{
		tagApi.Post("/create", jwtMiddle.JWTProtected, tag.CreateTag)
		tagApi.Post("/all", jwtMiddle.JWTProtected, tag.QueryAllTag)
		tagApi.Post("/createRel", jwtMiddle.JWTProtected, tag.CreateTagRel)
	}
	circleApi := api.Group("/circle")
	{
		circleApi.Post("/create", jwtMiddle.JWTProtected, circle.CreateCircle)
		circleApi.Get("/:id", jwtMiddle.JWTProtected, circle.QueryCircleDetail)
	}
	storageApi := api.Group("/storage")
	{
		storageApi.Post("/avatar/upload", storage.UploadAvatar)
		storageApi.Get("/download", storage.Download)
	}
}
