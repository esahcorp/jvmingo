1.1 指针

* `&` 取址苻
* `*` 取值

1.2 flag 包处理命令行信息

1.3 struct 结构体与类相似

1.4 go install & go build

* `go install` 生成的可执行文件在 `$GOPATH/bin` 目录下  
* `go build` 生成的可执行文件在当前目录下

1.5 方法接受者什么时候用指针

    可能修改结构体 -> newXX 函数返回是指针 -> 方法定义为指针
    即方法接受者与构造函数返回值保持一致

2.1 defer 语句用于延迟执行代码，适合关闭资源等操作

3.1 Big Endian 大端是字节序的一种，低地址位存放数据的高位，与 Little Endian 相反

3.2 recover() 与 panic()

* `panic()` 终止当前程序，开始执行 defer 栈
* `recover()` 只能在 defer 中调用，捕获 panic 造成的异常，阻止程序崩溃

3.3 字节码结构

``` c
struct ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count - 1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
};
```

3.4 字段结构

``` c
struct field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attribute_count;
    attribute_info attributes[attributes_count];
};
```

3.5 常量结构

``` c
struct cp_info {
    u1 tag;
    u1 info;
};
```

3.4 magic number 是格式的一种惯例，非 Java 独有