import aio_pika
import json

class Producer:
    def __init__(self, amqp_url: str, queue_name: str):
        self.amqp_url = amqp_url
        self.queue_name = queue_name

    async def produce(self, message: dict):
        connection = await aio_pika.connect_robust(self.amqp_url)
        async with connection:
            channel = await connection.channel()
            await channel.default_exchange.publish(
                aio_pika.Message(body=json.dumps(message).encode()),
                routing_key=self.queue_name,
            )
