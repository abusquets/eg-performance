package eg

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func RunApi() {

	targeter := vegeta.NewStaticTargeter(
		vegeta.Target{
			Method: "GET",
			URL:    "https://events-guard.fly.dev/api/hello/hello",
		},
	)

	// Set up the attack: Request of requests and duration
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second

	// Execute the attack
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics

	for res := range attacker.Attack(targeter, rate, duration, "Test GET") {
		metrics.Add(res)
	}

	// End and print the results
	metrics.Close()

	// Show the latencies
	fmt.Println("Latències:")
	fmt.Printf("Mean: %s\n", metrics.Latencies.Mean)
	fmt.Printf("95th Percentile: %s\n", metrics.Latencies.P95)
	fmt.Printf("Max: %s\n", metrics.Latencies.Max)

	// Other useful statistics
	fmt.Printf("Peticions totals: %d\n", metrics.Requests)
	fmt.Printf("Èxit: %.2f%%\n", metrics.Success*100)
	fmt.Printf("Codi d'estat: %+v\n", metrics.StatusCodes)
}
