package routers

import (
	"Go_PersonalWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/sendEmail", &controllers.SendEmail{}, "post:ContactMe")

	/**
	多级路由导致页面不能正常加载css样式,好像是views设置不对
	我解决的方法，是“../”，找到css的位置，而不是直接使用相对位置
	 */
	beego.Router("/blog/editor/", &controllers.Editor{})
	beego.Router("/blog", &controllers.Blog{}, "get:ShowBlog")
	beego.Router("/blog/show/*", &controllers.Blog{}, "get:ShowFile")
	beego.Router("/blog/blockChain",&controllers.Blog{}, "get:ShowBlockChain")
	beego.Router("/blog/Cryptography",&controllers.Blog{}, "get:ShowCryptography")
	beego.Router("/blog/Technology",&controllers.Blog{}, "get:ShowTechnology")
	/**
	相册展示
	 */
	beego.Router("/album", &controllers.Album{})
	beego.Router("/album/show", &controllers.ShowAlbum{})

	/**
	自我简介
	 */
	beego.Router("/about", &controllers.About{})

	/**
	抖音骚操作
	 */
	beego.Router("/list", &controllers.ShowLove{}, "get:GetList")
	beego.Router("/love1", &controllers.ShowLove{},"get:Get1")
	beego.Router("/love2", &controllers.ShowLove{},"get:Get2")

}
