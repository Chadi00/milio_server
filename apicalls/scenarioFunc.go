package apicalls

func SoftwareCall(message string) string {
	req := SoftwarePrompt + message
	output := "I can't help you with that, sorry!"

	res, err := CallMistralAPI(req)
	if err != nil {
		output = "I can't help you with that now, please try later."
	}

	answer := res.Choices[0].Message.Content

	switch answer[0:2] {
	case "01":
		output := openApp(message)
		return output
	case "02":
		output := closeApp()
		return output
	case "03":
		output := openFile()
		return output
	case "04":
		output := closeFile()
		return output
	case "05":
		output := createFile()
		return output
	case "06":
		output := renameFile()
		return output
	case "07":
		output := createFolder()
		return output
	case "08":
		output := renameFolder()
		return output
	case "09":
		output := deleteFile()
		return output
	case "10":
		output := deleteFolder()
		return output
	case "11":
		output := takeScreenshot()
		return output
	case "12":
		output := playMusic()
		return output
	case "13":
		output := PauseMusic()
		return output
	case "14":
		output := OpenURL()
		return output
	case "15":
		output := sendEmail()
		return output
	case "16":
		output := readPDF()
		return output
	case "17":
		output := setTimer()
		return output
	case "18":
		output := setAlert()
		return output
	case "19":
		output := startChrono()
		return output
	case "20":
		output := textToSpeachClipboard()
		return output
	}

	return output + " " + answer
}

func HardwareCall(message string) string {
	return "Hardware"
}

func DomoCall(message string) string {
	return "Domo"
}

func SearchCall(message string) string {
	return "Search"
}

func LogicCall(message string) string {
	return "Logic"
}

func CreativeCall(message string) string {
	return "Creative"
}

func CsCall(message string) string {
	return "Computer science"
}
