# 分销商登录

| api                  | 方法 | 含义       | 参数                  | 成功响应              | 失败响应                     |
| -------------------- | ---- | ---------- | --------------------- | --------------------- | ---------------------------- |
| /fenxiao/user/signin | POST | 分销商登录 | 参数 Phone, Password  | {Ok:true , Data:User} | {Ok:false,Data:"用户不存在"} |
| /fenxiao/user/signup | POST | 分销商注册 | 参数 Phone,Password , | {Ok:true,Data:User}   | {Ok:false,Data:"错误原因"}   |
