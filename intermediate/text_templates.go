package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {

	// tmpl := template.New("example")                                   // this returns template
	// tmpl, err := tmpl.Parse("Welcome, {{.name}}! How are you doing?") // this returns template as well as an error
	// previus line uses gopher assignment because err is new variable NOTE
	// make it in one line as following:
	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n") //
	// if err != nil {
	// 	panic(err)
	// }

	// you can use template.Must() to handle the errors. so we don't need the variable err
	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

	// // Define data for the welcome message template
	// data := map[string]any{
	// 	"name": "Ahmed",
	// }

	// tmpl.Execute(os.Stdout, data) // Excute() returns error

	// data["name"] = os.Stdin

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Define named templates for different types of
	templates := map[string]string{ // templates are strings now before parsing
		"welcome":      "Welcome, {{.name}}! W're glad you joined.",
		"notification": "{{.nm}}, you have new notification: {{.ntf}} ",
		"error":        "Opps! An error occured: {{.em}}",
	}

	// Parse and store templates NOTE:Remember when you intialize map use make() function
	parsedTemplates := make(map[string]*template.Template) // now it is parsed!
	// let's now fill the parsedTemplates map
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		//show menu
		fmt.Println("\nMenu:")
		fmt.Println("1. join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]any{"name": name} // name variable which urser alrady entered
		case "2":
			fmt.Println("Enter you notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]any{"nm": name, "ntf": notification}
		case "3":
			fmt.Println("Enter you error message:")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]any{"em": errorMessage}

		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice Please select a valid option.")
			continue

		}

		//render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}

	}

}
