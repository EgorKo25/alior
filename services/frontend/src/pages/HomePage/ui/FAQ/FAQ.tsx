import { MainButton } from "src/shared/ui/MainButton";
import { FAQContainer } from "./ui/FAQContainer";
import { SendForConsult } from "./ui/SendForConsult";
import { BookGuySVG } from "./ui/BookGuySVG";

export const FAQ = () => {
  return (
    <section
      className={` my-0 mx-auto mt-48 pt-48 flex flex-col items-center justify-center w-full bg-blue-100 pb-10 relative md:pt-14 
        before:absolute before:content-[''] before:h-40 before:w-full before:top-1 before:-translate-y-[100%] before:bg-gradient-to-b before:from-transparent before:via-blue-100 before:via-80% before:to-blue-100
        after:absolute after:content-[''] after:h-40 after:w-full after:bottom-1 after:translate-y-[100%] after:bg-gradient-to-b after:to-transparent after:from-5% after:via-20% after:via-blue-100/90 after:from-blue-100`}
    >
      <SendForConsult />
      <div
        className={
          "w-96 h-96 z-10 relative md:self-start md:ml-10 lg:w-431 lg:h-96 2xl:ml-20 2xl:w-540 2xl:h-500"
        }
      >
        <BookGuySVG
          className={
            "object-cover object-center z-10 absolute top-36 w-full h-full 2xl:top-48"
          }
        />
      </div>
      <FAQContainer />
      <MainButton
        className=" absolute bottom-1%"
        colorSchema={" btn-accent-white"}
        title="На консультацию"
      />
    </section>
  );
};
