export const Guys = () => {
  return (
    <div className="guys flex flex-col mt-64 gap-20 md:grid md:grid-cols-2 md:gap-4 lg:w-[60vw] lg:mt-40 xl:w-[42vw] xl:mt-32">
      <div className="selfie-guy flex flex-col items-center border relative rounded-3xl w-64 h-56 md:w-auto md:h-[330px] lg:h-[295px]">
        <img
          src="./images/selfie.svg"
          alt="Guy With Phone"
          className="absolute top-0 -translate-y-1/2 w-48 md:w-72 lg:w-64"
        />
        <h4 className="mt-auto mb-7 font-bold text-xl md:mb-11">
          Разрабатываем
        </h4>
      </div>
      <div className="guy-in-box flex flex-col items-center border relative rounded-3xl w-64 h-48 ml-auto md:w-auto md:ml-0 md:h-72 lg:h-64">
        <img
          src="./images/inthebox.svg"
          alt="Guy In The Box"
          className="absolute top-0 -translate-y-1/3 w-44 md:w-64 "
        />
        <h4 className="mt-auto mb-5 font-bold text-xl md:mb-10">
          Поддерживаем
        </h4>
      </div>
    </div>
  );
};
