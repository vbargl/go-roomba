package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/vbargl/go-roomba/pkg/roombapass"
)

func GetPasswordCloud() {
	flags := flag.NewFlagSet("get-password-cloud", flag.ExitOnError)
	accUsername := flags.String("username", "", "iRobot account username")
	accPassword := flags.String("password", "", "iRobot account password")
	apikey := flags.String("apikey", "", "Gigya API key (optional)")
	filter := flags.String("filter", "", "filter robots either by uid or name (optional)")
	flags.Parse(os.Args[2:])

	if len(*accUsername) == 0 || len(*accPassword) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify iRobot account username and password!\n")
		flags.Usage()
		os.Exit(1)
	}

	robots, err := roombapass.GetPasswordFromCloud(
		roombapass.WithCredentials(*accUsername, *accPassword),
		roombapass.WithAPIKey(*apikey),
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}

	// if filter is set, print only the password of the robot with the given UID or name
	if *filter != "" {
		for _, robot := range robots {
			if robot.UID == *filter || robot.Name == *filter {
				fmt.Fprintf(os.Stderr, "iRobot %s password is:", robot.Name)
				fmt.Println(robot.Password)
				return
			}
		}

		fmt.Fprintf(os.Stderr, "Robot with UID or name %s not found!\n", *filter)
		os.Exit(1)
	}

	// otherwise print all robot passwords
	for _, robot := range robots {
		fmt.Printf("iRobot %s password is: %s\n", robot.Name, robot.Password)
	}
}
