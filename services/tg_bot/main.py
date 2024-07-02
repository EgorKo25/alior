import logging
import asyncio
from aiogram import Bot, Dispatcher
from aiogram.types import BotCommand
from app.config.config import config
from app.core.handlers.public.handlers import register_handlers_public


async def set_commands(bot: Bot):
    commands = [
        BotCommand(command="/start", description="Начать возню"),
    ]
    await bot.set_my_commands(commands)


async def main():
    logging.basicConfig(level=logging.INFO)

    bot = Bot(token=config['token'])
    dp = Dispatcher(bot=bot)

    register_handlers_public(dp)

    await set_commands(bot)

    try:
        await dp.start_polling(bot)
    finally:
        await bot.close()


if __name__ == '__main__':
    asyncio.run(main())