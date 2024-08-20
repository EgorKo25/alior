import { ButtonColorSchema, MainButton } from "src/shared/ui/MainButton";
import { TabletsCanvas } from "./TabletsCanvas";
import { Placeholders } from "./Placeholders";

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
      <div className="canvas-wrapper w-full h-500 md:w-[350px] lg:w-[500px] lg:h-598 xl:w-[700px] 2xl:w-[950px] 2xl:h-[700px] md:absolute md:-right-5 md:-top-[15%] 2xl:-top-[20%]">
        {isActive ? <TabletsCanvas url={projectUrl} /> : <Placeholders />}
      </div>

      <div className={` flex flex-col gap-10 md:w-[53%] md:mt-10 xl:w-[45%]`}>
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
        <a href={projectUrl} target="_blank">
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
