package notify

import (
	"github.com/ouqiang/gocron/internal/models"
	"github.com/ouqiang/gocron/internal/modules/logger"
	"github.com/ouqiang/gocron/internal/modules/utils"
	gw "github.com/xen0n/go-workwx"
	"strconv"
	"strings"
)

type Workwx struct{}

func (workwx *Workwx) Send(msg Message) {
	model := new(models.Setting)
	config, err := model.Workwx()
	logger.Debugf("%+v", config)
	if err != nil {
		logger.Error("#mail#从数据库获取mail配置失败", err)
		return
	}
	if config.CorpId == "" || config.Secret == "" || config.AgentId == "" {
		logger.Error("#config#企微配置无效")
		return
	}

	msg["content"] = parseNotifyTemplate(config.Template, msg)
	toUsers := workwx.getActiveMailUsers(config, msg)
	workwx.send(config, msg, toUsers)
}

func (workwx *Workwx) send(config models.WorkwxConfig, msg Message, toUsers []string) {
	logger.Infof("send workwx message, config:%v \n msg:%+v \n toUsers:%v", config, msg["content"], toUsers)
	agentId, err := strconv.ParseInt(config.AgentId, 10, 64)
	if err != nil {
		logger.Error(err)
		return
	}
	w := gw.New(config.CorpId)
	app := w.WithApp(config.Secret, agentId)
	recipient := gw.Recipient{UserIDs: toUsers}
	err = app.SendTextMessage(&recipient, msg["content"].(string), true)
	if err != nil {
		logger.Error(err)
	}
}

func (workwx *Workwx) getActiveMailUsers(config models.WorkwxConfig, msg Message) []string {
	taskReceiverIds := strings.Split(msg["task_receiver_id"].(string), ",")
	var users []string
	for _, v := range config.Users {
		if utils.InStringSlice(taskReceiverIds, strconv.Itoa(v.Id)) {
			users = append(users, v.UserId)
		}
	}
	return users
}
