package define

// 定义 get 请求的请求体结构体
type GetReq struct {
	Key string `json:"key"`
}

// 定义 get 请求的响应结构体
type GetResp struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// 定义 set 请求的请求体结构体
type SetReq struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

// 定义 set 请求的响应结构体
type SetResp struct {
	Message string `json:"message"`
}

// 定义 get_rank 请求的请求体结构体
type GetRankReq struct {
	Rank int `json:"rank"`
}

// 定义 get_rank 请求的响应结构体
type GetRankResp struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

// 定义 get_rank 请求的响应结构体
type GetSortedQueueResp struct {
	SortedQueue []int `json:"sorted_queue"`
}
