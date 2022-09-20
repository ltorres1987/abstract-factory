package main

import "fmt"

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type EmailNotification struct {
}

func (e EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via EMAIL")
}

func (e EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (e EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (e EmailNotificationSender) GetSenderChannel() string {
	return "aws"
}

type SMSNotification struct {
}

func (s SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

func (s SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (s SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (s SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

func getNotificationfactory(nType string) (INotificationFactory, error) {

	if nType == "SMS" {
		return &SMSNotification{}, nil
	}

	if nType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification type")
}

func SendNotification(n INotificationFactory) {
	n.SendNotification()
}

func GetMethod(n INotificationFactory) {
	fmt.Println(n.GetSender().GetSenderMethod())
}

func GetChannel(n INotificationFactory) {
	fmt.Println(n.GetSender().GetSenderChannel())
}

func main() {

	s, _ := getNotificationfactory("SMS")
	e, _ := getNotificationfactory("Email")

	SendNotification(s)
	GetMethod(s)
	GetChannel(s)

	SendNotification(e)
	GetMethod(e)
	GetChannel(e)
}
