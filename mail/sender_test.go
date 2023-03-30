package mail

import (
	"testing"

	"github.com/fredele20/Golang-backend-master/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	
	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
		<h1>Hello World</h1>
		<p>This is a test message from <a href="https://github.com/fredele20">Fredel</a></p>
	`

	to := []string{"frederickvictor93@gmail.com"}
	attachFiless := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiless)

	require.NoError(t, err)
}
