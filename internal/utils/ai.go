package utils

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/nepu-SidneyYu/blog-grpc/internal/logs"
	"github.com/nepu-SidneyYu/blog-grpc/internal/model"
	"go.uber.org/zap"
)

// // 存储到Response的中的提问消息和问答消息
// type Message struct {
// 	Role    string `json:"role"`
// 	Content string `json:"content"`
// }

// // 向Bot提出的Request，
// type request struct {
// 	Model       string     `json:"model"`
// 	Messages    []*Message `json:"messages"`
// 	ServiceCode string     `json:"serviceCode"`
// 	MaxTokens   int        `json:"maxTokens"`
// 	KeyType     string     `json:"keyType"`
// 	Tag         string     `json:"tag"`
// }

// // 从bot收到的应答
// type Response struct {
// 	Content string `json:"content"`
// 	Stop    bool   `json:"stop"`
// }

type botAiChat struct {
	// mp  map[string]*request
	// req *request
}

// func NewBotAI() AI {
// 	return &botAiChat{
// 		mp: make(map[string]*request),
// 		req: &request{
// 			Model:       "baidu-eb-instant",
// 			Messages:    make([]*Message, 0),
// 			ServiceCode: "SI_KG_PRODUCE",
// 			KeyType:     "Z120",
// 			MaxTokens:   2048,
// 			Tag:         "service",
// 		},
// 	}
// }

// func Botformatmap(mp map[string]*request, botsess string, req *request, opration string) (*request, bool) {
// 	var l sync.Mutex
// 	l.Lock()
// 	defer l.Unlock()
// 	if opration == "store" {
// 		mp[botsess] = req
// 		return nil, true
// 	} else if opration == "load" {
// 		s, ok := mp[botsess]
// 		return s, ok
// 	} else if opration == "delete" {
// 		delete(mp, botsess)
// 		return nil, true
// 	}
// 	return nil, false
// }
// func (b *botAiChat) CreateSession() (string, error) {
// 	str := utils.NewStringID()
// 	Botformatmap(b.mp, str, b.req, "store")
// 	return str, nil
// }
// func (b *botAiChat) CloseSession(sid string) {
// 	if _, ok := b.mp[sid]; !ok {
// 		log.Println("Session 不存在")
// 		return
// 	}
// 	Botformatmap(b.mp, sid, b.req, "delete")
// }

func Chat(req *model.Request, fn func(*model.Response)) error {
	url := "https://ai-platform-cloud-proxy.polymas.com/chat-llm/v1/completions/stream/multi"
	jsonvalue, _ := json.Marshal(req)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonvalue))
	if err != nil {
		logs.Error(context.Background(), "创建请求失败", zap.String("error", err.Error()))
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		logs.Error(context.Background(), "请求失败", zap.String("error", err.Error()))
		return err
	}
	go func() {
		defer resp.Body.Close()
		// 读取响应体数据

		reader := bufio.NewReader(resp.Body)
		var respe = new(model.Response)
		//count := 0
		var mess = new(model.Message)
		var result = new(model.Response)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				logs.Error(context.Background(), "读取响应失败", zap.String("error", err.Error()))
				return
			}
			line = bytes.TrimSpace(line)
			if len(line) == 0 {
				continue
			}
			//res := bytes.TrimPrefix(line, []byte("data:"))

			err = json.Unmarshal(line, &result)
			if err != nil {
				logs.Error(context.Background(), "解析响应失败", zap.String("error", err.Error()))
				return
			}
			respe.Content = result.Content
			respe.Stop = result.Stop
			mess.Content += result.Content
			fn(respe)
			if result.Stop {
				respe.Stop = true
				mess.Role = "assistant"
				break
			}
		}
	}()
	return nil
}
