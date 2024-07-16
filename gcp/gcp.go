package gcp

import (
	"log"
	"os/exec"
	"scheduler/utils"
)

func Action(p utils.Params, b utils.KBody) (string, error) {
	switch p.Action {
	case "up":
		return utils.CmdUp(b)
	case "down":
		return utils.CmdDown(b)
	}
	return "", nil
}

func CallGCP(cmd_l string) (string, error) {
	log.Println(cmd_l)

	cmd := exec.Command("/bin/sh", "-c", cmd_l)

	_, err := cmd.Output()

	if err != nil {
		log.Println("Something is wrong", cmd_l)
	} else {
		log.Println("Change applied: ", cmd_l)
	}

	return cmd_l, err
}
