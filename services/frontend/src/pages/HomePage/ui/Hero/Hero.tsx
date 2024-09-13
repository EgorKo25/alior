import { TgButton } from "src/shared/ui/TgButton";
import { Guys } from "./ui/Guys";
import { CheckDemo } from "./ui/CheckDemo";
import { MainTitle } from "./ui/MainTitle";
import { ArrowSVG } from "./ui/ArrowSVG";

export const Hero = () => {
  return (
    <section className=" mt-10 md:mt-24 mx-5 sm:mx-10 xl:mx-28 ">
      <MainTitle />
      <div className=" relative flex justify-end items-end w-fit ml-auto gap-3 sm:gap-4 mt-4 sm:mt-12 lg:mt-24 2xl:mr-[5vw]">
        <span className=" absolute left-0 sm:-left-12 -translate-x-[70%] translate-y-2 text-lg sm:static sm:translate-x-0 sm:translate-y-0 sm:mb-4">
          Пишите нам
        </span>
        <ArrowSVG className=" w-36 sm:w-44 text-accent fill-accent" />
        <TgButton className=" bg-[#487CD4] hover:bg-[#487CD4]/80  transition-all" />
      </div>
      <div className="flex flex-col xl:flex-row xl:gap-24">
        <Guys />
        <CheckDemo />
      </div>
    </section>
  );
};
