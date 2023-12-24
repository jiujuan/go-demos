## main.go
main.go 代码来自官方 example 代码：[hertz-examples/parameter/query/main.go](https://github.com/cloudwego/hertz-examples/blob/main/parameter/query/main.go)

- `DefaultQuery()` ： 获 URL 参数值，如果返回值为空，还可以赋默认值
- `Query()`：获 URL 参数值

请求 1 ：
```
GET http://127.0.0.1:8080/welcome
```
返回内容 :
```
Hello: 'firstname: Guest' 'lastname: ', favorite food: []
```
参数 `firstname` 没有返回默认 `Guest`` ，`lastname` 返回空。

请求 2 ：
```
GET http://127.0.0.1:8080//welcome?firstname=Jane&lastname=Doe&food=apple&food=fish
```
返回内容 :
```
Hello: 'firstname: Jane' 'lastname: Doe', favorite food: [apple fish]
```
所以的参数内容都获取到了

请求 3：
```
GET http://127.0.0.1:8080/hello/tom?number=12
```
返回内容 ：
```
{
    "name": "tom",
    "num": "12"
}
```