import { MainButton } from "src/shared/ui/MainButton";
import { FAQContainer } from "./ui/FAQContainer";
import { SendForConsult } from "./ui/SendForConsult";

export const FAQ = () => {
  return (
    <section
      className={
        " my-0 mx-auto mt-48 pt-48 flex flex-col items-center justify-center w-full bg-blue-100 pb-10 relative md:pt-14"
      }
    >
      <SendForConsult />
      <div
        className={
          "w-96 h-96 z-10 relative md:self-start md:ml-10 lg:w-431 lg:h-96 2xl:ml-20 2xl:w-540 2xl:h-500"
        }
      >
        <img
          className={
            "object-cover object-center z-10 absolute top-36 w-full h-full 2xl:top-48"
          }
          src="/images/bookworm.svg"
        ></img>
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
