package main

import (
	"fmt"

	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

type Registry interface {
	Register(string, string, string, string, int) error
	Deregister(string) error
}

type client struct {
	client *consulapi.Client
}

type Service interface {
	GetServices() (map[string]*consulapi.AgentService, error)
	GetService(string) (*consulapi.AgentService, *consulapi.QueryMeta, error)
	FilterService(string) (map[string]*consulapi.AgentService, error)
}

func NewConfig(addr string) *consulapi.Config {
	config := consulapi.DefaultConfig()
	if addr != "" {
		config.Address = addr
	}
	return config
}

func NewClient(addr string) (*client, error) {
	config := NewConfig(addr)

	c, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{client: c}, nil
}

func (c *client) Register(name, id, tags, ip string, port int) error {
	// 服务注册信息
	reg := &consulapi.AgentServiceRegistration{
		ID:      id,             // 唯一服务 ID
		Name:    name,           // 服务名
		Tags:    []string{tags}, // 标签，可以标识相同服务
		Port:    port,           // 端口号
		Address: ip,             // 所在节点 ip 地址
	}
	// 服务健康检查
	reg.Check = &consulapi.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", ip, port, "/health-check"), // http形式检查，检测默认/health路径
		Timeout:                        "5s",
		Interval:                       "5s",      // 每隔5s检查一次
		DeregisterCriticalServiceAfter: "40s",     // 服务40s不可达时，注销服务
		Status:                         "passing", // 默认服务状态正常
	}
	return c.client.Agent().ServiceRegister(reg)
}

func (c *client) Deregister(id string) error {
	return c.client.Agent().ServiceDeregister(id)
}

// 获取所有服务
func (c *client) GetServices() (map[string]*consulapi.AgentService, error) {
	return c.client.Agent().Services()
}

// 获取单个服务
func (c *client) GetService(id string) (*consulapi.AgentService, *consulapi.QueryMeta, error) {
	return c.client.Agent().Service(id, nil)
}

// 条件过滤服务
func (c *client) FilterService(filter string) (map[string]*consulapi.AgentService, error) {
	return c.client.Agent().ServicesWithFilter(filter)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("health check"))
}

func main() {
	client, err := NewClient("")
	if err != nil {
		fmt.Println("consul client error: ", err.Error())
	}
	err = client.Register("hello-service", "hello-service", "hello", "127.0.0.1", 81)
	if err != nil {
		fmt.Println("Register err:", err.Error())
	}

	err = client.Register("hello-world-service", "hello-world-service", "hello", "127.0.0.1", 81)
	if err != nil {
		fmt.Println("Register err2:", err.Error())
	}

	getService, _, err := client.GetService("hello-service")
	if err != nil {
		fmt.Println("get service error: ", err.Error())
	}
	fmt.Println("get service val: ", getService)

	http.HandleFunc("/health-check", healthCheck)
	err = http.ListenAndServe(":81", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}
