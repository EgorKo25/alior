import { Message } from "src/shared/ui/Message";
import { Thesis } from "./ui/Thesis";
import { MobileTree } from "./ui/MobileTree";
import { MiddleTree } from "./ui/MiddleTree";
import { LargeTree } from "./ui/LargeTree";

export const Features = () => {
  return (
    <section
      id="features"
      className={` flex flex-col my-0 mx-auto mt-48 w-90% relative overflow-hidden`}
    >
      <div
        className={` flex flex-col md:flex-row md:justify-between gap-8 mb-20`}
      >
        <h2
          className={` flex flex-col  font-bold clamp-h2 leading-[1] lg:flex-row lg:items-center lg:gap-4 whitespace-nowrap`}
        >
          Почему нас
          <span
            className={` font-caveat font-normal text-accent clamp-span lg:-mt-4`}
          >
            выбирают
          </span>
        </h2>
        <Message
          className=" rounded-bl-none max-w-96 lg:text-base"
          title="Лучше всего за нас скажут наши работы. Но если нужны тезисы, они ниже слайдом"
        />
      </div>
      <div
        className={` thesis_container_mobile self-center max-w-335 md:max-w-688 lg:max-w-944 md:thesis_container_middle lg:thesis_container_large relative`}
      >
        <MobileTree
          className={` absolute -z-10 -left-7% -top-52 md:hidden lg:hidden text-accent/50`}
        />
        <MiddleTree
          className={` absolute -top-32 left-17% hidden lg:hidden md:block -z-10 text-accent/50`}
        />
        <LargeTree
          className={` absolute -top-40 w-668 h-598 xl:w-874 xl:h-583 xl:left-4% xl:-top-20 hidden left-10% lg:block -z-10 text-accent/50 `}
        />
        <Thesis index={0} className={` item1 lg:min-w-72 md:-mt-7 xl:mt-0`} />
        <Thesis index={1} className={` item2 md:-mt-7 lg:mt-0`} />
        <Thesis index={2} className={` item3 lg:mt-8 md:mt-8 xl:mt-8`} />
        <Thesis
          index={3}
          className={` item4 md:mt-8 md:ml-3 lg:mt-4 xl:ml-5 xl:mt-10`}
        />
        <Thesis
          index={4}
          className={` item5 lg:mt-6 lg:min-w-72 md:mt-11 xl:mt-0`}
        />
        <Thesis index={5} className={` item6 md:mt-20 lg:mt-0`} />
        <Thesis index={6} className={` item7 mt-14 md:mt-0`} />
      </div>
      <div
        className={` opacity-0 md:opacity-10 md:left-0 size-44 rounded-full bg-accent absolute bottom-14 left-7% -z-10`}
      ></div>
      <svg
        className={` opacity-0 md:right-0 md:opacity-100 absolute right-7% top-8 -z-10`}
        width="282"
        height="282"
        viewBox="0 0 282 282"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          opacity="0.1"
          d="M140.024 1.33077C140.259 0.290445 141.741 0.290439 141.976 1.33077L167.335 113.909C167.42 114.286 167.714 114.58 168.091 114.665L280.669 140.024C281.71 140.259 281.71 141.741 280.669 141.976L168.091 167.335C167.714 167.42 167.42 167.714 167.335 168.091L141.976 280.669C141.741 281.71 140.259 281.71 140.024 280.669L114.665 168.091C114.58 167.714 114.286 167.42 113.909 167.335L1.33077 141.976C0.290445 141.741 0.290439 140.259 1.33077 140.024L113.909 114.665C114.286 114.58 114.58 114.286 114.665 113.909L140.024 1.33077Z"
          fill="#FF5678"
        />
      </svg>
    </section>
  );
};
