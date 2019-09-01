package scheduler

import (
	"github.com/bamzi/jobrunner"
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
)

func Run() {
	runner := []string{
		data.SchedulerSpot,
		data.SchedulerEatWhat,
		data.SchedulerWeiboFollow,
		data.SchedulerWeiboLoginChecker,
		data.SchedulerWeiboMessage,
		data.SchedulerWeiboTopic,
		data.SchedulerWeiboGroupSender,
		data.SchedulerTaobaoToWeibo,
	}
	for x := range runner {
		if data.GetConfig(runner[x]) != "" {
			log.Info().Str("runner", runner[x]).Msg("run init")
		}
	}
	jobrunner.Start(32, 0)

	jobrunner.Schedule("@every 5m", HotSpotRunner{})

	jobrunner.Schedule("0 0 11,17 * * *", EatWhat{do: Choose})
	jobrunner.Schedule("0 0 12,18 * * *", EatWhat{do: Result})

	jobrunner.Schedule("@every 10m", WeiboTopicRunner{})

	jobrunner.Schedule("@every 3m", WeiboLoginChecker{1})
	jobrunner.Schedule("@every 10m", GroupSender{})
	jobrunner.Schedule("0 0/5 * * * *", TaobaoToWeibo{})
}
