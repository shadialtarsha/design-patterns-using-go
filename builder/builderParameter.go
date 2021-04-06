package builder

import "strings"

type email struct {
	from, to, subject, body string
}

type emailBuilder struct {
	email
}

func (b *emailBuilder) from(from string) *emailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *emailBuilder) to(to string) *emailBuilder {
	b.email.to = to
	return b
}

func (b *emailBuilder) subject(subject string) *emailBuilder {
	b.email.subject = subject
	return b
}

func (b *emailBuilder) body(body string) *emailBuilder {
	b.email.body = body
	return b
}

func sendMailImpl(email *email) {
	// Email sending implementation
}

type build func(*emailBuilder)

func sendEmail(action build) {
	builder := emailBuilder{}
	action(&builder)
	sendMailImpl(&builder.email)
}
