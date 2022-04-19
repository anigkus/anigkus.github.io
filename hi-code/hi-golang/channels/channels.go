/*
Copyright 2022 The https://github.com/anigkus Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func Main() {
	fmt.Println("channels...")

	rand.Seed(time.Now().Unix())

	// Create a buffered channel to manage the employee vs project load.
	projects := make(chan string, 10)

	// Launch 5 goroutines to handle the projects.
	goRoutine.Add(5)
	for i := 1; i <= 5; i++ {
		go employee(projects, i)
	}

	for j := 1; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	// Close the channel so the goroutines will quit
	close(projects)
	goRoutine.Wait()
}

func employee(projects chan string, employee int) {
	defer goRoutine.Done()
	for {
		// Wait for project to be assigned.
		project, result := <-projects

		if !result {
			// This means the channel is empty and closed.
			fmt.Printf("Employee : %d : Exit\n", employee)
			return
		}

		fmt.Printf("Employee : %d : Started   %s\n", employee, project)

		// Randomly wait to simulate work time.
		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// Display time to wait
		fmt.Println("\nTime to sleep", sleep, "ms")

		// Display project completed by employee.
		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}

}
