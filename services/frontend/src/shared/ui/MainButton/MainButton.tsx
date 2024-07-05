import { CircleIcon } from "./CircleIcon";
import { ArrowIcon } from "./ArrowIcon";

export const MainButton: React.FC<{
  title: string;
  onClick?: () => void;
  type?: React.ButtonHTMLAttributes<HTMLButtonElement>["type"];
  colorSchema?:
    | " btn-accent-white"
    | " btn-white-accent"
    | " btn-orange-white"
    | " btn-blue-white"
    | " btn-black-white"
    | " btn-white-black";
  className: string;
}> = ({ title, onClick, type, colorSchema, className }) => {
  return (
    <button
      onClick={onClick}
      className={
        className +
        ` main-button group h-12 text-center rounded-[32px] cursor-pointer leading-6 border font-medium gap-2 py-2 px-3.5 md:py-3 md:px-5` +
        colorSchema
      }
      type={type}
    >
      <div className={`flex items-center justify-center mx-auto`}>
        <ArrowIcon />
        <span className=" title translate-x-0 animate-textBackwards -ml-6 group-hover:animate-textForwards">
          {title}
        </span>
        <CircleIcon />
      </div>
    </button>
  );
};
