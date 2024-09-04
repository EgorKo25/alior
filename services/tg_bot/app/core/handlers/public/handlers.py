from aiogram import types, Dispatcher
from aiogram.filters import CommandStart


def register_handlers_public(dp: Dispatcher):
    @dp.message(CommandStart())
    async def start_command(message: types.Message):
        await message.answer("Здарова чурка я ALior bot и буду делать тебя рот")