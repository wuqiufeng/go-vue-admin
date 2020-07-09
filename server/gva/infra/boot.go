package infra

import (
	"reflect"

	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

// 应用启动管理器
type BootApplication struct {
	IsTest         bool
	conf           kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{conf: conf, starterContext: StarterContext{}}
	//b.starterContext[keyProps] = conf
	b.starterContext.SetProps(conf)
	return b
}

func (b *BootApplication) Start() {
	//1. 初始化starter
	b.init()
	//2. 安装starter
	b.setup()
	//3. 启动starter
	b.start()
}

func (b *BootApplication) init() {
	logrus.Info("Initializing starters...")
	for _, starter := range GetStarters() {
		typ := reflect.TypeOf(starter)
		logrus.Infof("Initializing: PriorityGroup=%d,Priority=%d,type=%s", starter.PriorityGroup(), starter.Priority(), typ.String())
		starter.Init(b.starterContext)
	}
}
func (b *BootApplication) setup() {
	logrus.Info("Setup starters...")
	for _, starter := range GetStarters() {
		typ := reflect.TypeOf(starter)
		logrus.Info("Setup: ", typ.String())
		starter.Setup(b.starterContext)
	}
}

//程序开始运行，开始接受调用
func (b *BootApplication) start() {
	logrus.Info("Starting starters...")
	for i, starter := range GetStarters() {
		typ := reflect.TypeOf(starter)
		logrus.Debug("Starting: ", typ.String())
		if starter.StartBlocking() {
			// 如果是最后一个可阻塞的，直接启动并阻塞
			if i+1 == len(GetStarters()) {
				starter.Start(b.starterContext)
			} else {
				// 如果不是，使用goroutine来异步启动
				//防止阻塞后面的starter
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}

//程序开始运行，开始接受调用
func (b *BootApplication) Stop() {
	logrus.Info("Stopping starters...")
	for _, starter := range GetStarters() {

		typ := reflect.TypeOf(starter)
		logrus.Debug("Stopping: ", typ.String())
		starter.Stop(b.starterContext)
	}
}
