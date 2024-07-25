from aiogram import Dispatcher, types
from aiogram.filters import Command
from app.core.common.interactive_keyboard import create_response_keyboard


def register_handlers_management(dp: Dispatcher):

    @dp.message(Command("delete_callback"))
    async def delete_callback_command(message: types.Message, command: Command):
        pass

    @dp.message(Command("get_callbacks_paginated"))
    async def get_callback_command(message: types.Message, command: Command):
        pass


