package generater

import (
	"fmt"
	"strings"
)

func MarketingGenerater(subject, event, otherword string) string {
	s := strings.Join([]string{subject, event, "是怎么回事呢?", subject, "相信大家都很熟悉,但是", subject, event, "是怎么回事呢,下面就让小编带大家一起了解吧<br>", subject, event, ",其实就是", otherword, ",大家可能会很惊讶", subject, "怎么会", event, "呢?但事实就是这样,小编也感到非常惊讶<br>这就是关于", subject, event, "的事情了,大家有什么想法呢,欢迎在评论区告诉小编一起讨论哦!"}, "")
	fmt.Println(s)
	return s
}

func NewsSubtitle(one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen string) string {
	s := strings.Join([]string{"此事件发生后,", one, "高度重视,立即连夜召开", two, "会议,", three, "立即作出", four, ",要求组织", five, ",妥善处理", six, ",迅速查清", seven, "严肃追究", eight, ",立即组织开展", nine, ",进一步做好", ten, ",防止", eleven, ".目前,", twelve, "情绪稳定,", thirteen, "深感内疚,自责,痛心"}, "")
	fmt.Println(s)
	return s
}
