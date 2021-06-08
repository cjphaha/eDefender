package plugin

// Task 任务结构
type Task struct {
	Type   string   `json:"type"`
	Netloc string   `json:"netloc"`
	Target string   `json:"target"`
	Meta   TaskMeta `json:"meta"`
}

// TaskMeta 任务额外信息
type TaskMeta struct {
	System   string   `json:"system"`
	PathList []string `json:"pathlist"`
	FileList []string `json:"filelist"`
	PassList []string `json:"passlist"`
}
