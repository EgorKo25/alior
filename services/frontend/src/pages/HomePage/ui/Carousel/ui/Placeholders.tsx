export const Placeholders = () => {
  return (
    <div className="img-wrapper w-full h-full grid place-items-center">
      {/* <LogoSVG className={" scale-[1.8] xl:scale-[2.5] 2xl:scale-[4]"} /> */}
      <img
        src="./images/slider-placeholders/mobile-tablet.png"
        alt="Alior"
        className=" xl:hidden"
      />
      <img
        src="./images/slider-placeholders/tv-desktop.png"
        alt="Alior"
        className="hidden xl:block opacity-0 xl:opacity-100"
      />
    </div>
  );
};
