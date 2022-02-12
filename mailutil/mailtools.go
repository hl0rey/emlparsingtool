package mailutil

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"main/parsemail"
	"main/printutil"
	"strconv"
)

func Printheader(email parsemail.Email) {

	for k, _ := range email.Header {
		printutil.Prettyprint(k+": "+email.Header.Get(k), printutil.INFO)
	}
	printutil.Prettyprint("Found "+strconv.Itoa(len(email.Header))+" headers.", printutil.GOOD)

}

func Printattach(email parsemail.Email) {

	printutil.Prettyprint("Found "+strconv.Itoa(len(email.Attachments))+" Attachments（附件）", printutil.GOOD)
	for _, v := range email.Attachments {
		printutil.Prettyprint(fmt.Sprintf("Filename: %s\tContentType: %s", v.Filename, v.ContentType), printutil.INFO)
	}
}

func Getmailcontent(email parsemail.Email) ([]byte, []byte) {

	notext := false
	nohtml := false
	var htmlbody []byte
	var textbody []byte
	var err error
	if email.HTMLBody != "" {
		htmlbody, err = base64.StdEncoding.DecodeString(email.HTMLBody)
		if err != nil {
			printutil.Prettyprint("Decode HTMLBody faild.", printutil.ERROR)
			err = nil
		}
	} else {
		nohtml = true
		//textbody=[]byte{}
	}
	if email.TextBody != "" {
		textbody, err = base64.StdEncoding.DecodeString(email.TextBody)
		if err != nil {
			printutil.Prettyprint("Decode TextBody faild.", printutil.ERROR)
			err = nil
		}

	} else {
		notext = true
		//textbody=[]byte{}
	}

	if notext && nohtml {
		printutil.Prettyprint("No content found.", printutil.ERROR)
	}
	return textbody, htmlbody
}

func Getattchdoc(email parsemail.Email) {

	var filedata []byte
	var err error
	for _, v := range email.Attachments {
		printutil.Prettyprint("Export "+v.Filename, printutil.INFO)
		filedata, err = ioutil.ReadAll(v.Data)
		if err != nil {
			printutil.Prettyprint("Error,when output "+v.Filename+",Message: "+err.Error(), printutil.ERROR)
			continue
		} else {
			ioutil.WriteFile(v.Filename, filedata, 0644)
			printutil.Prettyprint("Success "+v.Filename, printutil.GOOD)
		}

	}

}

//Received X-Originating-IP Date From To Cc Bcc Subject X-Mailer
func Printbrief(email parsemail.Email) {
	//解析出关键信息
	//subject
	printutil.Prettyprint("Subject（主题）: "+parsemail.DecodeMimeSentenceGBK(email.Subject), printutil.GOOD)
	//From
	printutil.Prettyprint("From（发件人）: "+email.Header.Get("From"), printutil.GOOD)
	//To
	printutil.Prettyprint("To（收件人）: "+email.Header.Get("To"), printutil.GOOD)
	//Cc
	printutil.Prettyprint("Cc（抄送）: "+email.Header.Get("Cc"), printutil.GOOD)
	//Bcc
	printutil.Prettyprint("Bcc（密送）: "+email.Header.Get("Bcc"), printutil.GOOD)
	//Date
	printutil.Prettyprint("Date（时间）: "+email.Header.Get("Date"), printutil.GOOD)
	//X-Originating-IP
	printutil.Prettyprint("X-Originating-IP（发件人IP）: "+email.Header.Get("Date"), printutil.GOOD)
	//X-Mailer
	printutil.Prettyprint("X-Mailer（发信客户端）: "+email.Header.Get("Received"), printutil.GOOD)
	//Received
	printutil.Prettyprint("Received（收发信息摘要）: "+email.Header.Get("Received"), printutil.GOOD)

	//var fromlist string
	//for k,v:=range email.From{
	//	Frommaxindex:=len(email.From)-1
	//	if k==Frommaxindex {
	//		fromlist=fromlist+fmt.Sprintf("%s(%s)",v.Name,v.Address)
	//		break
	//	}
	//	fromlist=fromlist+fmt.Sprintf("%s(%s),",v.Name,v.Address)
	//}
	//printutil.Prettyprint("From: "+fromlist,printutil.GOOD)

}

//rfc2047
func Decodeheader() {

}
