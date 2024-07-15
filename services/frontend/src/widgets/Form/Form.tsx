import { useEffect, useState } from "react";
import { CustomRadioInput } from "./ui/CustomRadioInput";
import { InputLabeled } from "./ui/InputLabeled";
import { MainButton } from "src/shared/ui/MainButton";
import axios from "axios";

export const Form = () => {
  const [formMessage, setFormMessage] = useState("");
  const [isFormValid, setIsFormValid] = useState(false);
  const [formData, setFormData] = useState({
    idea: "",
    name: "",
    phone: "",
    callbackType: "call",
  });

  useEffect(() => {
    setIsFormValid(
      regexName.test(formData.name) && regexTel.test(formData.phone)
    );
  }, [formData]);

  const regexName = /^[a-zA-Zа-яА-ЯЁё\- ]+$/;
  const regexTel = /^[0-9]{11}$/;

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (isFormValid == false) {
      setFormMessage("Заполните форму правильно");
      return;
    }
    axios
      .post("/callback", formData)
      .then((response) => {
        setFormMessage(response.statusText);
        console.log(response.data);
      })
      .catch((error) => {
        setFormMessage(error.message);
        console.error(error);
      });
  };

  const handleChangeValue = (
    e: React.ChangeEvent<HTMLInputElement>,
    id: string
  ) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
      callbackType: id,
    });
  };

  return (
    <form
      action=""
      className="bg-accent w-full px-5 py-8 lg:px-10 2xl:px-16 2xl:py-12 rounded-[24px] md:w-[55%] xl:w-[70%]"
      onSubmit={handleSubmit}
    >
      <div className="inputs flex flex-col gap-8">
        <InputLabeled
          label={"Опишите идею коротко"}
          placeholder={"Комментарий о проекте"}
          type={"text"}
          name={"idea"}
          onChange={(e) => handleChangeValue(e, formData.callbackType)}
          isIputValueValid={true}
        />
        <div className="name-and-phone flex w-full gap-6">
          <InputLabeled
            label={"Имя"}
            placeholder={"Иван"}
            type={"text"}
            name={"name"}
            onChange={(e) => handleChangeValue(e, formData.callbackType)}
            isIputValueValid={
              regexName.test(formData.name) || formData.name === ""
            }
          />
          <InputLabeled
            label={"Телефон"}
            placeholder={"7 800 555 3535"}
            type={"tel"}
            name={"phone"}
            onChange={(e) => handleChangeValue(e, formData.callbackType)}
            isIputValueValid={
              regexTel.test(formData.phone) || formData.phone === ""
            }
          />
        </div>
        <div className="callback-types-and-button flex flex-col justify-between xl:flex-row xl:items-center">
          <div className="types-of-callback">
            <fieldset>
              <legend className="text-white text-lg sm:text-xl">
                Способ связи
              </legend>
              <div className="radio-options mt-4 flex flex-wrap gap-2 xl:justify-around">
                <CustomRadioInput
                  name={"callbackType"}
                  label={"Звонок"}
                  id={"call"}
                  isChecked={formData.callbackType === "call"}
                  onChange={(e) => {
                    handleChangeValue(e, "call");
                  }}
                />
                <CustomRadioInput
                  name={"callbackType"}
                  label={"WhatsApp"}
                  id={"whatsapp"}
                  isChecked={formData.callbackType === "whatsapp"}
                  onChange={(e) => {
                    handleChangeValue(e, "whatsapp");
                  }}
                />
                <CustomRadioInput
                  name={"callbackType"}
                  label={"Telegram"}
                  id={"telegram"}
                  isChecked={formData.callbackType === "telegram"}
                  onChange={(e) => {
                    handleChangeValue(e, "telegram");
                  }}
                />
              </div>
            </fieldset>
          </div>
          <div className="button-and-text flex flex-col">
            <h4 className=" text-center text-white text-sm sm:text-base font-light lg:text-right mt-6">
              {formMessage.length > 0
                ? formMessage
                : "Обещием, что ваши данные в безопасности"}
            </h4>
            <MainButton
              title={"На консультацию"}
              className={" border-none w-full mt-4 lg:w-auto lg:ml-auto"}
              colorSchema={" btn-white-black"}
            />
          </div>
        </div>
      </div>
    </form>
  );
};
