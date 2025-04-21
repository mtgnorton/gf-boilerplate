## 概述
使用`dao.User.Ctx(ctx)`这种方式调用默认是链式安全的,所以使用需要使用赋值的方式替换原有的模型对象
```
m := user.Where("status", g.Slice{1,2,3})
if vip {
// 查询条件通过赋值叠加
m = m.Where("money>=?", 1000000)
} else {
// 查询条件通过赋值叠加
m = m.Where("money<?",  1000000)
}
```
## Where
- Where 支持的参数为任意的 string/map/slice/struct/*struct 类型。在需要考虑索引的情况建议使用字符串的方式
```golang
Where("level=? OR money >=?", 1, 1000000)
Where(g.Map{"uid" : 1})
```
- WherePri 支持主键的查询条件
- Wheref 格式化条件字符串


## 查询结果
- ~~- All 用于查询并返回多条记录的列表/数组。很少使用~~
- ~~- One 用于查询并返回单条记录。很少使用~~ 
- Array 用于查询指定字段列的数据，返回数组。
```golang
// SELECT `name` FROM `user` WHERE `score`>60
Model("user").Fields("name").Where("score>?", 60).Array()
```
- Value 用于查询并返回一个字段值，往往需要结合 Fields 方法使用。
```golang
// SELECT `name` FROM `user` WHERE `uid`=1 LIMIT 1
Model("user").Fields("name").Where("uid", 1).Value()
```
- Count 用于查询并返回记录数。
```golang
// SELECT COUNT(1) FROM `user` WHERE `status` IN(1,2,3)
Model("user").Where("status", g.Slice{1,2,3}).Count()

```
- All和One方法使用 Scan方法替代
```golang
// 查询单条
var user *User
g.Model("user").Where("id", 1).Scan(&user)

// 查询多条
var users []*User
g.Model("user").Scan(&users)
```

## Insert
- Insert 如果写入的数据中存在主键或者唯一索引时，返回失败，否则写入一条新数据。
- InsertAndGetId 写入数据时并直接返回自增字段的 ID
- Replace 如果写入的数据中存在主键或者唯一索引时，会删除原有的记录，必定会写入一条新记录。
- Save 如果写入的数据中存在主键或者唯一索引时，更新原有数据，否则写入一条新数据
- OnDuplicate/OnDuplicateEx需要结合 Save 方法一起使用，用于指定 Save 方法的更新/不更新字段，参数为字符串、字符串数组、 Map
- Insert,Save 方法都支持批量操作,并支持通过Batch方法设置分批写入条数数量
    ```
    g.Model("user").Data(g.List{
    {"name": "john_1"},
    {"name": "john_2"},
    {"name": "john_3"},
    }).Batch(2).Insert()
    ```
## Update
- Update
```golang
// UPDATE `user` SET `name`='john guo' WHERE name='john'
g.Model("user").Data(g.Map{"name" : "john guo"}).Where("name", "john").Update()

```
- Increment/Decrement 自增/减
```golang
// UPDATE `article` SET `views`=`views`+10000 WHERE `id`=1
g.Model("article").Where("id", 1).Increment("views", 10000)
```

## Delete
```golang
g.Model("user").Where("uid", 10).Delete()

```

## Cache
