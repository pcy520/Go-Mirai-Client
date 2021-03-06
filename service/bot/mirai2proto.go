package bot

import (
	"strconv"

	"github.com/Mrs4s/MiraiGo/message"
	"github.com/ProtobufBot/Go-Mirai-Client/proto_gen/onebot"
)

func MiraiMsgToProtoMsg(messageChain []message.IMessageElement) []*onebot.Message {
	msgList := make([]*onebot.Message, 0)
	for _, element := range messageChain {
		switch elem := element.(type) {
		case *message.TextElement:
			msgList = append(msgList, MiraiTextToProtoText(elem))
		case *message.AtElement:
			msgList = append(msgList, MiraiAtToProtoAt(elem))
		case *message.ImageElement:
			msgList = append(msgList, MiraiImageToProtoImage(elem))
		case *message.FaceElement:
			msgList = append(msgList, MiraiFaceToProtoFace(elem))
		case *message.VoiceElement:
			msgList = append(msgList, MiraiVoiceToProtoVoice(elem))
		}
	}
	return msgList
}

func MiraiTextToProtoText(elem *message.TextElement) *onebot.Message {
	return &onebot.Message{
		Type: "text",
		Data: map[string]string{
			"text": elem.Content,
		},
	}
}

func MiraiImageToProtoImage(elem *message.ImageElement) *onebot.Message {
	return &onebot.Message{
		Type: "image",
		Data: map[string]string{
			"file": elem.Url,
			"url":  elem.Url,
		},
	}
}

func MiraiAtToProtoAt(elem *message.AtElement) *onebot.Message {
	return &onebot.Message{
		Type: "at",
		Data: map[string]string{
			"qq": func() string {
				if elem.Target == 0 {
					return "all"
				}
				return strconv.FormatInt(elem.Target, 10)
			}(),
		},
	}
}

func MiraiFaceToProtoFace(elem *message.FaceElement) *onebot.Message {
	return &onebot.Message{
		Type: "face",
		Data: map[string]string{
			"id": strconv.Itoa(int(elem.Index)),
		},
	}
}

func MiraiVoiceToProtoVoice(elem *message.VoiceElement) *onebot.Message {
	return &onebot.Message{
		Type: "record",
		Data: map[string]string{
			"file": elem.Url,
			"url":  elem.Url,
		},
	}
}
