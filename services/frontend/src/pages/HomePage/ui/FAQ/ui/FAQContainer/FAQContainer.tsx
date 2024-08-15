import { Message } from "src/shared/ui/Message";
import { FirstArrow } from "../FirstArrow";
import { SecondArrow } from "../SecondArrow";

export const FAQContainer = () => {
  return (
    <div
      className={
        " flex flex-col rounded-40 border-solid border-6 border-orange-100 w-full px-5 pt-28 text-blue-900 md:w-90% md:px-8 lg:px-16 2xl:px-52"
      }
    >
      <div className={" flex flex-col items-end self-end w-min relative"}>
        <h2
          className={" flex flex-col items-end font-bold clamp-h2 leading-[1]"}
        >
          Самый популярный
          <span className=" font-caveat font-normal clamp-span  ">Вопрос</span>
        </h2>
        <FirstArrow />
        <Message
          className={
            " border-blue-900 text-blue-900 rounded-br-none max-w-60 mt-28 lg:max-w-72 lg:text-xl 2xl:max-w-96 2xl:text-2xl"
          }
          title="Как будет выглядеть разработка Тех задания?"
        />
      </div>
      <div className={" flex flex-col items-start self-start relative"}>
        <span className=" font-caveat font-normal clamp-span  ">Отвечаем</span>
        <SecondArrow />
        <div className={" flex flex-col md:flex-row mt-44 mb-24"}>
          <div className={" "}>
            <Message
              className={
                " border-blue-900 text-blue-900 text-nowrap rounded-bl-none max-w-335 font-medium text-xl lg:max-w-96 lg:text-2xl 2xl:max-w-lg 2xl:text-3xl"
              }
              title="Уточним, обсудим ваш проект"
            />
            <Message
              className={
                " border-blue-900 text-blue-900 rounded-bl-none max-w-64 mt-2 lg:max-w-80 lg:text-xl 2xl:text-2xl 2xl:max-w-420"
              }
              title="Поговорим о сроках и оплате. Все это сделаем в открытой и понятной форме, чтобы вы чувствовали себя комфортно на протяжении всего сотрудничества с нами"
            />
            <Message
              className={
                " border-blue-900 text-blue-900 rounded-bl-none max-w-80 lg:max-w-96 mt-7 lg:text-xl 2xl:text-2xl 2xl:max-w-md"
              }
              title="Далее наш менеджер заботливо соберет все ваши идеи, опишет все технологии, которые мы будем использовать и объяснит их актуальность и целесообразность."
            />
          </div>
          <div
            className={
              "w-80 h-64 mt-14 md:mt-0 md:mb-28 self-center lg:w-431 lg:h-96 2xl:w-540 2xl:h-440"
            }
          >
            <img
              className={"object-cover object-center w-full h-full"}
              src="./images/meditation.svg"
            ></img>
          </div>
        </div>
      </div>
    </div>
  );
};
