package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhako0999/student-manager/internal/student"
)

func main() {
	store := student.NewStore("students.json")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		runAdd(store, os.Args[2:])
	case "list":
		runList(store)
	case "curve":
		runCurve(store, os.Args[2:])
    default:
		fmt.Println("Unknown command:", os.Args[1])
		printUsage()
		os.Exit(1)

	}

	
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  student-manager add --name=<name> --score=<0-100>")
	fmt.Println("  student-manager list  ")
	fmt.Println("  student-manager curve --points=<int>")

}

func runAdd(store student.Store, args[] string) {
	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)

	
	name := addCmd.String("name", "", "student name (required)")
	score := addCmd.Int("score", -1, "student score 0-100 (required)")


	//Prevent flag from print to stdout on errors

	addCmd.SetOutput(os.Stdout)

	err := addCmd.Parse(args)

	if err != nil {
		os.Exit(1)
	}

	if *name == "" || *score < 0 || *score > 100 {
		fmt.Println("Invalid input. Example:")
		fmt.Println("  student-manager add --name=Alex --score=88")
		os.Exit(1)
	}

	students, err := store.Load()

	if err != nil {
		fmt.Println("Failed to load students:", err)
		os.Exit(1)
	}

	students = append(students, student.New(*name, *score))

	if err := store.Save(students); err != nil {
		fmt.Println("Failed to save students:", err)
		os.Exit(1)
	}

	fmt.Println("Added:", student.New(*name, *score))
}

func runList(store student.Store) {
	students, err := store.Load()

	if err != nil {
		fmt.Println("Failed to load students:", err)
		os.Exit(1)
	}

	if len(students) == 0 {
		fmt.Println("No students yet. Add one with:")
		fmt.Println("  student-manager add --name=Alex --score=88")
		return
	}

	fmt.Println("Students:")

	for _, s := range students {
		fmt.Println(s)
	}
}

func runCurve(store student.Store, args []string) {
	curveCmd := flag.NewFlagSet("curve", flag.ContinueOnError)
	points := curveCmd.Int("points", 0, "points to add/subtract for all students")

	curveCmd.SetOutput(os.Stdout)

	if err := curveCmd.Parse(args); err != nil {
		os.Exit(1)
	}

	if *points == 0 {
		fmt.Println("No change (points=0) Example:")
		fmt.Println("   student-manager curve --points=5")
		os.Exit(1)
	}

	students, err := store.Load()

	if err != nil {
		fmt.Println("Failed to load students: ", err)
		os.Exit(1)
	}

	if len(students) == 0 {
		fmt.Println("No students to curve. Add students first.")
		os.Exit(1)
	
	}

	for i := range students {
		students[i].AddPoints(*points)
	}

	if err := store.Save(students); err != nil {
		fmt.Println("Failed to save students", err)
		os.Exit(1)
	}

	fmt.Println("Curved class by ", *points, "points.")
}