# 使用django orm数据库管理

* django是一个python的web框架， django对与中小型站点有着良好的支撑。
* django的orm与mysql,postgre等关系型数据库结合非常好，可以基于版本地自动管理数据库的表结构变更历史， 并自动将表结构的变更同步到各个环境的数据库中， 避免了手动维护表结构变更语句历史sql文件的重复工作和人为错误。

## django管理数据库的优点

1. 无需手动维护表结构变更历史文件， 无需关注表结构变更文件执行顺序， django自动根据修改记录将表结构修改同步到数据库。
2. 修改表结构时，无需手动操作数据库；更改python数据库模型定义即可，熟悉之后，比直接操作数据库更直观。
3. 如果有必要可以提供django的admin后台，将各个表数据管理使用web起来。

## django管理数据库的缺点
1. django会自动创建 django*, auth*的数据表， 目前版本是10个， 大多是为了django基本运行的表，影响不大。
2. django支持外键，如果使用外键，脱离django之后使用可能存在麻烦。

## 基本操作
1. 修改python数据库model。
   edit file in gadmin/models
2. 执行makemigrations, 生成migrations文件记录。
   python manage.py makemigrations
3. 执行migrate，将表结构变动同步到数据库。
   python manage.py migrate

python manage.py runserver 0.0.0.0:8188