from aiogram.types.inline_keyboard_button import InlineKeyboardButton
from aiogram.types.inline_keyboard_markup import InlineKeyboardMarkup

def create_response_keyboard(response, offset, limit):
    if response:
        callback = response.callbacks[0]
        text = f"Имя: {callback.Name}\nДата: {callback.Date}\nНомер: {callback.Number}\n"

        keyboard_buttons = []
        if offset > 0:
            keyboard_buttons.append(
                InlineKeyboardButton(text="⬅️ Предыдущий", callback_data=f"prev_{offset - limit}")
            )
        if offset < response.total_items - 1:
            keyboard_buttons.append(
                InlineKeyboardButton(text="➡️ Следующий", callback_data=f"next_{offset + limit}")
            )
        delete_button = [InlineKeyboardButton(text="🗑️ Удалить", callback_data=f"delete_{callback.Number}_{limit}_{offset}")]

        keyboard = InlineKeyboardMarkup(inline_keyboard=[keyboard_buttons, delete_button])
        return text, keyboard
    return "хз ниче в ответе нет", None