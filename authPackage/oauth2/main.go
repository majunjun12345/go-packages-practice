package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

/*
	比如:尝试使用 github 登录 github 并进行身份验证
	条件: 用户有 github 的账号

	在 Oauth 流程都有三个参与者:
		客户端(用户): 登录的人员或用户
		使用者(应用服务器): 客户端想要登录的应用程序 gitlab
		服务提供者(认证服务器): 用户通过其进行身份验证的外部应用程序 github

	流程:
		1. github 要求用户登录,用户选择使用 github 进行登录
		2. gitlab 提供的 github 链接将会把用户带到 github 登录界面
		   应用index: http://localhost:8000/
		   应用登录: http://localhost:8000/GithubLogin
		   会获取到这个 url: https://github.com/login/oauth/authorize?client_id=6155ffa6f3921c4e67d5&redirect_uri=http%3A%2F%2Flocalhost%3A8000%2FGithubCallback&response_type=code&scope=user+repo&state=random
		   重定向这个链接,会发出请求,定位到 github 的登录界面: https://github.com/login?client_id=6155ffa6f3921c4e67d5&return_to=%2Flogin%2Foauth%2Fauthorize%3Fclient_id%3D6155ffa6f3921c4e67d5%26redirect_uri%3Dhttp%253A%252F%252Flocalhost%253A8000%252FGithubCallback%26response_type%3Dcode%26scope%3Duser%2Brepo%26state%3Drandom
		   点击登录后, 会请求这个 url: https://github.com/login/oauth/authorize?client_id=6155ffa6f3921c4e67d5&redirect_uri=http%3A%2F%2Flocalhost%3A8000%2FGithubCallback&response_type=code&scope=user+repo&state=random
		   最后请求回调 url: http://localhost:8000/GithubCallback?code=9c82da7fb8693660399e&state=random
		3. 在 github 界面登录后,选择授权,也就是允许 gitlab 获取用户资料
		4. 浏览器通过几次跳转后返回 gitlab,这时候我们已经完成了认证登录;

	跳转过程中 gitlab 与 github 的几次交互:
		1. 选择授权的时候,github 服务器会根据 gitlab 转到 github 时候给出的重定向链接返回给 github 一个 code.
		   这个 code 代表 github 的登录服务器认可 gitlab 这个应用服务器的这个请求是合法的并给予放行;
		2. gitlab 这时候拿到 code 后再次向 github 登录服务器发起请求 token
		3. gitlab 拿到 token 后, 向 github 服务提交的请求头带上该 token, token 验证通过就可以拿到用户信息;
*/

const htmlIndex = `<html><body>
<a href="/GithubLogin">Log in with Github</a>
</body></html>
`

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize", // github oauth 流程的 oauth 网关, 所有 oauth 提供商都有一个 url 网关
	TokenURL: "https://github.com/login/oauth/access_token",
}

var githubOauthConfig = &oauth2.Config{
	ClientID:     "6155ffa6f3921c4e67d5", // 客户端 id, 代表哪个应用的请求验证
	ClientSecret: "f1429e0462a5fae29b290f893d0e9a3910324023",
	RedirectURL:  "http://localhost:8000/GithubCallback",
	Scopes:       []string{"user", "repo"}, // 表示申请的权限范围
	Endpoint:     endpoint,
}

const oauthStateString = "random" //

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/GithubLogin", handleGithubLogin)
	http.HandleFunc("/GithubCallback", handleGithubCallback)

	http.ListenAndServe(":8000", nil)
}

func handleMain(w http.ResponseWriter, t *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL(oauthStateString)
	/*
		response_type: 授权类型, 有 code 简化模式 密码模式 客户端模式
		redirect_uri: 重定向 uri, 验证以后的回调地址,一般用来接收返回的code, 以及做下一步处理
		state: 表示客户端的当前状态, 可以指定任意值, 认证服务器会原封不动的返回这个值,作为安全校验
	*/
	// https://github.com/login/oauth/authorize?client_id=6155ffa6f3921c4e67d5&redirect_uri=http%3A%2F%2Flocalhost%3A8000%2FGithubCallback&response_type=code&scope=user+repo&state=random
	fmt.Println("url:", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGithubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected %s, got %s", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println("state:", state)        // random
	fmt.Println("url:", r.URL.String()) //  /GithubCallback?code=999674a0436c5b8423bb&state=random
	code := r.FormValue("code")
	fmt.Println("code:", code) // 999674a0436c5b8423bb

	token, err := githubOauthConfig.Exchange(oauth2.NoContext, code)
	/*
		&{
		  AccessToken:5502b8ccd0f3bb5c502d3ecfb2498a93ad54ca10
		  TokenType:bearer
		  RefreshToken: Expiry:0001-01-01 00:00:00 +0000 UTC
		  raw:map[
			  access_token:[5502b8ccd0f3bb5c502d3ecfb2498a93ad54ca10]
		  	  scope:[repo,user]
			  token_type:[bearer]
			]
		}
	*/
	fmt.Printf("token:%+v", token)
	if err != nil {
		fmt.Printf("code exchange failed with %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}

	response, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "%s\n", contents)
}

/*
	{"login":"majunjun12345","id":30212741,"node_id":"MDQ6VXNlcjMwMjEyNzQx",
	"avatar_url":"https://avatars3.githubusercontent.com/u/30212741?v=4","gravatar_id":"",
	"url":"https://api.github.com/users/majunjun12345","html_url":"https://github.com/majunjun12345",
	"followers_url":"https://api.github.com/users/majunjun12345/followers",
	"following_url":"https://api.github.com/users/majunjun12345/following{/other_user}",
	"gists_url":"https://api.github.com/users/majunjun12345/gists{/gist_id}",
	"starred_url":"https://api.github.com/users/majunjun12345/starred{/owner}{/repo}",
	"subscriptions_url":"https://api.github.com/users/majunjun12345/subscriptions",
	"organizations_url":"https://api.github.com/users/majunjun12345/orgs",
	"repos_url":"https://api.github.com/users/majunjun12345/repos",
	"events_url":"https://api.github.com/users/majunjun12345/events{/privacy}",
	"received_events_url":"https://api.github.com/users/majunjun12345/received_events",
	"type":"User","site_admin":false,"name":"mamenlgi","company":"myself","blog":"www.mamengli.com",
	"location":null,"email":"598959594@qq.com","hireable":null,"bio":"a fighter","public_repos":29,
	"public_gists":0,"followers":0,"following":3,"created_at":"2017-07-16T12:58:59Z",
	"updated_at":"2019-06-17T14:04:27Z","private_gists":0,"total_private_repos":32,
	"owned_private_repos":32,"disk_usage":676949,"collaborators":0,"two_factor_authentication":false,
	"plan":{"name":"free","space":976562499,"collaborators":0,"private_repos":10000}}
*/
