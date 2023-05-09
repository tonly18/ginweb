# ginweb
 gin web是在gin框架基础上，进行了二次封装，方便开发。可直接用于项目开发。

##### 特点:
##### 1、使用endless支持热重启。
##### 2、使用模板模式，二次封装了http请求处理方法(controller)。
##### 3、使用logrus处理日志。
##### 4、支持链路追踪。
##### 5、使用viper处理配置文件。


var signals = [...]string{
    // 这里省略N行。。。。
 
    /** 找到此位置添加如下 */
    16: "SIGUSR1",
    17: "SIGUSR2",
    18: "SIGTSTP",
 
}
 
/** 兼容windows start */
func Kill(...interface{}) {
    return;
}
const (
    SIGUSR1 = Signal(0x10)
    SIGUSR2 = Signal(0x11)
    SIGTSTP = Signal(0x12)
)
