package scheduler

import (
	"github.com/rs/zerolog/log"
	"go-go-go/src/data"
	"go-go-go/src/utils"
)

type TaobaoToWeibo struct {
}

func (runner TaobaoToWeibo) Run() {
	defer utils.CommonRecover()
	if data.GetConfig("adasdasd") == "" {
		return
	}
	log.Info().Msg("start TaobaoToWeibo")
	weibo_follow()
}
