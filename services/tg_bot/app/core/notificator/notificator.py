import asyncio
from aiogram import Bot
from app.config.config import config

class NotificationSender:
    def __init__(self):
        self.bot = Bot(token=config['bot']['token'])
        self.chat_id = config['bot']['chat_id']

    async def send_telegram_message(self, message: str):
        await self.bot.send_message(self.chat_id, message)
