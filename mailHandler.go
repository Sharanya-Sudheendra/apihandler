package handlers

import(
	"fmt"
	"github.com/sharanya-sudheendra/apihandler" 
)

type MailHandler struct {
	
}

func NewMailHandler() *MailHandler {
	return &MailHandler{}
}


func (mh *MailHandler) SendMail() {

	message := mh.NewMessage(sender, subject, body, recipient)

	resp, id, err := mh.SendMail()

	fmt.Printf("ID: %s Resp: %s\n", id, resp) 
}  

