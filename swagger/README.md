// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 object user.CreateResponse 成功后返回值
// @Failure 409 object model.Result 失败后返回值
// @Router /user [post]

// @Summary 通过文章 id 获取单个文章内容
// @version 1.0
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 object model.Result 成功后返回值
// @Failure 409 object model.Result 失败后返回值
// @Router /article/{id} [get]

// @Summary 获取banner活动
// @Tags  banner 活动管理
// @Accept  json
// @Produce json
// @Param   school query string false "通过学校搜索"
// @Param   status query int false "通过活动状态搜索，1 表示进行中，2 表示已结束"
// @Param   release_section query string false "通过活动板块搜索"
// @Param   nickname query string false "通过发布人昵称搜索"
// @Success 200 object app.Response
// @Failure 500 {string} json  "{"code":500,"data":{},"msg":"fail"}"
// @Router /api/banner-activities [get]

Summary：简单阐述 API 的功能
Description：API 详细描述
Tags：API 所属分类
Accept：API 接收参数的格式
Produce：输出的数据格式，这里是 JSON 格式
Param：参数，分为 6 个字段，其中第 6 个字段是可选的，各字段含义为：
    参数名称
    参数在 HTTP 请求中的位置（body、path、query）
    参数类型（string、int、bool 等）
    是否必须（true、false）
    参数描述
    选项，这里用的是 default() 用来指定默认值

Success：成功返回数据格式，分为 4 个字段
    HTTP 返回 Code
    返回数据类型
    返回数据模型
    说明

路由格式，分为 2 个字段：
    API 路径
    HTTP 方法