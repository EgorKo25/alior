from aiogram import Dispatcher, types
from app.api.grpc_client import GRPCClient
from app.api.proto import tn_pb2

grpc_client = GRPCClient()

def register_handlers_management(dp: Dispatcher):
    @dp.message_handler(commands=["get_callback"])
    async def get_callback_command(message: types.Message):
        try:
            callback_id = int(message.get_args())
            response = grpc_client.get_callback(callback_id)
            await message.answer(f"Callback: Name: {response.Name}, Date: {response.Date}, Number: {response.Number}")
        except Exception as e:
            await message.answer(f"Ашыбка: {str(e)}")

    @dp.message_handler(commands=["create_callback"])
    async def create_callback_command(message: types.Message):
        try:
            data = message.get_args().split(',')
            if len(data) != 3:
                raise ValueError("Пример для тупых: /create_callback <Имя блять>,<дата блять>,<номер нахуй>")

            callback = tn_pb2.CallBack(Name=data[0], Date=data[1], Number=data[2])
            grpc_client.create_callback(callback)
            await message.answer("Создано на кайфе")
        except Exception as e:
            await message.answer(f"Error: {str(e)}")

    @dp.message_handler(commands=["delete_callback"])
    async def delete_callback_command(message: types.Message):
        try:
            callback_id = int(message.get_args())
            grpc_client.delete_callback(callback_id)
            await message.answer("Убрал его из твоей жизни")
        except Exception as e:
            await message.answer(f"Error: {str(e)}")

    @dp.message_handler(commands=["get_all_callbacks"])
    async def get_all_callbacks_command(message: types.Message):
        try:
            number = message.get_args()
            responses = grpc_client.get_all_callbacks(number)
            callbacks = [f"Имя: {callback.Name}, Дата: {callback.Date}, Номер: {callback.Number}" for callback in responses]
            if callbacks:
                await message.answer("\n\n".join(callbacks))
            else:
                await message.answer("Нету таких.")
        except Exception as e:
            await message.answer(f"Error: {str(e)}")