import { Tablet1024SVG } from "./ui/Tablet1024SVG";
import { Tablet768SVG } from "./ui/Tablet768SVG";
import { TabletMobileSVG } from "./ui/TabletMobileSVG";
import { Tv1440SVG } from "./ui/Tv1440SVG";
import { Tv1920SVG } from "./ui/Tv1920SVG";

export const Placeholders = () => {
  return (
    <div className="img-wrapper w-full h-full grid place-items-center">
      <TabletMobileSVG />
      <Tablet768SVG />
      <Tablet1024SVG />
      <Tv1440SVG />
      <Tv1920SVG />
    </div>
  );
};
