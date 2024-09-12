import clsx from "clsx";
import { FeaturesThesesData } from "src/features/appData";
export const Thesis: React.FC<{
  index: number;
  className: string;
}> = ({ index, className }) => {
  return (
    <div className={clsx(` flex flex-col`, className)}>
      <h3 className=" font-caveat font-bold text-2xl lg:text-3xl xl:text-4xl">
        {FeaturesThesesData[index].title}
      </h3>
      <p className=" font-light text-sm lg:text-base xl:text-lg">
        {FeaturesThesesData[index].description}
      </p>
    </div>
  );
};
