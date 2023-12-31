# It's a common frame using Gin and Gorm


# 这是一个使用GIN和GORM的基础框架，提供了用户的登录注册等


  ``git clone https://github.com/mahaonan001/GIN_GORM.git``

  
  ``go mod init GIN_GORM``

  
  ``go mod tidy``


## Second


### install mysql and set a database called TEST and change your mysql password or change the file in config,the application.yml file.There is a datesourse password change to your own password.


### 电脑需要安装MySQL并且有一个TEST命名的数据库，默认密码是root，可以在项目中的config/application.yml文件中更改password为自己的mysqlpassword


  ``go build -o server main.go(mac or linux)``

  
  ``go build -o server.exe main.go(Windows)``

  
## Finally


  ``./server(mac or linux)``

  
  ``.\server.exe(windows)``
