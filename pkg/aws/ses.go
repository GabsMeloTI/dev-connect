package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"treads/internal/model"
)

type SESClient struct {
	Client *ses.SES
}

func NewSESClient(region string) *SESClient {
	if region == "" {
		panic("A configuração da região AWS está ausente")
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &SESClient{
		Client: ses.New(sess),
	}
}

func (s *SESClient) SendEmail(data model.EmailResponse) error {
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(data.To)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(data.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(data.Subject),
			},
		},
		Source: aws.String("devconnectsocialmedia@gmail.com"),
	}

	output, err := s.Client.SendEmail(input)
	if err != nil {
		return fmt.Errorf("erro ao enviar e-mail via SES: %v", err)
	}

	fmt.Printf("E-mail enviado com sucesso! Message ID: %s\n", *output.MessageId)
	return nil
}
