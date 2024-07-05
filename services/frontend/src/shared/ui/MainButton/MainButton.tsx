import clsx from "clsx";
import styles from "./MainButton.module.scss";
import { CircleIcon } from "./CircleIcon";
import { ArrowIcon } from "./ArrowIcon";

export const MainButton: React.FC<{
  title: string;
  onClick?: () => void;
  type?: React.ButtonHTMLAttributes<HTMLButtonElement>["type"];
  className?: string;
  circleClassName: string;
}> = ({ title, onClick, type, className, circleClassName }) => {
  return (
    <button
      onClick={onClick}
      className={clsx(styles.button, className)}
      type={type}
    >
      <div className={styles.arrow}>
        <ArrowIcon />
      </div>
      <span className={styles.title}>{title}</span>
      <CircleIcon className={" " + circleClassName} />
    </button>
  );
};
