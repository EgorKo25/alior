import { HashLink } from "react-router-hash-link";
import { useEffect } from "react";
import gsap from "gsap";
import { ScrollTrigger } from "gsap/all";
import { MainButton } from "src/shared/ui/MainButton";
import { FAQContainer } from "./ui/FAQContainer";
import { SendForConsult } from "./ui/SendForConsult";
import { BookGuySVG } from "./ui/BookGuySVG";

export const FAQ = () => {
  useEffect(() => {
    gsap.registerPlugin(ScrollTrigger);

    // Книжный мужичок
    gsap.to(".book-guy", {
      scrollTrigger: {
        trigger: ".book-guy",
        scrub: 1,
        start: "top 500",
        end: "+=400",
      },
      filter: "opacity(1)",
    });

    // Самый популярный
    gsap.to(".most-popular", {
      scrollTrigger: {
        trigger: ".most-popular",
        scrub: 1,
        start: "top 700",
        end: "+=400",
      },
      filter: "opacity(1)",
    });

    // Вопрос стрелка и текст вопроса
    gsap.to(".question", {
      scrollTrigger: {
        trigger: ".question",
        scrub: 1,
        start: "top 600",
        end: "+=400",
      },
      filter: "opacity(1)",
    });

    // Отвечаем
    gsap.to(".answering", {
      scrollTrigger: {
        trigger: ".answering",
        scrub: 1,
        start: "top 600",
        end: "+=400",
      },
      filter: "opacity(1)",
    });

    // стрелка после отвечаем
    gsap.to(".answer-arrow", {
      scrollTrigger: {
        trigger: ".answer-arrow",
        scrub: 1,
        start: "top 600",
        end: "+=400",
      },
      filter: "opacity(1)",
    });

    // сам ответ
    gsap.to(".answer-text", {
      scrollTrigger: {
        trigger: ".answer-text",
        scrub: 1,
        start: "top 600",
        end: "+=400",
      },
      filter: "opacity(1)",
    });

    //медитирует сидит
    gsap.to(".meditation-girl", {
      scrollTrigger: {
        trigger: ".meditation-girl",
        scrub: 1,
        start: "top 550",
        end: "+=400",
      },
      filter: "opacity(1)",
    });
  }, []);

  return (
    <section
      className={` my-0 mx-auto mt-20 pt-48 flex flex-col items-center justify-center w-full bg-blue-100 pb-10 relative md:pt-14 
        before:absolute before:content-[''] before:h-40 before:w-full before:top-1 before:-translate-y-[100%] before:bg-gradient-to-b before:from-transparent before:via-blue-100 before:via-80% before:to-blue-100
        after:absolute after:content-[''] after:h-40 after:w-full after:bottom-1 after:translate-y-[100%] after:bg-gradient-to-b after:to-transparent after:from-5% after:via-20% after:via-blue-100/90 after:from-blue-100`}
    >
      <SendForConsult />
      <div
        className={
          "book-guy w-56 h-56  relative md:self-start md:ml-24 lg:w-72 lg:h-72 2xl:ml-40 2xl:w-540 2xl:h-500"
        }
        style={{ filter: "opacity(0)" }}
      >
        <BookGuySVG
          className={
            "object-cover object-center  absolute top-20 w-full h-full 2xl:top-48 scale-150 lg:scale-125 lg:top-28 "
          }
        />
      </div>
      <FAQContainer />
      <HashLink
        smooth
        to={"#consult"}
        className=" absolute bottom-0 -translate-y-5 2xl:-translate-y-[0.75rem]"
      >
        <MainButton
          className=""
          colorSchema={" btn-accent-white"}
          title="На консультацию"
        />
      </HashLink>
    </section>
  );
};
