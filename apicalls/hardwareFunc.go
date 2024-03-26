package apicalls

import "strings"

func volumeUp(message string) string {
	message = "Read the user's request below, which asks for assistance with turning up the volume. Your task is to identify the percentage or number the user wants to increase the volume by, based solely on the provided input. Respond with the number or the percentage the user wants to increase the volume by, if it can be clearly identified from the request. If the request does not contain enough information to determine the number confidently, respond with '0'. For example is the user wants to increase volume by 5 respond with '5', if they wants to increase the volume by 10% respond with '%10'. If there is a percentage it is crucial that you place the percentage symbol '%' before the number. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the number the user wants to increase the volume by or '0', directly addressing the user's request. \n User request" + message

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
	return "101" + answer[0]
}

func volumeDown(message string) string {
	message = "Read the user's request below, which asks for assistance with turning down the volume. Your task is to identify the percentage or number the user wants to decrease the volume by, based solely on the provided input. Respond with the number or the percentage the user wants to decrease the volume by, if it can be clearly identified from the request. If the request does not contain enough information to determine the number confidently, respond with '0'. For example is the user wants to decrease volume by 5 respond with '5', if they wants to decrease the volume by 10% respond with '%10'. If there is a percentage it is crucial that you place the percentage symbol '%' before the number. It is crucial that you provide a single response based on the user's current request, without offering additional examples, further explanations, or multiple attempts. Your response must consist of only the number the user wants to decrease the volume by or '0', directly addressing the user's request. \n User request" + message

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
	return "102" + answer[0]
}
