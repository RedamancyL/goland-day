# 1. 设置go语言代码快速补全
- 1. ctrl+shift+p
- 2. 输入snippets,选择 Configure User Snippets
- 3. 选择go
- 4. 输入json代码
```
	// }

	"pln": {
		"prefix": "pln",                 // 输入内容
		"body": "fmt.Println($0)",       // 输出格式
		"description": "fmt.Println()"   // 介绍信息
	},
	
	"plf": {
		"prefix": "plf",
		"body": "fmt.Printf(\"$0\\n\",)",
		"description": "fmt.Printf()"
	}
```

注意：
- 1. vscode 国内不支持go module 


