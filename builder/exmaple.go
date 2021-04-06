package builder

import "fmt"

func Example() {
	builder := newHTMLBuilder("ul")
	builder.addChild("li", "cheese").
		addChild("li", "potato").
		addChild("li", "bannana")
	fmt.Println(builder.string())

	pb := newPersonBuilder()
	pb.lives().
		at("str1 108").
		in("Berlin").
		withPostcode("10115").
		works().
		at("Jodel").
		as("Backend Developer").
		earning(10)
	person := pb.build()
	fmt.Println(*person)

	sendEmail(func(b *emailBuilder) {
		b.from("shadi@bar.com").
			to("bar@shadi.com").
			subject("testing").
			body("hello")
	})

	b := employeeBuilder{}
	e := b.called("Shadi").worksAs("dev").build()
	fmt.Print(*e)
}
