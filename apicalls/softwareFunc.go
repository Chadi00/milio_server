package apicalls

import (
	"strings"
)

func openApp(message string) string {
	message = "Read the user's request below, which asks for assistance with an application. Your task is to identify the name of the application they are referring to, based solely on the provided input. Respond with the name of the application if it can be clearly identified from the request. If the request does not contain enough information to determine the application's name confidently, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the application's name or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "001" + answer[0]
}

func closeApp(message string) string {
	message = "Read the user's request below, which asks for assistance with an application. Your task is to identify the name of the application they are referring to, based solely on the provided input. Respond with the name of the application if it can be clearly identified from the request. If the request does not contain enough information to determine the application's name confidently, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the application's name or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "002" + answer[0]
}

func openFile(message string) string {
	message = "Read the user's request below, which asks for assistance with opening a file. Your task is to identify the name of the file and its type (e.g., txt, mp4, cvs) based solely on the provided input. Respond with the name of the file and its type, separated by a dash, like 'filename-txt' or 'videoname-mp4'. If the request does not contain enough information to confidently determine both the name of the file and its type, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the file name and type in the specified format or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "003" + answer[0]
}

func closeFile(message string) string {
	message = "Read the user's request below, which asks for assistance with closing a file. Your task is to identify the name of the file and its type (e.g., txt, mp4, cvs) based solely on the provided input. Respond with the name of the file and its type, separated by a dash, like 'filename-txt' or 'videoname-mp4'. If the request does not contain enough information to confidently determine both the name of the file and its type, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the file name and type in the specified format or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "004" + answer[0]
}

func createFile(message string) string {
	message = "Read the user's request below, which asks for assistance with creating a file. Your task is to identify the name of the file and its type (e.g., txt, mp4, cvs) based solely on the provided input. Respond with the name of the file and its type, separated by a dash, like 'filename-txt' or 'videoname-mp4'. If the request does not contain enough information to confidently determine both the name of the file and its type, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the file name and type in the specified format or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "005" + answer[0]
}

func renameFile(message string) string {
	message = "Read the user's request carefully, where they seek assistance with renaming a specific file. Your objective is to discern three key pieces of information from their input: the current name of the file, the file's extension (indicating its type, such as 'txt', 'mp4', 'csv'), and the new name the user wishes to assign to this file. You must then formulate your response by concatenating the current file name, its type, and the new name, each separated by a dash. The format for your response should strictly adhere to 'currentFilename-fileType-newFilename'. For instance, if the task is to rename 'doc.txt' to 'newdoc', your response should be 'doc-txt-newdoc'. Ensure that your response includes only the concatenated string in the specified format or '0'. The response '0' should be given if the user's request lacks sufficient details for you to confidently extract the current file name, its type, and the intended new name. It is essential to provide a singular, accurate response based on the user's initial request alone. Do not attempt to offer multiple solutions, additional explanations, or ask for further clarification. Your answer must either be the precise file name and type in the requested format or '0', directly addressing what the user has asked for. \nUser request:" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "006" + answer[0]
}

func deleteFile(message string) string {
	message = "Read the user's request below, which asks for assistance with deleting a file. Your task is to identify the name of the file and its type (e.g., txt, mp4, cvs) based solely on the provided input. Respond with the name of the file and its type, separated by a dash, like 'filename-txt' or 'videoname-mp4'. If the request does not contain enough information to confidently determine both the name of the file and its type, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the file name and type in the specified format or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "007" + answer[0]
}

func createFolder(message string) string {
	message = "Read the user's request below, which asks for assistance with creating a folder. Your task is to identify the name of the folder they want to create, based solely on the provided input. Respond with the name of the folder if it can be clearly identified from the request. If the request does not contain enough information to determine the folder's name confidently, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the folder's name or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "008" + answer[0]
}

func renameFolder(message string) string {
	message = "Read the user's request carefully, where they seek assistance with renaming a specific folder. Your objective is to discern two key pieces of information from their input: the current name of the folder and the new name the user wishes to assign to this folder. You must then formulate your response by concatenating the current folder name and the new name, each separated by a dash. The format for your response should strictly adhere to 'currentFoldername-newFoldername'. For instance, if the task is to rename 'currentFolder' to 'newFolder', your response should be 'currentFolder-newFolder'. Ensure that your response includes only the concatenated string in the specified format or '0'. The response '0' should be given if the user's request lacks sufficient details for you to confidently extract the current folder name and the intended new name. It is essential to provide a singular, accurate response based on the user's initial request alone. Do not attempt to offer multiple solutions, additional explanations, or ask for further clarification. Your answer must either be the precise folder name and new name in the requested format or '0', directly addressing what the user has asked for. \nUser request:" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "009" + answer[0]
}

func deleteFolder(message string) string {
	message = "Read the user's request below, which asks for assistance with deleting a folder. Your task is to identify the name of the folder they want to delete, based solely on the provided input. Respond with the name of the folder if it can be clearly identified from the request. If the request does not contain enough information to determine the folder's name confidently, respond with '0'. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the folder's name or '0', directly addressing the user's request. \n User request" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "010" + answer[0]
}

func takeScreenshot() string {
	return "011"
}

func playMusic(message string) string {
	message = "Read the user's request carefully, where they seek assistance with playing music. Your objective is to discern three key pieces of information from their input: the name of the music they want to play, the name of the playlist and the name of the music application (spotify or apple music). You must then formulate your response by concatenating the name of the music, the name of the playlist and the name of the music app, each separated by a dash. The format for your response should strictly adhere to 'nameOfMusic-nameOfPlayslist-musicApp'. For instance, if the task is to play 'hello' from playlist 'today' on 'Spotify', your response should be 'hello-today-spotify'. Ensure that your response includes only the concatenated string in the specified format or '0'. The response '0' should be given if the user's request lacks sufficient details for you to confidently extract the name of the music, the name of the playlist and name of the music app. It is essential to provide a singular, accurate response based on the user's initial request alone. Do not attempt to offer multiple solutions, additional explanations, or ask for further clarification. Your answer must either be the precise folder name and new name in the requested format or '0', directly addressing what the user has asked for. \nUser request:" + message

	res, err := CallOpenAIAPI(message, 10)
	if err != nil {
		return "something happened, try again later"
	}

	answer := []string{"0"}
	if res.Choices != nil && len(res.Choices) > 0 {
		answer = strings.SplitN(res.Choices[0].Message.Content, "\n", 2)
	} else {
		return "0"
	}

	answer = strings.SplitN(answer[0], ".", 2)
	return "012" + answer[0]
}

func PauseMusic() string {
	return "013"
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
