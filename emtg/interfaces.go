package emtg

type TelegramBot interface {
    SendMessage(interface{}, string) (Message, error)
    GetEventsChannel() chan interface{}
}
