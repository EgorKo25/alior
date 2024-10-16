import { useEffect } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/all";
import { TgButton } from "src/shared/ui/TgButton";
import { Guys } from "./ui/Guys";
import { CheckDemo } from "./ui/CheckDemo";
import { MainTitle } from "./ui/MainTitle";
import { ArrowSVG } from "./ui/ArrowSVG";

export const Hero = () => {
  useEffect(() => {
    gsap.registerPlugin(ScrollTrigger);

    gsap.to(".demo-suggestion", {
      scrollTrigger: {
        trigger: ".demo-suggestion",
        scrub: 1,
        start: "top 600",
        end: "+=300",
      },
      translateX: 0,
    });

    gsap.to(".selfie-guy", {
      scrollTrigger: {
        trigger: ".selfie-guy",
        scrub: 1,
        start: "-150 600",
        end: "+=300",
      },
      translateX: 0,
    });

    gsap.to(".guy-in-box", {
      scrollTrigger: {
        trigger: ".guy-in-box",
        scrub: 1,
        start: "-150 600",
        end: "+=300",
      },
      translateX: 0,
    });
  }, []);

  return (
    <section className=" mt-10 md:mt-24 mx-5 sm:mx-10 xl:mx-28 overflow-hidden">
      <MainTitle />
      <div className=" relative flex justify-end items-end w-fit ml-auto gap-3 sm:gap-4 mt-4 sm:mt-12 lg:mt-24 2xl:mr-[5vw]">
        <span className=" absolute left-0 sm:-left-12 -translate-x-[70%] translate-y-2 text-lg md:text-2xl sm:static sm:translate-x-0 sm:translate-y-0 sm:mb-4">
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
