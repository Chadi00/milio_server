package apicalls

import "log"

func SoftwareCall(message string) string {
	req := SoftwarePrompt + message
	output := "I can't help you with that, sorry!"

	res, err := CallOpenAIAPI(req, 5)
	if err != nil {
		output = "I can't help you with that now, please try later."
	}

	answer := output

	if res.Choices != nil && len(res.Choices) > 0 {
		answer = res.Choices[0].Message.Content
	} else {
		output = "No choices returned from the API."
		return output
	}

	switch answer[0:2] {
	case "01":
		output := openApp(message)
		return output
	case "02":
		output := closeApp(message)
		return output
	case "03":
		output := openFile(message)
		return output
	case "04":
		output := closeFile(message)
		return output
	case "05":
		output := createFile(message)
		return output
	case "06":
		output := renameFile(message)
		return output
	case "07":
		output := deleteFile(message)
		return output
	case "08":
		output := createFolder(message)
		return output
	case "09":
		output := renameFolder(message)
		return output
	case "10":
		output := deleteFolder(message)
		return output
	case "11":
		output := takeScreenshot()
		return output
	case "12":
		output := playMusic(message)
		return output
	case "13":
		output := PauseMusic()
		return output
	case "14":
		output := OpenURL(message)
		return output
	case "15":
		output := sendEmail(message)
		return output
	case "16":
		output := readPDF(message)
		return output
	case "17":
		output := setTimer(message)
		return output
	case "18":
		output := startChrono()
		return output
	}

	return output + " " + answer
}

func HardwareCall(message string) string {
	req := HardwarePrompt + message
	output := "I can't help you with that, sorry!"

	res, err := CallOpenAIAPI(req, 5)
	if err != nil {
		output = "I can't help you with that now, please try later."
	}

	answer := output

	if res.Choices != nil && len(res.Choices) > 0 {
		answer = res.Choices[0].Message.Content
	} else {
		output = "No choices returned from the API."
		return output
	}

	switch answer[0:2] {
	case "01":
		output := volumeUp(message)
		return output
	case "02":
		output := volumeDown(message)
		return output
	case "03":
		return "103"
	case "04":
		return "104"
	case "05":
		return "105"
	case "06":
		return "106"
	case "07":
		return "107"
	}

	return output + " " + answer
}

func DomoCall(message string) string {
	return "Domo"
}

func SearchCall(message string) string {
	message = SearchPrompt + message

	res, err := CallOpenAIAPI(message, 50)
	if err != nil {
		return "something happened, try again later"
	}

	if len(res.Choices) == 0 {
		log.Println("res.Choices len == 0")
		return "something happened, try again later"
	}

	answer := res.Choices[0].Message.Content
	return "3-" + answer
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
