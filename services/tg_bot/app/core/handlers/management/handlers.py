from aiogram import Dispatcher, types
from aiogram.filters import Command
from aiogram.types.inline_keyboard_button import InlineKeyboardButton
from aiogram.types.inline_keyboard_markup import InlineKeyboardMarkup
from app.api.grpc_client import GRPCClient

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
            args = command.args.split()
            offset = int(args[0])
            callbacks_quantity = grpc_client.get_callbacks_quantity().quantity
            responses = list(grpc_client.get_callbacks_paginated(limit, offset))
            if responses:
                callback = responses[0]
                text = f"Имя: {callback.Name}\n Дата: {callback.Date}\n Номер: {callback.Number}\n"

                keyboard_buttons = []
                if offset > 0:
                    keyboard_buttons.append(
                        InlineKeyboardButton(text="⬅️ Предыдущий", callback_data=f"prev_{offset - limit}"))
                if offset < callbacks_quantity - 1:
                    keyboard_buttons.append(
                        InlineKeyboardButton(text="➡️ Следующий", callback_data=f"next_{offset + limit}"))
                keyboard_buttons.append(InlineKeyboardButton(text="🗑️ Удалить", callback_data=f"delete"))

                keyboard = InlineKeyboardMarkup(inline_keyboard=[keyboard_buttons])

                await message.answer(text, reply_markup=keyboard)
            else:
                await message.answer("Нету таких.")
        except Exception as e:
            await message.answer(f"Error: {str(e)}")

    @dp.callback_query(lambda c: c.data and c.data.startswith('prev'))
    async def process_prev_callback(callback_query: types.CallbackQuery, limit=1):
        _, offset = callback_query.data.split('_')
        offset = int(offset)

        callbacks_quantity = grpc_client.get_callbacks_quantity().quantity
        responses = list(grpc_client.get_callbacks_paginated(limit, offset))

        if responses:
            callback = responses[0]
            text = f"Имя: {callback.Name}\n Дата: {callback.Date}\n Номер: {callback.Number}\n"

            keyboard_buttons = []
            if offset > 0:
                keyboard_buttons.append(
                    InlineKeyboardButton(text="⬅️ Предыдущий", callback_data=f"prev_{offset - limit}"))
            if offset < callbacks_quantity - 1:
                keyboard_buttons.append(
                    InlineKeyboardButton(text="➡️ Следующий", callback_data=f"next_{offset + limit}"))
            keyboard_buttons.append(InlineKeyboardButton(text="🗑️ Удалить", callback_data=f"delete"))

            keyboard = InlineKeyboardMarkup(inline_keyboard=[keyboard_buttons])

        await callback_query.message.edit_text(text, reply_markup=keyboard)

    @dp.callback_query(lambda c: c.data and c.data.startswith('next'))
    async def process_next_callback(callback_query: types.CallbackQuery, limit=1):
        _, offset = callback_query.data.split('_')
        offset = int(offset)

        callbacks_quantity = grpc_client.get_callbacks_quantity().quantity
        responses = list(grpc_client.get_callbacks_paginated(limit, offset))

        if responses:
            callback = responses[0]
            text = f"Имя: {callback.Name}\n Дата: {callback.Date}\n Номер: {callback.Number}\n"

            keyboard_buttons = []
            if offset > 0:
                keyboard_buttons.append(
                    InlineKeyboardButton(text="⬅️ Предыдущий", callback_data=f"prev_{offset - limit}"))
            if offset < callbacks_quantity - 1:
                keyboard_buttons.append(
                    InlineKeyboardButton(text="➡️ Следующий", callback_data=f"next_{offset + limit}"))
            keyboard_buttons.append(InlineKeyboardButton(text="🗑️ Удалить", callback_data=f"delete"))

            keyboard = InlineKeyboardMarkup(inline_keyboard=[keyboard_buttons])

            await callback_query.message.edit_text(text, reply_markup=keyboard)

    @dp.callback_query(lambda c: c.data and c.data.startswith('delete'))
    async def process_delete_callback(callback_query: types.CallbackQuery):
        _, callback_id = callback_query.data.split('_')
        callback_id = int(callback_id)
        grpc_client.delete_callback(callback_id)
        await callback_query.answer("Удалено")
        await callback_query.message.delete()
