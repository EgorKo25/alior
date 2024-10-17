import { HashLink } from "react-router-hash-link";

export const SendForConsult = () => {
  return (
    <div
      className={
        " absolute -top-14 rounded-3xl bg-accent px-5 py-10 w-full flex flex-col gap-10 md:items-center  md:px-10 md:justify-between md:flex-row md:w-90% lg:gap-0 lg:px-20 xl:px-40 xl:w-4/5 2xl:w-90%"
      }
    >
      <h2
        className={` flex flex-col  font-bold text-40 leading-[1] md:flex-row md:items-center md:gap-4 lg:text-64 whitespace-nowrap text-center`}
      >
        Идём на
        <span
          className={` font-caveat font-normal text-white text-64 md:leading-[0.9] md:-mt-4 lg:text-[80px]`}
        >
          консультацию
        </span>
      </h2>
      <HashLink smooth to={"#consult"}>
        <button className=" group rounded-full flex size-14 lg:size-16 xl:size-20 justify-center items-center bg-white text-black cursor-pointer hover:bg-black hover:text-white mx-auto">
          <svg
            className=" group-hover:animate-arrowForwards"
            width="30%"
            height="50%"
            viewBox="0 0 12 16"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M0 10L1.45 8.6L5 12.15L5 0L7 0L7 12.15L10.55 8.6L12 10L6 16L0 10Z"
              fill="currentColor"
            />
          </svg>
        </button>
      </HashLink>
    </div>
  );
};
