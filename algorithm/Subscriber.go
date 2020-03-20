package algorithm

import "fmt"

type Channel struct {
	userName string
}
type Channels struct{
	chann map[string]map[string]Channel
}
//注册订阅者
func (ch *Channels) Reg(appId string,userId string,userName string)  {
	c := new(Channel)
	c.userName = userName
	if ch.chann == nil {
		ch.chann = make(map[string]map[string]Channel)
	}
	if ch.chann[appId] == nil {
		ch.chann[appId] = make(map[string]Channel)
	}
	ch.chann[appId][userId] = *c
}
//取消订阅
func (ch *Channels) Remove(appId string,userId string)  {
	if ch.chann[appId] == nil {
		fmt.Println("该通道不存在")
	}else {
		delete(ch.chann[appId],userId)
	}
}
//发送消息
func (ch *Channels) Send(appId string,msg string)  {
	for k,v := range ch.chann[appId] {
		fmt.Println(appId,k,v.userName,msg)
	}
}