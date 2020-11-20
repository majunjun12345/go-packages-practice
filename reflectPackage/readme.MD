### 0、 好文
  [Go语言反射reflect](https://www.cnblogs.com/itbsl/p/10551880.html)

反射是指在程序运行期对程序本身进行访问和修改的能力。
程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，
并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

#### 一、 通过反射获取类型信息
![](https://img2020.cnblogs.com/blog/720430/202006/720430-20200626110101552-361035114.png)

  使用 reflect.TypeOf() 函数可以获得任意值的类型对象 reflect.Type;
  ```go
  type Type interface {
    Name() string                        // 类型变量名称
    Kind() Kind                          // 基础类型
    Elem() Type                          // 通过反射获取指针指向的元素类型
    
    // 通过反射获取结构体的成员类型, 当类型不是结构体或索引超界时发生宕机
    NumField() int                                                  // 返回结构体成员字段数量
    Field(i int) StructField                                        // 根据索引，返回索引对应的结构体字段的信息
    FieldByName(name string) (StructField, bool)                    // 根据给定字符串返回字符串对应的结构体字段的信息, 没有找到时 bool 返回 false
    FieldByIndex(index []int) StructField                           // 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息, 没有找到时返回零值
    FieldByNameFunc(match func(string) bool) (StructField, bool)    // 根据匹配函数匹配需要的字段

      type StructField struct {
        Name string         // 字段名
        Type      Type      // 字段反射的类型对象
        Tag       StructTag // 字段的结构体标签
        Offset    uintptr   // 字段在结构体中的相对偏移
        Index     []int     // Type.FieldByIndex中的返回的索引值
        Anonymous bool      // 是否为匿名字段
        PkgPath string      // 字段路径
      }

    Align() int
    FieldAlign() int
    Method(int) Method                   // 类型的方法
    MethodByName(string) (Method, bool)  // 是否存在该方法
    NumMethod() int                      // 类型方法数量
    
    PkgPath() string
    Size() uintptr
    String() string
    
    Implements(u Type) bool
    AssignableTo(u Type) bool
    ConvertibleTo(u Type) bool
    Comparable() bool
    Bits() int
    ChanDir() ChanDir
    IsVariadic() bool
    
    
    
    
    
    In(i int) Type
    Key() Type
    Len() int
    
    NumIn() int
    NumOut() int
    Out(i int) Type
    common() *rtype
    uncommon() *uncommonType
  }
  ```



#### 一、 通过反射获取值信息
![](https://img2018.cnblogs.com/blog/720430/201903/720430-20190318141014788-581422387.png)

- 当值不是结构体或索引超界时发生宕机
NumField() int 返回结构体成员字段数量。
Field(i int) Value 根据索引，返回索引对应的结构体成员字段的反射值对象。
FieldByName(name string) Value 根据给定字符串返回字符串对应的结构体字段, 没有找到时返回零值。
FieldByIndex(index []int) Value 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的值, 没有找到时返回零值。
FieldByNameFunc(match func(string) bool) Value 根据匹配函数匹配需要的字段, 找到时返回零值。

- 判断反射值的空和有效性
IsNil() bool 返回值是否为 nil。如果值类型不是通道（channel）、函数、接口、map、指针或 切片时发生 panic，类似于语言层的v== nil操作。
IsValid() bool 判断值是否有效。 当值本身非法时，返回 false，例如 reflect Value不包含任何值，值为 nil 等。

- 通过反射修改变量值
  - 条件：
    可被寻址
    被导出

  Elem() Value 取值指向的元素值，类似于语言层*操作。当值类型不是指针或接口时发生宕机，空指针时返回 nil 的 Value
  Addr() Value 对可寻址的值返回其地址，类似于语言层&操作。当值不可寻址时发生宕机
  CanAddr() bool 表示值是否可寻址
  CanSet() bool 返回值能否被修改。要求值可寻址且是导出的字段

- 通过反射调用方法