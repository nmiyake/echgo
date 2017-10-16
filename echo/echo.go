package echo

func Echo(in string) string {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
