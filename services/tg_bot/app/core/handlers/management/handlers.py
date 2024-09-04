from aiogram import Dispatcher, types
from aiogram.filters import Command
from app.core.common.interactive_keyboard import create_response_keyboard
from app.core.rabbitmq.producer import Producer
from app.config.config import config
import json

def register_handlers_management(dp: Dispatcher):

    @dp.message(Command("delete_callback"))
    async def delete_callback_command(message: types.Message, command: Command):
        callback_id = message.get_args()
        producer = Producer(config['rabbitmq']['url'], "request")
        await producer.produce({"command": "delete_callback", "id": int(callback_id)})
        await message.reply(f"Запрос на удаление callback с ID {callback_id} отправлен.")

    @dp.message(Command("get_callbacks"))
    async def get_callback_command(message: types.Message, command: Command):
        producer = Producer(config['rabbitmq']['url'], "request")
        await producer.produce({"command": "get_callback", "limit": 1, "offset": 0})
        await message.reply("Запрос на получение callback отправлен, ожидайте ответ.")

    @dp.callback_query_handler(lambda c: c.data and (c.data.startswith('prev_') or c.data.startswith('next_')))
    async def navigate_callback(callback_query: types.CallbackQuery):
        direction, current_id = callback_query.data.split('_')
        current_id = int(current_id)
        offset = current_id - 1 if direction == 'prev' else current_id + 1

        producer = Producer(config['rabbitmq']['url'], "request")
        await producer.produce({"command": "get_callback", "limit": 1, "offset": offset})
        await bot.answer_callback_query(callback_query.id)
        await bot.send_message(callback_query.from_user.id, "Запрос отправлен, ожидайте ответ.")

    @dp.callback_query_handler(lambda c: c.data and c.data.startswith('delete_'))
    async def delete_callback(callback_query: types.CallbackQuery):
        callback_id = int(callback_query.data.split('_')[1])
        producer = Producer(config['rabbitmq']['url'], "request")
        await producer.produce({"command": "delete_callback", "id": callback_id})
        await bot.answer_callback_query(callback_query.id)
        await bot.send_message(callback_query.from_user.id, f"Callback {callback_id} удален.")
