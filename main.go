package main

import (
	"encoding/json"
	"fmt"
	"memory_kv_db/define"
	"memory_kv_db/lib"
	"net/http"
	"strconv"
	"sync"
)

// todo 数据持久化，定时器
// 定义内存数据库
type InMemoryDB struct {
	data  map[string]string
	queue *lib.PriorityQueue
	mutex sync.RWMutex
}

// 初始化内存数据库
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data:  make(map[string]string),
		queue: lib.NewPriorityQueue(lib.SortableQueue{}),
	}
}

// 根据 key 查询值
func (db *InMemoryDB) GetHandler(w http.ResponseWriter, r *http.Request) {
	var req define.GetReq
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db.mutex.RLock()
	defer db.mutex.RUnlock()

	value, exists := db.data[req.Key]
	if !exists {
		http.Error(w, fmt.Sprintf("key %+v not found", req.Key), http.StatusInternalServerError)
		return
	}

	resp := define.GetResp{
		Key:   req.Key,
		Value: value,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// 写入 key 和 value 到内存数据库中
func (db *InMemoryDB) SetHandler(w http.ResponseWriter, r *http.Request) {
	var req define.SetReq
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db.mutex.RLock()
	defer db.mutex.RUnlock()

	db.data[req.Key] = strconv.Itoa(req.Value)
	db.queue.Push(lib.DataItem{Key: req.Key, Value: req.Value})

	resp := define.SetResp{
		Message: fmt.Sprintf("key: %+v,value: %+v, write success", req.Key, req.Value),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// 取内存数据库中第rank大的元素
func (db *InMemoryDB) GetRankKVHandler(w http.ResponseWriter, r *http.Request) {
	var req define.GetRankReq
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if req.Rank > lib.SortQueueLimit {
		http.Error(w, fmt.Sprintf("rank %+v exceed sortqueue limit: %+v", req.Rank, lib.SortQueueLimit), http.StatusInternalServerError)
		return
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()

	//每次请求第rank的元素都先把堆里元素重新排序
	db.queue.Sort()
	data := db.queue.GetRank(req.Rank)

	resp := define.GetRankResp{
		Key:   data.Key,
		Value: data.Value,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (db *InMemoryDB) GetSortedQueueHandler(w http.ResponseWriter, r *http.Request) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	//每次请求第rank的元素都先把堆里元素重新排序
	db.queue.Sort()
	sortedQueue := db.queue.GetSortedQueue()

	resp := define.GetSortedQueueResp{
		sortedQueue,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// 创建一个新的内存数据库
	db := NewInMemoryDB()

	// 注册HTTP处理函数
	http.HandleFunc("/get", db.GetHandler)
	http.HandleFunc("/set", db.SetHandler)
	http.HandleFunc("/get_rank", db.GetRankKVHandler)
	http.HandleFunc("/get_sorted_queue", db.GetSortedQueueHandler)

	// 启动HTTP服务器
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":7069", nil)
}
