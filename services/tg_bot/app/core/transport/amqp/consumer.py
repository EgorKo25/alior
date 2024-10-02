import aio_pika
import asyncio
from app.config.config import config
import json

class Consumer:
    def __init__(self, bot, amqp_url: str, queue_name: str):
        self.bot = bot
        self.amqp_url = amqp_url
        self.queue_name = queue_name

    async def on_message(self, message: aio_pika.IncomingMessage):
        async with message.process():
            print(f" [x] Received message {message.body.decode()}")
            await self.bot.send_message(config['bot']['chat_id'], "Новый callback")

    async def consume(self):
        connection = await aio_pika.connect_robust(self.amqp_url)
        async with connection:
            channel = await connection.channel()
            queue = await channel.declare_queue(self.queue_name, durable=True)
            await queue.consume(self.on_message)
            print(" [*] Waiting for messages. To exit press CTRL+C")
            await asyncio.Future()  # Run forever
