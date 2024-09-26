package config

import "time"

type AppCfg struct {
	Server      ServerCfg     `mapstructure:"server" json:"server" yaml:"server"`                   //服务器配置
	Remote      RemoteCfg     `mapstructure:"remote" json:"remote" yaml:"remote"`                   //远程配置
	Logger      LogCfg        `mapstructure:"logger" json:"logger" yaml:"logger"`                   //log配置
	JWT         JWT           `mapstructure:"jwt" json:"jwt" yaml:"jwt"`                            //jwt配置
	DBCfg       DBCfg         `mapstructure:"dbcfg" json:"dbcfg" yaml:"dbcfg"`                      // 数据库配置
	Cache       CacheCfg      `mapstructure:"cache" json:"cache" yaml:"cache"`                      // 缓存
	Cors        CORS          `mapstructure:"cors" json:"cors" yaml:"cors"`                         //cors配置
	Extends     any           `mapstructure:"extend" json:"extend" yaml:"extend"`                   //扩展配置
	Gen         GenCfg        `mapstructure:"gen" json:"gen" yaml:"gen"`                            //是否可生成
	GrpcServer  GrpcServerCfg `mapstructure:"grpc-server" json:"grpc-server" yaml:"grpc-server"`    //grpc服务配置
	AccessLimit AccessLimit   `mapstructure:"access-limit" json:"access-limit" yaml:"access-limit"` //访问限制配置
}

type ServerCfg struct {
	Name         string `mapstructure:"name" json:"name" yaml:"name"`                            //appname
	RemoteEnable bool   `mapstructure:"remote-enable" json:"remote-enable" yaml:"remote-enable"` //是否开启远程配置
	Mode         string `mapstructure:"mode" json:"mode" yaml:"mode"`                            //模式
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                            //启动host
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                            //端口
	ReadTimeout  int    `mapstructure:"read-timeout" json:"read-timeout" yaml:"read-timeout"`    //读超时 单位秒
	WriteTimeout int    `mapstructure:"write-timeout" json:"write-timeout" yaml:"write-timeout"` //写超时 单位秒
	FSType       string `mapstructure:"fs-type" json:"fs-type" yaml:"fs-type"`                   //文件系统
	I18n         bool   `mapstructure:"i18n" json:"i18n" yaml:"i18n"`                            //是否开启多语言
	Lang         string `mapstructure:"lang" json:"lang" yaml:"lang"`                            //默认语言
	CloseWait    int    `mapstructure:"close-wait" json:"close-wait" yaml:"close-wait"`          //服务关闭等待 秒
}

type GrpcServerCfg struct {
	Enable bool   `mapstructure:"enable" json:"enable" yaml:"enable"` //启用Grpc服务
	Name   string `mapstructure:"name" json:"name" yaml:"name"`       //服务名，不设置默认为ServerName+"_grpc"
	Host   string `mapstructure:"host" json:"host" yaml:"host"`       //启动host
	Port   int    `mapstructure:"port" json:"port" yaml:"port"`       //端口
}

func (e *GrpcServerCfg) GetHost() string {
	if e.Host == "" {
		e.Host = "0.0.0.0"
	}
	return e.Host
}

func (e *GrpcServerCfg) GetPort() int {
	if e.Port < 1 {
		e.Port = 7789
	}
	return e.Port
}

func (e *ServerCfg) GetLang() string {
	if e.Lang == "" {
		e.Lang = "zh-CN"
	}
	return e.Lang
}

func (e *ServerCfg) GetHost() string {
	if e.Host == "" {
		e.Host = "0.0.0.0"
	}
	return e.Host
}

func (e *ServerCfg) GetPort() int {
	if e.Port < 1 {
		e.Port = 7788
	}
	return e.Port
}

func (e *ServerCfg) GetCloseWait() int {
	if e.CloseWait < 1 {
		e.CloseWait = 1
	}
	return e.CloseWait
}

func (e *ServerCfg) GetReadTimeout() int {
	if e.ReadTimeout < 1 {
		e.ReadTimeout = 20
	}
	return e.ReadTimeout
}

func (e *ServerCfg) GetWriteTimeout() int {
	if e.WriteTimeout < 1 {
		e.WriteTimeout = 20
	}
	return e.WriteTimeout
}

type RemoteCfg struct {
	Provider      string `mapstructure:"provider" json:"provider" yaml:"provider"`                   //提供方
	Endpoint      string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`                   //端点
	Path          string `mapstructure:"path" json:"path" yaml:"path"`                               //路径
	SecretKeyring string `mapstructure:"secret-keyring" json:"secret-keyring" yaml:"secret-keyring"` //安全
	ConfigType    string `mapstructure:"config-type" json:"config-type" yaml:"config-type"`          //配置类型
}

func (e *RemoteCfg) GetConfigType() string {
	if e.ConfigType == "" {
		e.ConfigType = "yaml"
	}
	return e.ConfigType
}

type Config struct {
	Enable      bool             `mapstructure:"enable" json:"enable" yaml:"enable"`
	Driver      string           `mapstructure:"driver" json:"driver" yaml:"driver"`
	Endpoints   []string         `mapstructure:"endpoints" json:"endpoints" yaml:"endpoints"`
	Scheme      string           `mapstructure:"scheme" json:"scheme" yaml:"scheme"`
	Timeout     time.Duration    `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
	Registers   []*RegisterNode  `mapstructure:"registers" json:"registers" yaml:"registers"`
	Discoveries []*DiscoveryNode `mapstructure:"discoveries" json:"discoveries" yaml:"discoveries"`
}

type RegisterNode struct {
	Namespace   string        `mapstructure:"namespace" json:"namespace" yaml:"namespace"`          //命名空间
	Id          string        `mapstructure:"id" json:"id" yaml:"id"`                               //服务id
	Name        string        `mapstructure:"name" json:"name" yaml:"name"`                         //服务名
	Addr        string        `mapstructure:"addr" json:"addr" yaml:"addr"`                         //服务地址
	Port        int           `mapstructure:"port" json:"port" yaml:"port"`                         //端口
	Protocol    string        `mapstructure:"protocol" json:"protocol" yaml:"protocol"`             //协议
	Weight      int           `mapstructure:"weight" json:"weight" yaml:"weight"`                   //权重
	Interval    time.Duration `mapstructure:"interval" json:"interval" yaml:"interval"`             //检测间隔
	Timeout     time.Duration `mapstructure:"timeout" json:"timeout" yaml:"timeout"`                //服务检测超时时间
	HealthCheck string        `mapstructure:"health-check" json:"health-check" yaml:"health-check"` //健康检查地址
	Tags        []string      `mapstructure:"tags" json:"tags" yaml:"tags"`                         //标签
	FailLimit   int           `mapstructure:"fail-limit" json:"fail-limit" yaml:"fail-limit"`       //失败次数限制，到达失败次数就会被禁用
}

type DiscoveryNode struct {
	Enable              bool   `mapstructure:"enable" json:"enable" yaml:"enable"`                                           //启用发现
	Namespace           string `mapstructure:"namespace" json:"namespace" yaml:"namespace"`                                  //命名空间
	Name                string `mapstructure:"name" json:"name" yaml:"name"`                                                 //服务名
	Tag                 string `mapstructure:"tag" json:"tag" yaml:"tag"`                                                    //标签
	SchedulingAlgorithm string `mapstructure:"scheduling-algorithm" json:"scheduling-algorithm" yaml:"scheduling-algorithm"` //调度算法
	FailLimit           int    `mapstructure:"fail-limit" json:"fail-limit" yaml:"fail-limit"`                               //已发现服务最大失败数
	RetryTime           int    `mapstructure:"retry-time" json:"retry-time" yaml:"retry-time"`                               //重试时间间隔 秒
}
