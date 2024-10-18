import gsap from "gsap";
import { ScrollTrigger } from "gsap/all";
import { Message } from "src/shared/ui/Message";
import { Form } from "src/widgets/Form";
import { FormGuySVG } from "./ui/FormGuySVG";
import { useEffect } from "react";

export const Consultation = () => {
  useEffect(() => {
    gsap.registerPlugin(ScrollTrigger);

    // танцующий возле формы мужичок
    gsap.to(".dancing-guy", {
      scrollTrigger: {
        trigger: ".dancing-guy",
        start: "top 600",
        end: "+=400",
        toggleActions: "restart pause restart pause",
      },
      duration: 1.5,
      rotationY: 360,
    });
    gsap.to(".dancing-guy", {
      scrollTrigger: {
        trigger: ".dancing-guy",
        start: "top 600",
        end: "+=400",
        toggleActions: "restart pause restart pause",
      },
      yoyo: true,
      repeat: 1,
      duration: 0.75,
      rotationX: 40,
    });
  }, []);
  return (
    <section id="consult" className="mt-16 lg:mt-24 mb-10 lg:mb-20">
      <div
        className={` flex flex-col md:flex-row md:justify-between gap-8 w-90% mx-auto`}
      >
        <h2 className={` font-bold clamp-h2 leading-[1]`}>
          Приглашаем на
          <span
            className={` font-caveat font-normal text-accent clamp-span leading-[0.8]`}
          >
            {" "}
            консультацию{" "}
          </span>
        </h2>
        <Message
          className=" rounded-bl-none sm:min-w-72 lg:text-base max-w-[65%] ml-auto  xl:block xl:mt-auto hidden "
          title="Расскажете детали, а мы поможем составить ТЗ"
        />
      </div>
      <div className="form-and-dancingguy mt-8 flex flex-col md:flex-row md:gap-4 md:justify-between md:w-90% md:mx-auto ">
        <Form />
        <div className="dancing-guy-wrapper md:mr-10">
          <Message
            className=" rounded-bl-none sm:min-w-72 sm:w-72 lg:text-base max-w-[65%] ml-auto hidden md:block xl:hidden md:mt-4"
            title="Расскажете детали, а мы поможем составить ТЗ"
          />
          <FormGuySVG className="dancing-guy mx-auto w-80 mt-10 md:mt-16 lg:w-96" />
        </div>
      </div>
    </section>
  );
};
