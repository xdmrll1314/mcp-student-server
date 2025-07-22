package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// 定义学生结构体
type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Grade   string `json:"grade"`
	Section string `json:"section"`
}

// 模拟数据库
var students = []Student{
	{ID: "1", Name: "张三", Age: 18, Grade: "高三", Section: "A班"},
	{ID: "2", Name: "李四", Age: 17, Grade: "高二", Section: "B班"},
	{ID: "3", Name: "王五", Age: 16, Grade: "高一", Section: "C班"},
	{ID: "4", Name: "赵六", Age: 17, Grade: "高二", Section: "A班"},
	{ID: "5", Name: "孙七", Age: 16, Grade: "高一", Section: "B班"},
}

// 定义获取学生列表的参数结构体
type GetStudentListArgs struct {
	Section string `json:"section"`
}

// 定义获取学生信息的参数结构体
type GetStudentInfoArgs struct {
	StudentID string `json:"student_id"`
}

func main() {
	// 创建MCP服务器
	s := server.NewMCPServer(
		"Student Information Server",
		"1.0.0",
		server.WithToolCapabilities(false),
	)

	// 添加获取学生列表工具
	getStudentListTool := mcp.NewTool("get_student_list",
		mcp.WithDescription("获取班级所有学生的列表"),
		mcp.WithString("section",
			mcp.Description("班级名称，可选参数。如果提供，则只返回该班级的学生"),
		),
	)

	// 添加获取学生列表工具处理函数
	s.AddTool(getStudentListTool, mcp.NewTypedToolHandler(
		func(ctx context.Context, request mcp.CallToolRequest, args GetStudentListArgs) (*mcp.CallToolResult, error) {
			// 如果指定了班级，则过滤学生
			var result []Student
			if args.Section != "" {
				for _, student := range students {
					if student.Section == args.Section {
						result = append(result, student)
					}
				}
			} else {
				result = students
			}

			// 格式化输出
			output := fmt.Sprintf("找到 %d 名学生:\n", len(result))
			for _, student := range result {
				output += fmt.Sprintf("ID: %s, 姓名: %s, 年龄: %d, 年级: %s, 班级: %s\n",
					student.ID, student.Name, student.Age, student.Grade, student.Section)
			}

			return mcp.NewToolResultText(output), nil
		},
	))

	// 添加获取单个学生信息工具
	getStudentInfoTool := mcp.NewTool("get_student_info",
		mcp.WithDescription("根据学生ID获取单个学生的详细信息"),
		mcp.WithString("student_id",
			mcp.Required(),
			mcp.Description("学生的唯一标识ID"),
		),
	)

	// 添加获取单个学生信息工具处理函数
	s.AddTool(getStudentInfoTool, mcp.NewTypedToolHandler(
		func(ctx context.Context, request mcp.CallToolRequest, args GetStudentInfoArgs) (*mcp.CallToolResult, error) {
			// 查找学生
			for _, student := range students {
				if student.ID == args.StudentID {
					// 格式化输出
					output := fmt.Sprintf("学生详细信息:\nID: %s\n姓名: %s\n年龄: %d\n年级: %s\n班级: %s",
						student.ID, student.Name, student.Age, student.Grade, student.Section)
					return mcp.NewToolResultText(output), nil
				}
			}

			return mcp.NewToolResultError(fmt.Sprintf("未找到ID为 %s 的学生", args.StudentID)), nil
		},
	))

	// 设置MCP服务器处理HTTP请求
	httpServer := server.NewStreamableHTTPServer(s)

	// 添加一个简单的健康检查端点
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("MCP Student Server is running!"))
	})

	// 启动HTTP服务器
	port := 8080
	log.Printf("MCP Student Server starting on port %d...", port)
	if err := httpServer.Start(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
