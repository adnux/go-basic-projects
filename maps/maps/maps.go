package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func MapsWebsite() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println("Websites=>", websites)
	fmt.Println("websites[Amazon Web Services]=>", websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println("Websites=>", websites)

	delete(websites, "Google")
	fmt.Println("Websites=>", websites)

}
