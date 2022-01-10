package process

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/Jeffail/gabs/v2"
)

func Process(ctx context.Context, project, file string, mapping_proposal map[string]string, quiet bool) {
	client, err := pubsub.NewClient(ctx, project)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var mapping = make(map[string]map[string]string)

	// restructure the mapping to something that is easier to use later on for matching messages
	for value, rule := range mapping_proposal {
		rules := strings.Split(rule, ":")
		if len(rules) != 2 {
			panic(fmt.Sprintf("Incorrect format for filter specified: '%s=%s'. Should be VALUE=field.name:pubsub/topic", value, rule))
		}
		if _, present := mapping[value]; !present {
			mapping[value] = make(map[string]string)
		}
		mapping[value][rules[0]] = rules[1]
	}

	var scanner *bufio.Scanner

	if file == "-" {
		// read from stdin
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		// read in from file
		file, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		// we expect single JSON object per line
		m, err := gabs.ParseJSON(scanner.Bytes())
		if err != nil {
			panic(err)
		}

		// check messages against mapping to figure out where (if anywhere) it should be sent
		for route_value, route_map := range mapping {
			for field, routing_info := range route_map {
				if value, ok := m.Path(field).Data().(string); ok {
					if value == route_value || route_value == "*" {
						// this message should be sent to routing_info topic
						topic := client.Topic(routing_info)
						res := topic.Publish(ctx, &pubsub.Message{Data: []byte(m.String())})
						msgID, err := res.Get(ctx)
						if err != nil {
							panic(err)
						}
						if !quiet {
							fmt.Println("Message sent: ", msgID)
						}
					}
				}
			}
		}
	}
}

// eof
