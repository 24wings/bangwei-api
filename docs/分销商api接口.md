# 分销商登录

| api                           | 方法 | 含义           | 参数                  | 成功响应              | 失败响应                     |
| ----------------------------- | ---- | -------------- | --------------------- | --------------------- | ---------------------------- |
| /fenxiao/user/signin          | POST | 分销商登录     | 参数 Phone, Password  | {Ok:true , Data:User} | {Ok:false,Data:"用户不存在"} |
| /fenxiao/user/signup          | POST | 分销商注册     | 参数 Phone,Password , | {Ok:true,Data:User}   | {Ok:false,Data:"错误原因"}   |
| /fenxiao/user/user-auth-code/ | Get  | 用户短信验证码 | 参数 Phone            | {Ok:true,Data:""}     | {Ok:false,Data:"错误原因"}   |
| /static/citys.json            | GET  | 省份和城市列表 | 无                    | {provinces:[]}        |                              |
| fenxiao/user/user             |
