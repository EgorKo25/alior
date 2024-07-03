import logging
import asyncio
from aiogram import Bot, Dispatcher
from aiogram.types import BotCommand
from app.config.config import config
from app.core.handlers.public.handlers import register_handlers_public
from app.core.handlers.management.handlers import register_handlers_management


async def set_commands(bot: Bot):
    commands = [
        BotCommand(command="/start", description="Начать возню"),
        BotCommand(command="/get_callback", description="Дай callback по ID сучка биля"),
        BotCommand(command="/create_callback", description="Накинуть новый калбэк"),
        BotCommand(command="/delete_callback", description="Убрать его фром жизнь по ID"),
        BotCommand(command="/get_all_callbacks", description="Хачю все калбэки па номеру жи ес"),
    ]
    await bot.set_my_commands(commands)


async def main():
    logging.basicConfig(level=logging.INFO)

    bot = Bot(token=config['bot']['token'])
    dp = Dispatcher(bot=bot)

    register_handlers_public(dp)
    register_handlers_management(dp)

    await set_commands(bot)

    try:
        await dp.start_polling(bot)
    finally:
        await bot.close()


if __name__ == '__main__':
    asyncio.run(main())