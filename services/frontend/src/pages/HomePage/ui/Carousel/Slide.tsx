import { ButtonColorSchema, MainButton } from "src/shared/ui/MainButton";
import { TabletCanvas } from "./ui/TabletCanvas";

const COLOR_SCHEMA_CONFIG: Record<
  string,
  {
    buttonColorSchema: ButtonColorSchema;
    background: string;
    textColor: string;
  }
> = {
  orange: {
    buttonColorSchema: " btn-orange-white",
    background: " bg-orange-100",
    textColor: " text-orange-900",
  },
  blue: {
    buttonColorSchema: " btn-blue-white",
    background: " bg-blue-100",
    textColor: " text-blue-900",
  },
  black: {
    buttonColorSchema: " btn-black-white",
    background: " bg-black",
    textColor: " text-white",
  },
};

export const Slide: React.FC<{
  ColorSchema: "orange" | "blue" | "black";
  title: string;
  description: string;
  projectUrl: string;
  isActive: boolean;
}> = ({ ColorSchema, title, description, projectUrl, isActive }) => {
  return (
    <div
      className={
        `slide px-5 lg:px-20 py-12 xl:py-20 rounded-40 relative mx-2.5 xl:mx-5` +
        COLOR_SCHEMA_CONFIG[ColorSchema].background
      }
    >
      {isActive ? (
        <div className="canvas-wrapper md:w-[50%] md:absolute md:-right-5 h-500 md:h-[130%] md:-top-[15%] xl:w-[60%]">
          <TabletCanvas url={projectUrl} />
        </div>
      ) : (
        <img
          src="/images/Rectangle.png"
          className=" w-full h-500 -mt-20 mb-5 overflow-visible md:w-1/2 md:absolute md:-right-5 md:top-1/2 md:translate-y-[-35%]"
        />
      )}

      <div className={` flex flex-col gap-10 md:w-[55%] md:mt-10 xl:w-[43%]`}>
        <h2
          className={
            ` font-bold text-40 lg:text-64 leading-none` +
            COLOR_SCHEMA_CONFIG[ColorSchema].textColor
          }
        >
          {title}
        </h2>
        <div
          className={
            ` text-sm line-clamp-4 h-[6em] lg:text-base` +
            COLOR_SCHEMA_CONFIG[ColorSchema].textColor
          }
        >
          {description}
        </div>
        <a href={projectUrl}>
          <MainButton
            title={"Посмотреть демо"}
            className={"w-full"}
            colorSchema={COLOR_SCHEMA_CONFIG[ColorSchema].buttonColorSchema}
          />
        </a>
      </div>
    </div>
  );
};
