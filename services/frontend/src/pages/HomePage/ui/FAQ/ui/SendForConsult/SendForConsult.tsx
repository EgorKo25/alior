export const SendForConsult = () => {
  return (
    <div
      className={
        " absolute -top-4% rounded-3xl bg-accent px-5 py-16 w-full flex flex-col gap-6 md:items-center md:pt-6 md:pb-8 md:px-10 md:justify-between md:flex-row md:w-90% lg:gap-0 lg:px-28 xl:pt-8 xl:pb-10 xl:w-4/5 2xl:w-95%"
      }
    >
      <h2
        className={` flex flex-col  font-bold text-40 leading-[1] md:flex-row md:items-center md:gap-4 whitespace-nowrap`}
      >
        Идём на
        <span
          className={` font-caveat font-normal text-white text-64 md:leading-[0.9] md:-mt-4`}
        >
          консультацию
        </span>
      </h2>
      <button className=" group rounded-full flex size-14 justify-center items-center bg-white text-black cursor-pointer hover:bg-black hover:text-white">
        <svg
          className=" group-hover:animate-arrowForwards"
          width="12"
          height="16"
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
    </div>
  );
};
