package apicalls

import "strings"

func openApp(message string) string {
	message = "Read the user's request below, which asks for assistance with an application. Your task is to identify the name of the application they are referring to, based solely on the provided input. Respond with the name of the application if it can be clearly identified from the request. If the request does not contain enough information to determine the application's name confidently, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the application's name or '0', directly addressing the user's request. \n User request" + message

	res, err := CallMistralAPI(message)
	if err != nil {
		return "something happened, try again later"
	}

	answer := strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	answer = strings.SplitN(answer[0], ".", 2)
	return "001" + answer[0]
}

func closeApp() string {
	return "close app"
}

func openFile() string {
	return "open file"
}

func closeFile() string {
	return "close file"
}

func createFile() string {
	return "create file"
}

func renameFile() string {
	return "rename file"
}

func createFolder() string {
	return "create folder"
}

func renameFolder() string {
	return "rename folder"
}

func deleteFile() string {
	return "delete file"
}

func deleteFolder() string {
	return "delete folder"
}

func takeScreenshot() string {
	return "take screenshot"
}

func playMusic() string {
	return "play music"
}

func PauseMusic() string {
	return "pause music"
}

func OpenURL() string {
	return "open url"
}

func sendEmail() string {
	return "send email"
}

func readPDF() string {
	return "read pdf"
}

func setTimer() string {
	return "set timer"
}

func setAlert() string {
	return "set alarm"
}

func startChrono() string {
	return "start chrono"
}

func textToSpeachClipboard() string {
	return "text to speach clipboard"
}
