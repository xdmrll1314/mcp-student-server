# MCP 学生信息服务器

这是一个使用Go语言和mark3labs/mcp-go库实现的MCP（Model Control Protocol）服务器，通过HTTP SSE（Server-Sent Events）进行通信。

## 功能

该服务器提供以下功能：

1. 获取学生列表（可选按班级筛选）
2. 根据学生ID获取单个学生的详细信息

## 技术栈

- Go语言
- mark3labs/mcp-go库
- HTTP SSE通信

## 安装和运行

### 前提条件

- Go 1.16或更高版本

### 安装依赖

```bash
go mod tidy
```

### 运行服务器

```bash
go run main.go
```

服务器将在8080端口启动。

## API端点

- `/mcp` - MCP服务器的SSE端点
- `/health` - 健康检查端点

## 工具说明

### 1. get_student_list

获取学生列表，可选择按班级筛选。

**参数：**

- `section` (可选): 班级名称，如果提供，则只返回该班级的学生

**示例：**

```json
{
  "section": "A班"
}
```

### 2. get_student_info

根据学生ID获取单个学生的详细信息。

**参数：**

- `student_id` (必填): 学生的唯一标识ID

**示例：**

```json
{
  "student_id": "1"
}
```

## 示例数据

服务器内置了5名学生的示例数据：

1. 张三 - 高三A班
2. 李四 - 高二B班
3. 王五 - 高一C班
4. 赵六 - 高二A班
5. 孙七 - 高一B班
