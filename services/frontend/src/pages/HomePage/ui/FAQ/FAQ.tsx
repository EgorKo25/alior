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
      <div className={"w-96 h-72 z-10 relative md:self-start md:ml-10"}>
        <img
          className={"object-cover object-center z-10 absolute top-8"}
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
