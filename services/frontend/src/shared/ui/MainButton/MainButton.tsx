import clsx from "clsx";
import styles from "./MainButton.module.scss";
import { CircleIcon } from "./CircleIcon";
import { ArrowIcon } from "./ArrowIcon";

export const MainButton = ({
  title,
  onClick,
  type,
  className,
  textColor,
}: {
  title: string;
  onClick?: () => void;
  type?: React.ButtonHTMLAttributes<HTMLButtonElement>["type"];
  className?: string;
  textColor: string;
}) => {
  return (
    <button
      onClick={onClick}
      className={clsx(styles.button, className)}
      type={type}
    >
      <div className={styles.arrow}>
        <ArrowIcon color={textColor} />
      </div>
      <span className={styles.title} style={{ color: textColor }}>
        {title}
      </span>
      <CircleIcon color={textColor} />
    </button>
  );
};
