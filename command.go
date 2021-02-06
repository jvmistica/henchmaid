package main

import "strings"

// CheckCommand checks if the command is valid and returns the appropriate response
func (i *Items) CheckCommand(cmd string) (string, error) {
	var msg string
	var params []string

	if len(strings.Split(cmd, " ")) > 1 {
		params = strings.Split(cmd, " ")[1:]
	}

	switch cmd {
	case "/start":
		msg = startMsg
	case "/listitems":
		if len(params) == 0 {
			msg = i.ListItems("")
			if msg == "" {
				msg = noItems
			}
		} else {
			msg = i.ListItems(strings.Join(params[0:], " "))
		}
	case "/showitem":
		msg, err := i.ShowItem(params)
		if err != nil {
			return "", err
		}
		return msg, nil
	case "/additem":
		msg, err := i.AddItem(params)
		if err != nil {
			return "", err
		}
		return msg, nil
	case "/updateitem":
		msg, err := i.UpdateItem(params)
		if err != nil {
			return "", err
		}
		return msg, nil
	case "/deleteitem":
		if len(params) == 0 {
			msg = deleteChoose
		} else {
			msg = i.DeleteItem(params)
		}
	default:
		msg = invalidMsg
	}
	return msg, nil
}
