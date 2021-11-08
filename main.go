// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		log.Printf("group id: %s", event.Source.GroupID)

		if event.Type == linebot.EventTypeMemberJoined {
			//welcome(event.ReplyToken)
		} else if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if !strings.HasPrefix(message.Text, "?") &&
					!strings.HasSuffix(message.Text, "?") &&
					!strings.HasPrefix(message.Text, "？") &&
					!strings.HasSuffix(message.Text, "？") {
					return
				}

				msg := strings.TrimPrefix(message.Text, "?")
				msg = strings.TrimPrefix(msg, "？")
				msg = strings.TrimSuffix(msg, "?")
				msg = strings.TrimSuffix(msg, "？")

				switch msg {
				case "指令", "常用指令":
					reply(event.ReplyToken, message.Text,
						linebot.NewMessageAction("交車", "交車？"),
						linebot.NewMessageAction("外觀", "外觀相關？"),
						linebot.NewMessageAction("內裝", "內裝相關？"),
						linebot.NewMessageAction("設定", "設定相關？"),
						linebot.NewMessageAction("行車記錄器", "行車記錄器？"),
						linebot.NewMessageAction("輪胎", "輪胎相關？"),
						linebot.NewMessageAction("防跳石網", "防跳石網？"),
						linebot.NewMessageAction("鑰匙皮套", "鑰匙皮套？"),
						linebot.NewMessageAction("遮陽簾", "遮陽簾？"),
						linebot.NewMessageAction("隔熱紙", "隔熱紙？"),
						linebot.NewURIAction("更多 (尚未更新)", "https://drive.google.com/file/d/1AM7PAPzMhp9BT3qKEP0lMdDKEx62kRSW/view"),
					)
				case "交車":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("交車前驗車檢查項目2.0", "https://drive.google.com/file/d/19N6rUajn42eWfQJMikYySdcyGEvr1QR4/view"),
						linebot.NewURIAction("正式交車檢查2.0", "https://drive.google.com/file/d/1S-XPfwNZFWAwQzc3gZbOj3vM8dP7TXR4/view"),
					)
				case "族貼", "族框":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("KAMIQ TW CLUB 族貼 | 族框", "https://kamiq.club/article?sid=350&aid=434"),
					)
				case "外觀相關":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("水簾洞與導水條", "https://kamiq.club/article?sid=324&aid=378"),
						linebot.NewURIAction("雨刷異音、會跳、立雨刷與更換", "https://kamiq.club/article?sid=324&aid=379"),
						linebot.NewURIAction("後視鏡指甲倒插問題", "https://kamiq.club/article?sid=324&aid=381"),
						linebot.NewURIAction("第三煞車燈水氣無法散去", "https://kamiq.club/article?sid=324&aid=382"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=324"),
					)
				case "內裝相關":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("車室異音-低速篇", "https://kamiq.club/article?sid=325&aid=383"),
						linebot.NewURIAction("車室異音-高速篇", "https://kamiq.club/article?sid=325&aid=384"),
						linebot.NewURIAction("車室靜音工程(含DIY與外廠安裝)", "https://kamiq.club/article?sid=325&aid=386"),
						linebot.NewURIAction("冷氣濾網更換", "https://kamiq.club/article?sid=325&aid=400"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=325"),
					)
				case "設定相關":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("搖控器啟閉車窗示範", "https://kamiq.club/article?sid=328&aid=375"),
						linebot.NewURIAction("Keyless鑰匙沒電手動開門方式", "https://kamiq.club/article?sid=328&aid=376"),
						linebot.NewURIAction("怠速引擎熄火判斷條件", "https://kamiq.club/article?sid=328&aid=377"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=328"),
					)
				case "行車記錄器":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("Garmin 66WD", "https://kamiq.club/article?sid=329&aid=394"),
						linebot.NewURIAction("HP S970 (電子後視鏡)", "https://kamiq.club/article?sid=329&aid=395"),
						linebot.NewURIAction("DOD RX900", "https://kamiq.club/article?sid=329&aid=503"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=328"),
					)
				case "輪胎相關":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("胎壓偵測器", "https://kamiq.club/article?sid=334&aid=388"),
						linebot.NewURIAction("有線/無線打氣機", "https://kamiq.club/article?sid=334&aid=456"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=334"),
					)
				case "防跳石網":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("防跳石網安裝", "https://kamiq.club/article?sid=335&aid=402"),
						linebot.NewURIAction("防跳石網配色參考", "https://kamiq.club/article?sid=335&aid=404"),
						linebot.NewURIAction("怠速引擎熄火判斷條件", "https://kamiq.club/article?sid=328&aid=377"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=335"),
					)
				case "鑰匙皮套":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("Hsu's 頑皮革", "https://kamiq.club/article?sid=338&aid=416"),
						linebot.NewURIAction("Story Leather", "https://kamiq.club/article?sid=338&aid=425"),
						linebot.NewURIAction("賽頓精品手工皮件", "https://kamiq.club/article?sid=338&aid=423"),
						linebot.NewURIAction("JC手作客製皮套", "https://kamiq.club/article?sid=338&aid=424"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=338"),
					)
				case "遮陽簾":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("晴天遮陽簾", "https://kamiq.club/article?sid=330&aid=438"),
						linebot.NewURIAction("徐府遮陽簾", "https://kamiq.club/article?sid=330&aid=439"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=330"),
					)
				case "隔熱紙":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("GAMA-E系列", "https://kamiq.club/article?sid=330&aid=403"),
						linebot.NewURIAction("Carlife X系列", "https://kamiq.club/article?sid=330&aid=417"),
						linebot.NewURIAction("3M極黑系列", "https://kamiq.club/article?sid=330&aid=499"),
						linebot.NewURIAction("Solar Gard 舒熱佳鑽石 LX 系列", "https://kamiq.club/article?sid=330&aid=500"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=330"),
					)
				case "避光墊":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("愛力美奈納碳避光墊", "https://kamiq.club/article?sid=333&aid=427"),
						linebot.NewURIAction("BSM專用仿麂皮避光墊", "https://kamiq.club/article?sid=333&aid=428"),
					)
				case "晴雨窗":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("晴雨窗", "https://kamiq.club/article?sid=333&aid=445"),
					)
				case "腳踏墊":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("3D卡固", "https://kamiq.club/article?sid=331&aid=406"),
						linebot.NewURIAction("Škoda原廠腳踏墊", "https://kamiq.club/article?sid=331&aid=420"),
						linebot.NewURIAction("台中裕峰訂製款", "https://kamiq.club/article?sid=331&aid=419"),
					)
				case "後車廂墊":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("後車廂墊", "https://kamiq.club/article?sid=331&aid=430"),
						linebot.NewURIAction("3M安美", "https://kamiq.club/article?sid=331&aid=418"),
					)
				case "車側飾板", "後廂護板":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("車側飾板|後廂護板", "https://kamiq.club/article?sid=336"),
					)
				case "其他週邊":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("旋轉杯架", "https://kamiq.club/article?sid=350&aid=436"),
						linebot.NewURIAction("後行李箱連動燈", "https://kamiq.club/article?sid=350&aid=448"),
						linebot.NewURIAction("光控燈膜", "https://kamiq.club/article?sid=350&aid=446"),
						linebot.NewURIAction("KAMIQ TW CLUB 族貼 | 族框", "https://kamiq.club/article?sid=350&aid=434"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=350"),
					)
				case "原廠週邊":
					reply(event.ReplyToken, message.Text,
						linebot.NewURIAction("原廠週邊價格表", "https://kamiq.club/article?sid=349&aid=407"),
						linebot.NewURIAction("原廠檔泥板", "https://kamiq.club/article?sid=349&aid=444"),
						linebot.NewURIAction("原廠門側垃圾桶", "https://kamiq.club/article?sid=349&aid=442"),
						linebot.NewURIAction("原廠多媒體底座", "https://kamiq.club/article?sid=349&aid=443"),
						linebot.NewURIAction("更多", "https://kamiq.club/article?sid=349"),
					)
				}

			}
		}
	}
}

func welcome(replyToken string) {
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(`
		新朋友您好!!
		歡迎加入KamiQ車主群，
		
		有任何問題都可以先爬文，不懂再詢問唷～
		==========
		KamiQ 車友群公開資訊
		https://kamiq.club/
		==========
		群組內的訊息很多，記得關提醒，尤其上班日
		記事本內有很多資料可以先爬文一下
		
		PS.KamiQ機器人會不定期進化成長!!
	`)).Do(); err != nil {
		log.Print(err)
	}
}

func reply(replyToken, msg string, actions ...linebot.TemplateAction) {
	contents := make([]*linebot.BubbleContainer, 0, len(actions))
	for _, act := range actions {
		btnComponent := make([]linebot.FlexComponent, 0)
		btnComponent = append(btnComponent, &linebot.ButtonComponent{
			Type:   linebot.FlexComponentTypeButton,
			Action: act,
			Style:  linebot.FlexButtonStyleTypePrimary,
			//Color:  "#8E8E8E",
		})
		contents = append(contents, &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Hero: &linebot.ImageComponent{
				Type:        linebot.FlexComponentTypeImage,
				URL:         "https://kamiq.club/upload/36/favicon_images/c1a630ef-c78f-43cc-b95e-0619f3f4da4d.jpg",
				Size:        linebot.FlexImageSizeTypeFull,
				AspectRatio: linebot.FlexImageAspectRatioType20to13,
				AspectMode:  linebot.FlexImageAspectModeTypeFit,
			},
			Footer: &linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeButton,
				Layout:   linebot.FlexBoxLayoutTypeVertical,
				Contents: btnComponent,
			},
		})
	}

	if _, err := bot.ReplyMessage(replyToken, linebot.NewFlexMessage(msg, &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: contents,
	})).Do(); err != nil {
		log.Println(err)
	}
}
