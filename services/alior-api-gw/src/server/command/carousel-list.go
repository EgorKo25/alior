package command

import "github.com/gin-gonic/gin"

type CarouselList struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func (c *CarouselList) Name() string {
	return "carousel/list"
}

func (c *CarouselList) Parse(ctx *gin.Context) error { return nil }

func (c *CarouselList) Apply() (any, error) {
	/*
		body, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}

		id := sha256.Sum256(body)
		request := &amqp.Publishing{
			CorrelationId: fmt.Sprintf("%x", id),
			Timestamp:     time.Now(),
			Type:          "create",
			ReplyTo:       "ANS",
			DeliveryMode:  1,
			ContentType:   "callback",
			Body:          body,
		}
	*/
	// TODO:отправка в Rabbit
	//
	// TODO: и ожидание ответа
	//
	// responce :=
	return nil, nil
}
