package infra

import (
	"reflect"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

/*
	启动器
*/
const (
	keyProps = "_conf"
)

//基础资源上下文结构体
type StarterContext map[string]interface{}

func (s StarterContext) Props() kvs.ConfigSource {
	p := s[keyProps]
	if p == nil {
		panic("配置还没有初始化")
	}
	return p.(kvs.ConfigSource)
}

func (s StarterContext) SetProps(conf kvs.ConfigSource) {
	s[keyProps] = conf
}

//资源启动器，每个应用少不了依赖其他资源，比如数据库，缓存，消息中间件等等服务
//启动器实现类，不需要实现所有方法，只需要实现对应的阶段方法即可，可以嵌入@BaseStarter
//通过实现资源启动器接口和资源启动注册器，友好的管理这些资源的初始化、安装、启动和停止。
//Starter对象注册器，所有需要在系统启动时需要实例化和运行的逻辑，都可以实现此接口
//注意只有Start方法才能被阻塞，如果是阻塞Start()，同时StartBlocking()要返回true
type Starter interface {
	//1.系统启动, 初始化一些基础资源
	Init(StarterContext)
	//2.资源安装
	Setup(StarterContext)
	//3.启动基础资源
	Start(StarterContext)
	//说明该资源启动器开始启动服务时，是否会阻塞
	StartBlocking() bool
	//4.资源停止和销毁  优雅的关闭
	Stop(StarterContext)
	PriorityGroup() PriorityGroup
	Priority() int
}

type PriorityGroup int

const (
	SystemGroup         PriorityGroup = 30
	BasicResourcesGroup PriorityGroup = 20
	AppGroup            PriorityGroup = 10

	INT_MAX          = int(^uint(0) >> 1)
	DEFAULT_PRIORITY = 10000
)

//基础空启动器实现,为了方便资源启动器的代码实现
type BaseStarter struct {
}

func (b *BaseStarter) Init(ctx StarterContext)      {}
func (b *BaseStarter) Setup(ctx StarterContext)     {}
func (b *BaseStarter) Start(ctx StarterContext)     {}
func (b *BaseStarter) StartBlocking() bool          { return false }
func (b *BaseStarter) Stop(ctx StarterContext)      {}
func (b *BaseStarter) PriorityGroup() PriorityGroup { return BasicResourcesGroup }
func (b *BaseStarter) Priority() int                { return DEFAULT_PRIORITY }

//启动器注册器
//不需要外部构造，全局只有一个
type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}

//返回所有启动器
func (r *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, r.nonBlockingStarters...)
	starters = append(starters, r.blockingStarters...)
	return starters
}

//注册启动器
func (r *starterRegister) Register(s Starter) {
	if s.StartBlocking() {
		r.blockingStarters = append(r.blockingStarters, s)
	} else {
		r.nonBlockingStarters = append(r.nonBlockingStarters, s)
	}
	typ := reflect.TypeOf(s)
	logrus.Infof("Register starter: %s", typ.String())
}

//方便外部调用
var StarterRegister = new(starterRegister)


type Starters []Starter

func (s Starters) Len() int {
	return len(s)
}

func (s Starters) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Starters) Less(i, j int) bool {
	return s[i].PriorityGroup() > s[j].PriorityGroup() && s[i].Priority() > s[j].Priority()
}

//提供外部注册
func Register(s Starter) {
	StarterRegister.Register(s)
}

func SortStarters() {
	sort.Sort(Starters(StarterRegister.AllStarters()))
}

func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}
