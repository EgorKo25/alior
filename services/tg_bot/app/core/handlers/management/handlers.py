from pprint import pprint

from aiogram import Dispatcher, types
from aiogram.filters import Command
from app.api.grpc_client import GRPCClient

from app.core.common.interactive_keyboard import create_response_keyboard

grpc_client = GRPCClient()


def register_handlers_management(dp: Dispatcher):

    @dp.message(Command("delete_callback"))
    async def delete_callback_command(message: types.Message, command: Command):
        try:
            callback_id = int(command.args)
            grpc_client.delete_callback(callback_id)
            await message.answer("Убрал его из твоей жизни")
        except Exception as e:
            await message.answer(f"Error: {str(e)}")

    @dp.message(Command("get_callbacks_paginated"))
    async def get_callback_command(message: types.Message, command: Command, limit=1):
        try:
            args = command.args.split() if command.args else []

            if len(args) == 1:
                offset = int(args[0])
            else:
                raise AttributeError("Слышь праститутка доку почитай и делай нормальные аргументы")

            response = grpc_client.get_callbacks_paginated(limit, offset)
            text, keyboard = create_response_keyboard(response, offset, limit)

            await message.answer(text, reply_markup=keyboard)

        except Exception as e:
            await message.answer(f"Error: {str(e)}")

    @dp.callback_query(lambda c: c.data and c.data.startswith('prev'))
    async def process_prev_callback(callback_query: types.CallbackQuery, limit=1):
        _, offset = callback_query.data.split('_')
        offset = int(offset)

        response = grpc_client.get_callbacks_paginated(limit, offset)
        text, keyboard = create_response_keyboard(response, offset, limit)

        await callback_query.message.edit_text(text, reply_markup=keyboard)

    @dp.callback_query(lambda c: c.data and c.data.startswith('next'))
    async def process_next_callback(callback_query: types.CallbackQuery, limit=1):
        _, offset = callback_query.data.split('_')
        offset = int(offset)

        response = grpc_client.get_callbacks_paginated(limit, offset)
        text, keyboard = create_response_keyboard(response, offset, limit)

        await callback_query.message.edit_text(text, reply_markup=keyboard)

    @dp.callback_query(lambda c: c.data and c.data.startswith('delete'))
    async def process_delete_callback(callback_query: types.CallbackQuery):
        _, callback_number, limit, offset = callback_query.data.split('_')
        callback_number = str(callback_number)

        grpc_client.delete_callback_by_number(callback_number)
        await callback_query.message.answer("Удалил его жи ес")

        # await callback_query.message.delete()
