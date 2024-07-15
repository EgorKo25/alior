import { Message } from "src/shared/ui/Message";
import { Form } from "src/widgets/Form";

export const Consultation = () => {
  return (
    <section className=" mt-48 mb-20">
      <div
        className={` flex flex-col md:flex-row md:justify-between gap-8 w-90% mx-auto`}
      >
        <h2 className={` font-bold clamp-h2 leading-[1]`}>
          Приглашаем вас на бесплатную
          <span
            className={` font-caveat font-normal text-accent clamp-span leading-[0.8]`}
          >
            {" "}
            консультацию{" "}
          </span>
          для обсуждения деталей
        </h2>
        <Message
          className=" rounded-bl-none sm:min-w-72 lg:text-base max-w-[65%] ml-auto md:hidden xl:block xl:mt-auto"
          title="Расскажете детали, а мы поможем составить ТЗ"
        />
      </div>
      <div className="form-and-dancingguy mt-8 flex flex-col md:flex-row md:gap-4 md:justify-between md:w-90% md:mx-auto ">
        <Form />
        <div className="guy-wrapper md:mr-10">
          <Message
            className=" rounded-bl-none sm:min-w-72 sm:w-72 lg:text-base max-w-[65%] ml-auto hidden md:block xl:hidden md:mt-4"
            title="Расскажете детали, а мы поможем составить ТЗ"
          />
          <img
            src="/images/formguy.svg"
            alt="flexing-guy"
            className=" mx-auto w-80 md:mt-16 lg:w-96"
          />
        </div>
      </div>
    </section>
  );
};
