# go 的错误处理

### 官方给出的四种错误检查方式

Whatever you do, always check your errors!

- 哨兵错误(sentinel errors)检查
    `var ErrNoRows = errors.New("sql: no rows in result set")`
