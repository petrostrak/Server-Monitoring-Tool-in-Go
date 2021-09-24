package main

import "log"

type job struct {
	HostServiceID int
}

func (j *job) Run() {
	repo.ScheduledCheck(j.HostServiceID)
}

func startMonitoring() {
	if preferenceMap["monitoring_live"] == "1" {
		// trigger a message to broadcast to all clients
		// that app is starting to monitor
		data := make(map[string]string)
		data["message"] = "Monitoring is starting..."
		if err := app.WsClient.Trigger("public-channel", "app-starting", data); err != nil {
			log.Println(err)
		}

		// get all of the services that we want to monitor
		servicesToMonitor, err := repo.DB.GetServicesToMonitor()
		if err != nil {
			log.Println(err)
		}

		log.Println("length of services to monitor", len(servicesToMonitor))

		// range through the services

		// get the schedule unit and number

		// create a job

		// save the id of the job so we can start and stop it

		// broadcast over websockets the fact that se service
		// is scheduled

		// end range
	}
}
