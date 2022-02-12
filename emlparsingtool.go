package main

import (
	"flag"
	"io/ioutil"
	"main/mailutil"
	"main/parsemail"
	"main/printutil"
	"os"
)

func main() {

	var filename string
	var attachdoc bool
	var mailcontent bool
	var allin bool
	var allheader bool

	flag.StringVar(&filename, "f", "", "the eml filename.")
	flag.BoolVar(&attachdoc, "a", false, "if you want export attach file.")
	flag.BoolVar(&mailcontent, "c", false, "if you want export mail content.")
	flag.BoolVar(&allin, "all", false, "export attach file and email content")
	flag.BoolVar(&allheader, "ah", false, "print all header.")
	flag.Parse()

	if filename == "" {
		flag.PrintDefaults()
		return
	}

	mailfp, err := os.Open(filename)
	if err != nil {
		println(err.Error())
		return
	}
	email, err := parsemail.Parse(mailfp)
	if err != nil {
		println(err.Error())
		return
	}
	if mailcontent || allin {
		textbody, htmlbody := mailutil.Getmailcontent(email)
		subject_decode := parsemail.DecodeMimeSentenceGBK(email.Subject)
		printutil.Prettyprint("Subject（主题）: "+subject_decode, printutil.GOOD)
		if len(textbody) != 0 {
			//textfile := strconv.Itoa(int(time.Now().Unix()))
			textfile := subject_decode
			err := ioutil.WriteFile(textfile+".txt", textbody, 0644)
			if err != nil {
				println(err.Error())
			} else {
				printutil.Prettyprint("text content wirte to "+textfile+".txt", printutil.INFO)
			}
		}
		if len(htmlbody) != 0 {
			//htmlfile := strconv.Itoa(int(time.Now().Unix()))
			htmlfile := subject_decode
			err := ioutil.WriteFile(htmlfile+".html", htmlbody, 0644)
			if err != nil {
				println(err.Error())
			} else {
				printutil.Prettyprint("text content wirte to "+htmlfile+".html", printutil.INFO)
			}
		}

	} else if attachdoc || allin {
		mailutil.Getattchdoc(email)
	} else if allheader {
		mailutil.Printheader(email)
	} else {
		//one param,print brief
		mailutil.Printbrief(email)
		//if have attachments then print
		if len(email.Attachments) != 0 {
			mailutil.Printattach(email)
		}
	}

}
