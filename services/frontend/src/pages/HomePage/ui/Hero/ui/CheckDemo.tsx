import { HashLink } from "react-router-hash-link";
import { MainButton } from "src/shared/ui/MainButton";

export const CheckDemo = () => {
  return (
    <div className="demo-suggestion flex flex-col mt-10 sm:flex-row sm:gap-10 lg:mt-16 xl:flex-col xl:justify-between xl:gap-12">
      <h2 className=" font-bold text-[40px] leading-[45px]  lg:text-[64px] sm:mb-0 xl:mt-20">
        <span className="font-caveat font-medium text-accent text-[64px] leading-[70px]  lg:text-[96px]">
          Посмотрите
        </span>
        <br />
        демо-версии
      </h2>
      <HashLink smooth to={"#cases"}>
        <MainButton
          title={"Ознакомиться"}
          className={"w-fit mt-10 sm:self-end xl:self-start xl:mb-10 xl:mt-0"}
          colorSchema=" btn-white-accent"
        />
      </HashLink>
    </div>
  );
};
