package main

type Config struct {
	ModuleList  []string `json:"module_list"`  // go module的目录
	GlobalFiles []string `json:"global_files"` // 需要全局更新的文件
	IgnoreFiles []string `json:"ignore_files"` // 忽略的文件
}
