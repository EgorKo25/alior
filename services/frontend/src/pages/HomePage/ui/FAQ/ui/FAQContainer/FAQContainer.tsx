import { Message } from "src/shared/ui/Message";
import { FirstArrow } from "./ui/FirstArrow";
import { SecondArrow } from "./ui/SecondArrow";
import { MeditationSVG } from "./ui/MeditationSVG";

export const FAQContainer = () => {
  return (
    <div
      className={
        " rounded-40 border-solid border-6 border-orange-100 w-full px-5 pt-28 text-blue-900 md:w-90% md:px-8 lg:px-16 2xl:px-52"
      }
    >
      <div className={"question-block ml-auto flex flex-col items-end mb-40"}>
        <h2
          className={
            " most-popular font-bold clamp-h2 leading-[1] text-right opacity-0"
          }
        >
          Самый популярный
        </h2>
        <span className=" question font-caveat font-normal clamp-span relative opacity-0">
          Вопрос
          <FirstArrow
            className={
              " absolute -left-[11.5rem] top-10 md:top-12 lg:top-16 xl:top-24 xl:scale-[1.4] xl:-left-56"
            }
          />
          <Message
            className={
              " border-blue-900 text-blue-900 rounded-br-none w-60 lg:w-72 2xl:w-[22rem] lg:text-xl 2xl:max-w-96 2xl:text-2xl absolute -left-24 top-44 md:top-[11.5rem] lg:top-48 xl:-left-32 xl:top-64"
            }
            title="Как будет выглядеть разработка Тех задания?"
          />
        </span>
      </div>
      <div className={" flex flex-col items-start self-start relative"}>
        <span className=" answering font-caveat font-normal clamp-span opacity-0 ">
          Отвечаем
        </span>
        <SecondArrow
          className={
            " answer-arrow absolute left-24 top-14 md:left-32 md:top-16 lg:left-40 lg:top-[4.5rem] xl:left-48 xl:scale-[1.4] xl:top-[7.5rem] opacity-0"
          }
        />
        <div className={" answer-text mt-40 xl:mt-60 mb-12 md:mb-20 opacity-0"}>
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
            "meditation-girl w-80 h-64 mb-12 md:absolute md:-right-5 bottom-28  lg:-right-10 lg:bottom-32 xl:right-10 2xl:-right-5 self-center lg:w-431 lg:h-96 2xl:w-540 2xl:h-440 opacity-0"
          }
        >
          <MeditationSVG
            className={
              "object-cover object-center w-full h-full mx-auto xl:scale-125"
            }
          />
        </div>
      </div>
    </div>
  );
};
